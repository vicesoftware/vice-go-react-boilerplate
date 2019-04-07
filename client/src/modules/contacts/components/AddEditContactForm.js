import React from "react";
import { reduxForm, Field } from "redux-form";
import forms from "../../forms";

const {
  components: { ValidatingInputField },
  validators: { required }
} = forms;

const AddEditContactForm = props => (
  <form onSubmit={props.handleSubmit(props.onSubmit)}>
    <Field
      component={ValidatingInputField}
      label="First Name"
      name="firstName"
      type="text"
      validators={[required]}
    />
    <Field
      component={ValidatingInputField}
      label="Last Name"
      name="lastName"
      type="text"
      validators={[required]}
    />
    <button type="submit">Add</button>
  </form>
);

export default reduxForm({
  form: "contactsForm"
})(AddEditContactForm);
