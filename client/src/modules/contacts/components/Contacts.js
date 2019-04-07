import React, { Fragment } from "react";

const Address = ({ line1 }) => (
  <ul>
    <li>{line1}</li>
  </ul>
);

const Contact = ({ firstName, lastName, addresses }) => (
  <Fragment>
    <li>{firstName + " " + lastName}</li>
    {addresses && addresses.map(a => <Address key={a.id} {...a} />)}
  </Fragment>
);

const Contacts = ({ contacts }) => (
  <Fragment>
    <p>Below are the contacts and addresses we got from the go server.</p>
    {contacts && (
      <ul>
        {contacts.map(c => (
          <Contact key={c.id} {...c} />
        ))}
      </ul>
    )}
  </Fragment>
);

export default Contacts;
