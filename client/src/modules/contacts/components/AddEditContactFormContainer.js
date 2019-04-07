import React, { Component } from "react";
import { connect } from "react-redux";
import AddEditContactForm from "./AddEditContactForm";
import * as actions from "../contacts.actions";

class AddEditFormContainer extends Component {
  onSubmit = values => {
    this.props.saveContact({ ...values }).then(this.props.getContacts);
  };

  render() {
    return <AddEditContactForm onSubmit={this.onSubmit} />;
  }
}

export default connect(
  null,
  { saveContact: actions.saveContact, getContacts: actions.getContacts }
)(AddEditFormContainer);
