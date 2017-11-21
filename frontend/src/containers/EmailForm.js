import React, { Component } from 'react';
import EmailFormComponent from '../components/EmailForm';
import Request from '../libs/Request';

export class EmailForm extends Component {

  constructor(props) {
    super(props)
    this.state = {
      to: '',
      from: '',
      subject: '',
      body: '',
    }
  }

  handleSubmit = (e) => {
    e.preventDefault()
    console.log(e.target.to.value);
    var data = {
      to: e.target.to.value,
      from: e.target.from.value,
      cc: e.target.cc.value,
      bcc: e.target.bcc.value,
      subject: e.target.subject.value,
      body: e.target.body.value,
    }
    Request.post("http://localhost:8080/email", data)
    .then((resp) => {
      e.target.to.value = '';
      e.target.subject.value = '';
      e.target.body.value = '';
      alert('Done');
      console.log(resp)
    }) 
    .catch(err => console.log(err));
  }

  handleChange = (event, fieldName) => {
    let state = {}
    state[fieldName] = event.target.value
    console.log(event.target.value)
    this.setState(state)
  }

  render() {
    const { to, from, subject, body, cc, bcc } = this.props;
    return (
      <EmailFormComponent handleSubmit={this.handleSubmit} handleChange={this.handleChange} to={to} from={from} subject={subject} body={body} cc={cc} bcc={bcc} />
    );
  }
}

export default EmailForm
