import React from 'react'

const InputText = ({ name, type, defaultValue, handleChange, required }) => (
  <input type={type} className="form-control" name={name} id={name} defaultValue={defaultValue} onChange={e => handleChange(e, name)} required={required} />
)

export default InputText
