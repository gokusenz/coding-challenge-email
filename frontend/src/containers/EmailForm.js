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
    }
  }

  componentDidMount() {
  }

  handleSubmit = (e) => {
    e.preventDefault()
    console.log(e.target.to.value);
    Request.post(SERVICE_CREATE_LOG, data)
    .catch(err => console.log(err));
    // if (result) {
    //   alert('บันทึกข้อมูลเรียบร้อย')
    //   e.target.to.value = ''
    // } else {
    //   alert('ส่งข้อมูลไม่สำเร็จ กรุณาลองใหม่อีกครั้ง')
    // }
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
