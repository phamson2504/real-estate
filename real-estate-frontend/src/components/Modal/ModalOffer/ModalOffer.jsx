import React, { useEffect, useState } from 'react'
import './ModalOffer.css'
import useAuth from '../../../hooks/useAuth';
import { useNavigate } from "react-router-dom";

import { toast, Toaster } from "react-hot-toast";
import { MdLogout } from "react-icons/md";

export default function ModalOffer({ isOpen, onClose, property }) {
    const { authAxios, user } = useAuth()

    const [offer, setOffer] = useState({
        propertyId: 0,
        buyerId: 0,
        sellerId: 0,
        amount: 0.0,
        dateOffer: ''
    });

    const checkUser = () => {
        if (isOpen && !user) {
            toast("Please login to make an offer", {
                icon: <MdLogout />,
                style: {
                    background: "#ff92b4",
                },
            });
            onClose();
            return false;
        }
        return true;
    };
    useEffect(() => {
        if (isOpen && !user) {
            toast("Please login to make an offer", {
                icon: <MdLogout />,
                style: {
                    background: "#ff92b4",
                },
            });
            onClose();
        }
    });

    useEffect(() => {
        const today = new Date().toISOString().split("T")[0];
        setOffer((prev) => ({ ...prev, dateOffer: today }));
        if(user) setOffer((prev) => ({ ...prev, propertyId: property.id, buyerId: user.id, sellerId: property.agent.id}))
    }, [user]);

    const handleInputChange = (e) => {
        const { name, value } = e.target;
        setOffer((prev) => ({
            ...prev,
            [name]: name === "amount" ? parseFloat(value) : value,
        }));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            if (checkUser) {
                await authAxios.post("/transaction", offer);
                alert("Offered successfully !!")
                onClose();
            }

        } catch (error) {
            console.error("Login error:", error);
        }
    };

    if (!isOpen) return null;
    return (
        <div className="md-offer-backdrop">
            <div className="md-offer-modal">
                <h2 className="md-offer-title">Make an Offer</h2>
                <form className="md-offer-form" onSubmit={handleSubmit}>
                    <div className="md-offer-group">
                        <label>Property Title</label>
                        <input type="text" value={property.title} disabled />
                    </div>

                    <div className="md-offer-group">
                        <label>Property Location</label>
                        <input
                            type="text"
                            value={property.location}
                            disabled
                        />
                    </div>
                    <div className="md-offer-row">
                        <div className="md-offer-group">
                            <label>Agent Name</label>
                            <input type="text" value={property.agent.name} disabled />
                        </div>
                        <div className="md-offer-group">
                            <label>Agent Email</label>
                            <input type="text" value={property.agent.email} disabled />
                        </div>
                    </div>

                    <div className="md-offer-group">
                        <label>
                            $ Offered Amount (
                            <span style={{ color: "green" }}>
                                Price range: ${property.minPrice} - ${property.maxPrice}
                            </span>
                            )
                        </label>
                        <input type="number"
                            name='amount'
                            placeholder="Enter your offer amount"
                            value={offer.amount}
                            onChange={handleInputChange}
                            required
                        />
                    </div>

                    <div className="md-offer-row">
                        <div className="md-offer-group">
                            <label>Buyer Name</label>
                            <input type="text" value={user && (user.username)} disabled />
                        </div>
                        <div className="md-offer-group">
                            <label>Buyer Email</label>
                            <input type="email" value={user && (user.email)} disabled />
                        </div>
                    </div>

                    <div className="md-offer-group">
                        <label>Buying Date</label>
                        <input type="date"
                            name="dateOffer"
                            value={offer.dateOffer}
                            onChange={handleInputChange}
                            required
                        />
                    </div>

                    <div className="md-offer-actions">
                        <button type="button" onClick={onClose} className="md-offer-cancel">
                            Cancel
                        </button>
                        <button type="submit" className="md-offer-submit">
                            Make Offer
                        </button>
                    </div>
                </form>
            </div>
            <Toaster />
        </div>
    )
}
