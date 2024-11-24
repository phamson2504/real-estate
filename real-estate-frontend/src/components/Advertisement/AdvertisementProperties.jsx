import React from 'react'
import "./AdvertisementProperties.css"

export default function AdvertisementProperties() {
    return (
        <div className="popular-properties">
            <h2>Our Popular Properties</h2>
            <div className="properties-grid">
                <div className="property-card">
                    <img src="./propertiesImage/property1.jpg" alt="Modern Luxury Apartment with Stunning Views" className="property-image" />
                    <div className="property-details">
                        <h3>Modern Luxury Apartment with Stunning Views</h3>
                        <p className="property-address">123 Main Street, Cityville</p>
                        <p className="property-price">$300,000 - $350,000</p>
                        <p className="property-status">Status: <span>verified</span></p>
                        <button className="details-button">Details</button>
                    </div>
                </div>
            </div>
            <button className="see-all-button">SEE ALL</button>
        </div>
    )
}
