import React, { useState, useEffect }  from 'react'
import PropertyDetail from '../../components/Property/PropertyDetail/PropertyDetail'
import AgentContact from '../../components/AgentContact/AgentContact'
import axios from 'axios';
import { useParams } from "react-router-dom";
import useAuth from '../../hooks/useAuth';
import './PropertyDetails.css'

export default function PropertyDetails() {
  const { user, authAxios } = useAuth()
  const { id } = useParams();
  const [property, setProperty] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [checkFavorite, setCheckFavorite] = useState(false); 

  useEffect(() => {
    const fetchProperty = async () => {
      try {
        const response = await axios.get(`/properties/property-details?id=${id}`);
        if (response.status !== 200) {
          throw new Error("Failed to fetch property details");
        }
        setProperty(response.data.property);
        if(user){
          const checkFavorite = await authAxios.get(`/properties/checkPropertyFavorite?propertyId=${id}`);
          setCheckFavorite(checkFavorite.data.checkIsFavorite)
        }
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchProperty();
  }, [id, user]);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>Error: {error}</div>;

  return (
    <div className='pd-contain'>
      {property && <PropertyDetail property={property} checkFavorite={checkFavorite}/>}
      {property &&<AgentContact agent={property.agent} />}
    </div>
  )
}
