import { buildAsyncActionType } from "../utilities/reduxUtilities";

export const GET_CONTACTS_ASYNC = buildAsyncActionType(
  "contacts",
  "GET_CONTACTS_ASYNC"
);

export const SAVE_CONTACT_ASYNC = buildAsyncActionType(
  "contacts",
  "SAVE_CONTACTS_ASYNC"
);
