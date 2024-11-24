import React from 'react'

import { MdOutlineDashboard, MdOutlineRateReview, MdLocalOffer, MdLogout } from "react-icons/md";
import { LuTableProperties,LuGitPullRequest } from "react-icons/lu";
import { FaHome } from "react-icons/fa";
import { IoIosLogOut } from "react-icons/io";
import { SiGoogletagmanager } from "react-icons/si";
import { GiSellCard } from "react-icons/gi";
import { CgProfile } from "react-icons/cg";

import { toast, Toaster } from "react-hot-toast";
import { useNavigate } from "react-router-dom";
import useAuth from '../../hooks/useAuth';

import "./NavDashboard.css"


export default function NavDashboard() {
    const { logout } = useAuth()
    const navigate = useNavigate()

    const handleLogOut = () => {
        logout()

        toast("User has been logged out", {
            icon: <MdLogout />,
            style: {
                background: "#ff92b4",
            },
        })

        navigate('/')
    };

    const handleClick = (e, path) => {
        e.preventDefault();
        navigate(path);
    };
    
    return (
        <div className="nav-db-sidebar">
            <div className="nav-db-logo">
                <h2>Dashboard</h2>
            </div>
            <ul className="nav-db-nav-list">
                <li><a href="/dashboard" onClick={(e) => handleClick(e, "/dashboard")}><span className="nav-db-icon"><CgProfile /></span>Profile</a></li>
                <li><a href="/dashboard/wishList"  onClick={(e) => handleClick(e, "/dashboard/wishList")}><span className="nav-db-icon"><MdOutlineDashboard /></span>Wish List</a></li>
                <li><a href="/dashboard/propertiesBought"  onClick={(e) => handleClick(e, "/dashboard/propertiesBought")}><span className="nav-db-icon"><LuTableProperties /></span>Properties Bought</a></li>
                <li><a href="/dashboard/propertiesOffered" onClick={(e) => handleClick(e, "/dashboard/propertiesOffered")}><span className="nav-db-icon"><MdLocalOffer /></span>Properties Offered</a></li>
                <li><a href="/dashboard/propertiesManager" onClick={(e) => handleClick(e, "/dashboard/propertiesManager")}><span className="nav-db-icon"><SiGoogletagmanager /></span>Properties Manager</a></li>
                <li><a href="/dashboard/soldProperties" onClick={(e) => handleClick(e, "/dashboard/soldProperties")}><span className="nav-db-icon"><GiSellCard /></span>Sold Properties</a></li>
                <li><a href="/dashboard/requestedProperties" onClick={(e) => handleClick(e, "/dashboard/requestedProperties")}><span className="nav-db-icon"><LuGitPullRequest /></span>Requested Properties</a></li>
                <li><a href="/dashboard/reviews" onClick={(e) => handleClick(e, "/dashboard/reviews")}><span className="nav-db-icon"><MdOutlineRateReview /></span>Reviews</a></li>
                <li><hr /></li>
                <li><a href="/"><span className="nav-db-icon" onClick={(e) => handleClick(e, "/")}><FaHome /></span>Home</a></li>
            </ul>
            <div className="nav-db-logout">
                <a onClick={handleLogOut} ><span className="nav-db-icon"><IoIosLogOut /></span>Logout</a>
            </div>
            <Toaster />
        </div>
    )
}
