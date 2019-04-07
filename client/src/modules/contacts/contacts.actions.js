import * as actionTypes from "./contacts.actionTypes";
import doAsync from "../doAsync";

export const getContacts = () => {
  return doAsync({
    actionType: actionTypes.GET_CONTACTS_ASYNC,
    url: "contacts",
    mapResponseToPayload: r => r,
    errorMessage: "Unable to retrieve contacts."
  });
};

export const saveContact = contact => {
  return doAsync({
    actionType: actionTypes.SAVE_CONTACT_ASYNC,
    url: "contacts",
    httpMethod: "post",
    httpConfig: {
      body: JSON.stringify(contact)
    },
    mapResponseToPayload: r => r,
    errorMessage: "Unable to retrieve contacts."
  });
};
