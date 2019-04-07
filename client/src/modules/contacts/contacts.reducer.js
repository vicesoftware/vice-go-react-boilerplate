import * as actionTypes from "./contacts.actionTypes";

export default function reducer(state = [], action) {
  switch (action.type) {
    case actionTypes.GET_CONTACTS_ASYNC.RECEIVED:
      return action.payload;

    default:
      return state;
  }
}
