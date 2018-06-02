import React, { Component } from 'react';
import './App.css';
import QrScanner from './components/qr_scanner';

class App extends Component {
  render() {
    return (
      <div className="App">
        <QrScanner/>
      </div>
    );
  }
}

export default App;
