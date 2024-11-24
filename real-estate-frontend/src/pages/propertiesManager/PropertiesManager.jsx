import React, { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom';
import PropertyManager from '../../components/Property/PropertyManager/PropertyManager'
import axios from 'axios';
import useAuth from '../../hooks/useAuth';
import './PropertiesManager.css'

export default function PropertiesManager() {
  const navigate = useNavigate();
  const { user } = useAuth()
  
  const [data, setData] = useState([]);

  const fetchData = async () => {
    try {
      const response = await axios.get(`/properties/properties-by-agent?id=${user.agent.id}`);
      setData(response.data.properties)
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  };
  useEffect(() => {
    fetchData();
  }, [])

  return (
    <div>
      <div className='pm-header'>
        <h2 className='pm-tilte'>Properties Manager</h2>
        <button className='pm-button-add' onClick={() => navigate('/dashboard/addProperties')}>Add Property</button>
      </div>
      <div className='pm-contain-content'>
        {data && data.map(p => (
          <PropertyManager properties={p} />
        ))}
      </div>

    </div>
  )
}
