import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
// import App from './App';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
// import { Router, Route, Link } from 'react-router'
import {
  HashRouter as Router,
  Route
} from 'react-router-dom';

import registerServiceWorker from './registerServiceWorker';
import Dashboard from './view/dashboard/Dashboard'
import NoMatch from './view/NoMatch'

const App = () => (
  <Router>
    <MuiThemeProvider>
    <div>
     <Route path="/" component={Dashboard}>
     <Route component={NoMatch}></Route>
     
      </Route>
      
    </div>

    </MuiThemeProvider>
  </Router>
);
ReactDOM.render(<App />, document.getElementById('root'));
registerServiceWorker();
