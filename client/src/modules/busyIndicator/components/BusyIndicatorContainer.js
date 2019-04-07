import React, { Component } from "react";
import { connect } from "react-redux";
import BusyIndicator from "./BusyIndicator";
import { isBusy } from "../busyIndicator.selector";

class BusyIndicatorContainer extends Component {
  render() {
    return <BusyIndicator isBusy={this.props.isBusy} />;
  }
}

const mapStateToProps = state => ({
  isBusy: isBusy(state)
});

export default connect(mapStateToProps)(BusyIndicatorContainer);
