import React from 'react'

export default function PropertySold({transactions}) {
    console.log(transactions)
    return (
        <div className='pb-list-card' key={transactions.id}>
            <img src='/PropertiesImage/property1.jpg' alt="Amage property" />
            <div className='pb-list-detail'>
                <h2>{transactions.Properties.title}</h2>
                <p className='pb-list-address'>{transactions.Properties.location}</p>
                <p className='pb-list-price'> Price Range: {transactions.Properties.minPrice}$ - {transactions.Properties.maxPrice}$</p>
                <p className='pb-list-price-bought'> Price Sold: {transactions.amount}$</p>
                <div className='pb-list-agent-info'>
                    <img src={transactions.Properties.agent.avatarAgent} alt='Amage agent'/>
                    <p>{transactions.Properties.agent.name}</p>
                </div>
                <p className='pb-list-date-bougtht'> Date Sold: {transactions.dateOffer}</p>
            </div>
        </div>
    )
}
