import React, { Component } from 'react';
import MyAwesomeReactComponent from '../../components/MyAwesomeReactComponent';

import './dashboard.css';

class Dashboard extends Component {
  render() {
    return (
      <div className="App">
           <MyAwesomeReactComponent />
      </div>
    );
  }
}

export default Dashboard;
