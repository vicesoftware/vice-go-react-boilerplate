import React, { Fragment } from "react";
import Indicator from "./indicator.svg";
import "./busyIndicator.css";

const BusyIndicator = ({ isBusy }) => (
  <Fragment>
    {isBusy && <img className="loader" src={Indicator} alt="Busy indicator" />}
  </Fragment>
);

export default BusyIndicator;
