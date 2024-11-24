import React from 'react'
import "./Review.css";

export default function Review() {
  return (
    <div className="table-container">
      <h2>Reviews</h2>
      <table className="styled-table">
        <thead>
          <tr>
            <th>Property title</th>
            <th>Property title</th>
            <th>Review time</th>
            <th>Your review</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>row.id</td>
            <td>row.name</td>
            <td>row.age</td>
            <td>"row.occupation"</td>
            <td><button className='table-delete'>Delete</button></td>
          </tr>
        </tbody>
      </table>
    </div>
  )
}
