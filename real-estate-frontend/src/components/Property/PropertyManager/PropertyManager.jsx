import React from 'react'
import './PropertyManager.css'

export default function PropertyManager({properties}) {
    return (
        <div className="pm-card" key={properties.id}>
            <img src={properties.image[0].url} alt="Modern Luxury Apartment" className="pm-image" />
            <div className="pm-details">
                <h3 className="pm-title">{properties.title}</h3>
                <p className="pm-address">{properties.location}</p>
                <div className="pm-agent-info">
                    <div className="pm-agent">
                        <img src={properties.agent.avatarAgent||"/user.jpg"} alt="Agent Icon" className="pm-agent-icon" />
                        <span className="pm-agent-name">{properties.agent.name}</span>
                    </div>
                </div>
                <div className="pm-status">
                    <span className="pm-status-label">Status:</span>
                    <span className="pm-status-verified" 
                    style={  properties.status === "spending"
                        ? { color: "#bfce48" }
                        : properties.status === "verify"
                        ? { color: "green" }
                        : { color: "#ff4d4d" }
                    }
                    >{properties.status}</span>
                </div>
                <p className="pm-price-range">Price Range: <span>{properties.minPrice} - {properties.maxPrice}</span></p>
                <div className="pm-actions">
                    <button className="pm-update-btn">Update</button>
                    <button className="pm-remove-btn">Remove</button>
                </div>
            </div>
        </div>
    )
}
