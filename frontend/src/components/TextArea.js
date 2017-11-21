import React from 'react'

const TextArea = ({ name, row, value, handleChange, required }) => (
  <textarea className="form-control" name={name} id={name} rows={row} value={value} onChange={e => handleChange(e, name)} required={required} />
)

export default TextArea
