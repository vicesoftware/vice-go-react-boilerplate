import React, { Component } from "react";
import { connect } from "react-redux";
import userContext from "../../../modules/userContext";
import contacts from "../../../modules/contacts";

const { getUserContext, isAuthenticated } = userContext.selectors;
const {
  components: { Contacts }
} = contacts;

class HomeContainer extends Component {
  render() {
    const { isAuthenticated } = this.props;

    return (
      <div>
        {!isAuthenticated ? (
          <div>
            <h1>Welcome Guest</h1>
            <p>
              You can login as ryan@vicesoftware.com with {"'password'"} for
              your password to see how authenctication works.
            </p>
          </div>
        ) : (
          <h1>Hi {this.props.userContext.displayName}</h1>
        )}
        <Contacts />
      </div>
    );
  }
}

const mapDispatchToProps = {
  ...userContext.actions
};

const mapStateToProps = state => ({
  userContext: getUserContext(state),
  isAuthenticated: isAuthenticated(state)
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(HomeContainer);
