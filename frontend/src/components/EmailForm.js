import React from 'react'
import InputText from './InputText'
import TextArea from './TextArea'

const EmailForm = ({ handleSubmit, handleChange, to, from, subject, body, cc, bcc }) => (
  <form className="form-horizontal col-md-9 col-md-offset-1 col-xs-12" onSubmit={handleSubmit}>
    <div className="form-group">
      <label htmlFor="daily_date" className="col-md-3 col-sm-2 control-label">To</label>
      <div className="col-md-8 col-sm-10" >
        <InputText name="to" type="text" defaultValue={to} handleChange={handleChange} required={true} />
      </div>
    </div>
    <div className="form-group">
      <label htmlFor="cc" className="col-md-3 col-sm-2 control-label">Cc</label>
      <div className="col-md-8 col-sm-10" >
        <InputText name="cc" type="text" defaultValue={cc} handleChange={handleChange} required={false} />
      </div>
    </div>
    <div className="form-group">
      <label htmlFor="bcc" className="col-md-3 col-sm-2 control-label">Bcc</label>
      <div className="col-md-8 col-sm-10" >
        <InputText name="bcc" type="text" defaultValue={bcc} handleChange={handleChange} required={false} />
      </div>
    </div>
    <div className="form-group">
      <label htmlFor="from" className="col-md-3 col-sm-2 control-label">From</label>
      <div className="col-md-8 col-sm-10" >
        <InputText name="from" type="text" handleChange={handleChange} required={true} />
      </div>
    </div>
    <div className="form-group">
      <label htmlFor="subject" className="col-md-3 col-sm-2 control-label">Subject</label>
      <div className="col-md-8 col-sm-10" >
        <InputText name="subject" type="text" handleChange={handleChange} required={true} />
      </div>
    </div>
    <div className="form-group">
      <label htmlFor="body" className="col-md-3 col-sm-2 control-label">Body</label>
      <div className="col-md-8 col-sm-10" >
        <TextArea name="body" row="5" handleChange={handleChange} required={true} />
      </div>
    </div>
    <div className="form-group">
      <div className="col-sm-offset-2 col-sm-10">
        <button type="submit" className="btn btn-success">Submit</button>
      </div>
    </div>
  </form>
)

export default EmailForm
