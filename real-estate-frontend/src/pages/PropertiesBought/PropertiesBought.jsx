import React, { useEffect, useState } from 'react'
import PropertyBought from '../../components/Property/PropertyBought/PropertyBought'
import "./PropertiesBought.css"
import useAuth from '../../hooks/useAuth';

export default function PropertiesBought() {
  const { user, authAxios } = useAuth()
  const [data, setData] = useState([]);

  const fetchData = async () => {
    try {
      const response = await authAxios.get(`/transaction/getTransactionBougth?id=${user.id}`);
      setData(response.data.transactions)
      console.log(response.data.transactions)
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  };

  useEffect(() => {
    fetchData();
  }, [])
  return (
    <div className='pb-contain-layout'>
      <h2 className='pb-tilte'>Property Bought</h2>
      <div className='pb-contain-content'>
        {data && data.map(t => (
          <PropertyBought transactions ={t}/>
        ))}
      </div>
    </div>
  )
}
