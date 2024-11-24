import React, { useEffect, useState } from 'react'
import PropertyOffered from '../../components/Property/PropertyOffered/PropertyOffered'
import useAuth from '../../hooks/useAuth';
import "./PropertiesOffered.css"

export default function PropertiesOffered() {
  const { user, authAxios } = useAuth()
  const [data, setData] = useState([]);

  const fetchData = async () => {
    try {
      const response = await authAxios.get(`/transaction/get-properties-offered?id=${user.id}`);
      setData(response.data.transactions)
      console.log(data)
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  };

  useEffect(() => {
    fetchData();
  }, [])

  return (
    <div className='po-contain-layout'>
      <h2 className='po-tilte'>Properties Offered</h2>
      <div className='po-contain-content'>
        {data && data.map(t => (
          <PropertyOffered transaction={t} />
        ))}
      </div>

    </div>
  )
}
