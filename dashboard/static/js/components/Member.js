class MemberLayout extends React.Component {
  constructor() {
    super();
    this.state = {
      state: STATE.FETCH,
      pending: false,
      error: null,
      success: null,
      data: this.EMPTY_DATA,
      member: null
    };
  }

  EMPTY_DATA = {
    name: "",
    email: "",
    publicKey: ""
  };

  handleState = (e, key) => {
    e.preventDefault();
    this.setState({state: key});
  };

  handleChange = (e) => {
    const target = e.target;
    const name = target.name;
    const value = target.value;
    this.setState({data: {...this.state.data, [name]: value}});
  };

  handleCreate = async () => {
    this.setState({pending: true});

    try {
      await api.MemberCreate({name: this.state.data.name, email: this.state.data.email});
      this.setState({success: "Member successfully confirmed"});
    } catch (e) {
      this.setState({error: e});
    }

    this.setState({pending: false});
  };

  handleClose = () => {
    this.setState({member: null, error: null, success: null, data: this.EMPTY_DATA});
  };

  handleFetch = async () => {
    this.setState({pending: true});

    try {
      const member = await api.MemberFetch({public_key: this.state.data.publicKey});
      this.setState({member: member, pending: false});
    } catch (e) {
      this.setState({error: e, pending: false});
    }

    this.setState({pending: false});
  };

  render() {
    if (this.state.pending) return <Preloader/>;
    if (this.state.error) return <Error error={this.state.error} close={this.handleClose}/>;
    if (this.state.success) return <Alert text={this.state.success} close={this.handleClose}/>;
    if (this.state.member) return <MemberCard member={this.state.member} close={this.handleClose}/>;

    return (
      <div className="card shadow">
        <div className="card-header">
          <ul className="nav nav-pills nav-fill">
            {
              Object.keys(STATE).map((key, index) => (
                <li
                  className="nav-item"
                  key={index}
                >
                  <a
                    className={`nav-link ${this.state.state === key && "active"}`}
                    onClick={(e) => this.handleState(e, key)}
                    href="#"
                  >
                    {key.capitalize()}
                  </a>
                </li>
              ))
            }
          </ul>
        </div>

        <div className="card-body">
          {
            this.state.state === STATE.FETCH ?
              <MemberFetchCard
                publicKey={this.state.data.publicKey}
                change={this.handleChange}
                submit={this.handleFetch}
              /> :
              <MemberCreateCard
                change={this.handleChange}
                name={this.state.data.name}
                email={this.state.data.email}
                submit={this.handleCreate}
              />
          }
        </div>
      </div>
    );
  }
}

const MemberCard = ({member, close}) => (
  <div className="card shadow">
    <div className="card-header">Member</div>

    <div className="card-body">
      <h5 className="card-title">{member.meta.name}</h5>
      <p className="card-text">{member.meta.email}</p>
    </div>

    <div className="card-footer">
      <button
        className="close"
        onClick={close}
      >
        <span>&times;</span>
      </button>
    </div>
  </div>
);

MemberCard.propTypes = {
  member: PropTypes.object.isRequired,
  close: PropTypes.func.isRequired
};

const MemberCreateCard = ({email, name, submit, change}) => (
  <section>
    <div className="form-group">
      <label>Email</label>
      <input
        name="email"
        onChange={change}
        value={email}
        type="email"
        className="form-control"
        placeholder="thomas@mail.sys"
      />
    </div>

    <div className="form-group">
      <label>Name</label>
      <input
        name="name"
        onChange={change}
        value={name}
        type="text"
        className="form-control"
        placeholder="Thomas Hobbs"
      />
    </div>

    <button
      className="btn btn-primary"
      onClick={submit}
    >
      Submit
    </button>
  </section>
);

MemberCreateCard.propTypes = {
  name: PropTypes.string.isRequired,
  email: PropTypes.string.isRequired,
  submit: PropTypes.func.isRequired,
  change: PropTypes.func.isRequired
};

const MemberFetchCard = ({submit, change, publicKey}) => (
  <section>
    <div className="form-group">
      <label>Public key</label>
      <input
        name="publicKey"
        onChange={change}
        value={publicKey}
        type="text"
        className="form-control"
        placeholder="787c8ef36e46f02a58f014ac7507c27fb29e757d0ca323ffd8d517ec70e3caa9"
      />
    </div>

    <button
      className="btn btn-primary"
      onClick={submit}
    >
      Submit
    </button>
  </section>
);

MemberFetchCard.propTypes = {
  submit: PropTypes.func.isRequired,
  change: PropTypes.func.isRequired,
  publicKey: PropTypes.string.isRequired
};
