import React, { useEffect, useState } from 'react'
import PropertySold from '../../components/Property/PropertySold/PropertySold'
import useAuth from '../../hooks/useAuth';

export default function PropertiesSold() {
  const { user, authAxios } = useAuth()
  const [data, setData] = useState([]);

  const fetchData = async () => {
    try {
      const response = await authAxios.get(`/transaction/getTransactionSold?id=${user.id}`);
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
    <div className='ps-contain-layout'>
      <h2 className='ps-tilte'>Property Sold</h2>
      <div className='ps-contain-content'>
        {data && data.map(t => (
          <PropertySold transactions={t} />
        ))}
      </div>
    </div>
  )
}
