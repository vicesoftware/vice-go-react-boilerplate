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
