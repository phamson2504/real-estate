import React, { useState, useEffect } from 'react'
import PropertyWishList from '../../components/Property/PropertyWishList/PropertyWishList'
import './WishList.css'
import useAuth from '../../hooks/useAuth';

export default function WishList() {
  const { user, authAxios } = useAuth()
  const [properties, setProperties] = useState([]);

  useEffect(() => {
    const fetchProperty = async () => {
      try {
        if (user) {
          const properties = await authAxios.get(`/properties/propertyFavoriteByUserId`);
          setProperties(properties.data.properties)
          console.log(properties.data.properties)
        }
      } catch (err) {
        console.log(err)
      }
    };

    fetchProperty();
  }, [user]);

  return (
    <div className='wl-contain-layout'>
      <h2 className='wl-tilte'>Property Wish List</h2>
      <div className='wl-contain-content'>
      {properties && properties.map(p => (
        <PropertyWishList property = {p}/>
      ))}
      </div>
    </div>

  )
}
