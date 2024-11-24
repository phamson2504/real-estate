import React, { useState } from 'react'
import { MdOutlineAddTask } from "react-icons/md";
import { IoLocationOutline } from "react-icons/io5";
import { MdMoney, MdOutlineBedroomChild, MdOutlineRealEstateAgent } from "react-icons/md";
import { IoIosArrowDropleft, IoIosArrowDropright } from 'react-icons/io';
import { FaBath } from "react-icons/fa";
import "./PropertyDetail.css";
import { useSwipeable } from 'react-swipeable';
import ModalOffer from '../../Modal/ModalOffer/ModalOffer';
import useAuth from '../../../hooks/useAuth';
import { toast, Toaster } from "react-hot-toast";
import { MdLogout } from "react-icons/md";

export default function PropertyDetail({ property, checkFavorite }) {
  console.log(checkFavorite)
  const { authAxios, user } = useAuth()

  const [isModalOpen, setModalOpen] = useState(false);
  const [currentImage, setCurrentImage] = useState(0);

  const handlers = useSwipeable({
    onSwipedLeft: () => setCurrentImage((prev) => (prev + 1) % property.image.length),
    onSwipedRight: () => setCurrentImage((prev) => (prev - 1 + property.image.length) % property.image.length),
  });

  const handleNext = () => {
    setCurrentImage((prev) => (prev + 1) % property.image.length);
  };

  const handlePrev = () => {
    setCurrentImage((prev) => (prev - 1 + property.image.length) % property.image.length);
  };

  async function addToWishList() {
    try {
      if (!user) {
        toast("Please login to add", {
          icon: <MdLogout />,
          style: {
            background: "#ff92b4",
          },
        });
        return false;
      }
      await authAxios.get(`/properties/GetPropertyFavoreat?propertyId=${property.id}`);
      alert("Added to with list successfully !!")
    } catch (error) {
      console.error("Login error:", error);
    }
  }

  return (
    <div className='pdl-contain'>
      <Toaster />
      <div className='pdl-header'>
        <h2>{property.title}</h2>
        <button
          className="pdl-btn-wish-list"
          onClick={addToWishList}
          disabled={checkFavorite}
        >
          <MdOutlineAddTask />
          <p>{checkFavorite ? "Already in Wishlist" : "Add to Wishlist"}</p>
        </button>
      </div>
      <div {...handlers} className="pdl-swipeable-container">
        <img
          src={property.image[currentImage].url}
          alt={`slide-${currentImage}`}
          className="pdl-swipeable-image"
        />
        {/* Left Arrow */}
        <button onClick={handlePrev} className="pdl-swipeable-arrow left">
          <IoIosArrowDropleft />
        </button>

        {/* Right Arrow */}
        <button onClick={handleNext} className="pdl-swipeable-arrow right">
          <IoIosArrowDropright />
        </button>
      </div>
      <div className='pdl-details-content'>
        <p className="pdl-address"><span><IoLocationOutline className='pdl-icon-local' /> </span>{property.location}</p>
        <p className="pdl-price"> <MdMoney className='pdl-icon-money' /> <span className='pdl-price-range'><span>{property.minPrice} $ - {property.maxPrice} $</span></span ></p>
        <p className="pdl-verifed"><span><IoLocationOutline className='pdl-icon-verifed' /></span>{property.status}</p>
        <p className='pdl-details-title'>Details:</p>
        <hr />
        <p className='pdl-details-description'>{property.description}</p>
        <div className='pdl-bottom'>
          <div className='pdl-item-icons'>
            <p><FaBath className='pdl-icons' /> <span>Bath Room: {property.bathrooms}</span> </p>
            <p><MdOutlineBedroomChild className='pdl-icons' /> <span>Bed Rooms: {property.bedrooms} </span></p>
            <p><MdOutlineRealEstateAgent className='pdl-icons' /> <span>Size: {property.square_feet} sq.ft</span></p>
          </div>
          <button onClick={() => setModalOpen(true)}>Make offer</button>
        </div>
      </div>
      <ModalOffer isOpen={isModalOpen} onClose={() => setModalOpen(false)} property={property} />
    </div>
  )
}
