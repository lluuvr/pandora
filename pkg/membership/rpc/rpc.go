package rpc

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"pandora/pkg/conf"
	"pandora/pkg/distribution"
	"pandora/pkg/pb"
	"pandora/pkg/utils/errors"
	"pandora/pkg/utils/log"
	"pandora/pkg/utils/network"
)

// RPC
type RPC struct {
	master   pb.MasterClient
	masterCC *grpc.ClientConn
}

// New returns mew membership rpc
func New() (*RPC, error) {
	creds, err := credentials.NewClientTLSFromFile(conf.Conf.TLS.Cert, "")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	discoveryCC, err := grpc.Dial(conf.Conf.Discovery.Endpoint, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer func() {
		_ = discoveryCC.Close()
	}()

	discoveryC := pb.NewDiscoveryClient(discoveryCC)

	ino := &pb.InitNetworkOpts{}
	tick := time.NewTicker(time.Millisecond * 500).C
	timer := time.NewTimer(time.Second * 3).C

loop:
	for {
		select {
		case <-tick:
			if ino, err = discoveryC.InitMembership(
				context.Background(),
				&pb.Endpoint{Endpoint: conf.Conf.Membership.Endpoint},
			); err != nil {
				continue
			}
			break loop
		case <-timer:
			return nil, errors.WithStack(err)
		}
	}

	masterCC, err := grpc.Dial(ino.Master, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	masterC := pb.NewMasterClient(masterCC)

	return &RPC{master: masterC, masterCC: masterCC}, nil
}

// Close close connection with others rpc
func (rpc *RPC) Close() {
	if err := rpc.masterCC.Close(); err != nil {
		log.Error(errors.WithStack(err))
	}
}

// ProposeMember propose member over rpc
func (*RPC) ProposeMember(ctx context.Context, in *pb.MemberMeta) (*pb.PublicKey, error) {
	key, err := distribution.NewMembership().ConfirmMember(in)
	if err != nil {
		return &pb.PublicKey{}, err
	}
	return key, nil
}

// SignCert sign cert over rpc
func (rpc *RPC) SignCert(ctx context.Context, in *pb.Cert) (*pb.Empty, error) {
	cert, err := distribution.NewMembership().SignCert(in)
	if err != nil {
		if err == errors.ErrNotFound {
			return &pb.Empty{}, status.Error(codes.NotFound, codes.NotFound.String())
		}
		return &pb.Empty{}, errors.WithStack(err)
	}

	if _, err := rpc.master.ProposeCert(context.Background(), cert); err != nil {
		return &pb.Empty{}, errors.WithStack(err)
	}

	return &pb.Empty{}, nil
}

// FetchMember fetch member over rpc
func (*RPC) FetchMember(ctx context.Context, in *pb.PublicKey) (*pb.Member, error) {
	mem, err := distribution.NewMembership().MemberFetch(in)
	if err != nil {
		if err == errors.ErrNotFound {
			return &pb.Member{}, status.Error(codes.NotFound, codes.NotFound.String())
		}
		return &pb.Member{}, err
	}
	return mem, nil
}

// Listen listen for rpc requests
func (rpc *RPC) Listen() error {
	creds, err := credentials.NewServerTLSFromFile(conf.Conf.TLS.Cert, conf.Conf.TLS.Key)
	if err != nil {
		return errors.WithStack(err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	defer s.GracefulStop()

	pb.RegisterMembershipServer(s, rpc)

	listen, err := net.Listen(network.TCP, network.PortWithSemicolon(conf.Conf.Membership.Endpoint))
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() {
		_ = listen.Close()
	}()

	if err := s.Serve(listen); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
