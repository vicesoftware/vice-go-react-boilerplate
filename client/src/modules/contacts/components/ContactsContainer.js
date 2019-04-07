import React, { Component } from "react";
import { connect } from "react-redux";
import * as actions from "../contacts.actions";
import { getContacts } from "../contacts.selectors";
import Contacts from "./Contacts";

class ContactsContainer extends Component {
  componentDidMount() {
    this.props.getContacts();
  }

  render() {
    return <Contacts contacts={this.props.contacts} />;
  }
}

const mapDispatchToProps = {
  getContacts: actions.getContacts
};

const mapStateToProps = state => ({
  contacts: getContacts(state)
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(ContactsContainer);
