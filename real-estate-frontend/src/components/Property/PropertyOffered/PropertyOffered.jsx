import React,{useState} from 'react'
import './PropertyOffered.css'
import ModalOffer from '../../Modal/ModalOffer/ModalOffer';

export default function PropertyOffered({transaction}) {
    const [isModalOpen, setModalOpen] = useState(false);
    console.log(transaction)
    return (
        <div className='po-list-card' key={transaction.id}>
            <img src='/PropertiesImage/property1.jpg' alt="Amage property" />
            <div className='po-list-detail'>
                <h2>{transaction.Properties.title}</h2>
                <p className='po-list-address'>{transaction.Properties.location}</p>
                <p className='po-list-price'> Price Range: {transaction.Properties.minPrice}$ - {transaction.Properties.maxPrice}$</p>
                <p className='po-list-price-offer'> Price Offered: {transaction.amount}$</p>
                <div className='po-list-agent-info'>
                    <img src= {transaction.Properties.agent.avatarAgent} alt='Amage agent' />
                    <p> {transaction.Properties.agent.name}</p>
                </div>
                <div className='po-list-action'>
                    <p> Date Offered: {transaction.dateOffer}</p>
                    <button onClick={() => setModalOpen(true)}>Re-offer</button>
                </div>
            </div>
            <ModalOffer isOpen={isModalOpen} onClose={() => setModalOpen(false)} property={transaction.Properties} />
        </div>
    )
}
