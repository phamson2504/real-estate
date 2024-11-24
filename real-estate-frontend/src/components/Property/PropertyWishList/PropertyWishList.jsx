import React from 'react'
import "./PropertyWishList.css"

export default function PropertyWishList({property}) {
    return (
        <div className='wl-list-grid'>
            <div className='wl-list-card'>
                <img src='/PropertiesImage/property1.jpg' alt="Amage property" />
                <div className='wl-list-detail'>
                    <h2>{property.title}</h2>
                    <p className='wl-list-address'>{property.location}</p>
                    <p className='wl-list-price'> Price Range: {property.minPrice}$ - {property.maxPrice}$</p>
                    <div className='wl-list-status'>
                        <span>{property.status}</span>
                    </div>
                    <div className='wl-list-agent-info'>
                        <img src={property.agent.avatarAgent} alt='Amage agent' />
                        <p>{property.agent.name}</p>
                    </div>
                    <div className='wl-list-action'>
                        <button className='wl-list-make-offer'>Make Offer</button>
                        <button className='wl-list-remove'>Remove</button>
                    </div>
                </div>
            </div>
        </div>
    )
}
