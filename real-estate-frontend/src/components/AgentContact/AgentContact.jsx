import React from 'react'
import './AgentContact.css'

export default function AgentContact({agent}) {
  return (
    <div className='ac-contain'>
        <img src={agent.avatarAgent || "/user.jpg"} alt='Agent'/>
        <p>Professional Brokers</p>
        <h3>{agent.name}</h3>
        <div className='ac-contact'>
            <p>{agent.email}</p>
            <p>{agent.contact}</p>
        </div>
    </div>
  )
}
