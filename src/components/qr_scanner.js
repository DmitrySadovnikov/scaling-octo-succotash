import React, { Component } from 'react';
import QrReader from 'react-qr-reader';
import axios from 'axios';

export default class QrScanner extends Component {
  constructor(props) {
    super(props);
    this.state      = {
      delay:  300,
      result: 'No result',
    };
    this.handleScan = this.handleScan.bind(this);
  }

  handleScan(data) {
    if (data) {
      axios({
        method:  'GET',
        url:     `http://localhost:3302/meats/${data}`,
        headers: {
          Accept:         'application/json',
          'Content-Type': 'application/json',
        },
      }).then(json => {
        this.setState({
          result: json.data.message,
        });
      });
    }
  }

  handleError(err) {
    console.error(err);
  }

  render() {
    return (
      <div>
        <div>
          <h2>Поднесите qr код продукта к сканеру</h2>
        </div>
        <QrReader
          delay={ this.state.delay }
          onError={ this.handleError }
          onScan={ this.handleScan }
          style={ { width: '30%', margin: 'auto' } }
        />
        <p>{ this.state.result }</p>
      </div>
    );
  }
}
