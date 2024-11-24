import React, { useState } from 'react';
import './NavBar.css';
import RegistrationModal from '..//Modal/RegisterModal';
import LoginModal from '..//Modal/loginModal';
import useAuth from '../../hooks/useAuth';
import { toast, Toaster } from "react-hot-toast";
import { MdLogout } from "react-icons/md";
import { useNavigate } from "react-router-dom";

export default function Nav() {
    const { user, logout } = useAuth()

    const navigate = useNavigate()
    const [isModalOpen, setModalOpen] = useState(false);
    const [isModalOpenLogin, setModalOpenLogin] = useState(false);
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
        <div>
            <nav className="navbar">
                <div className="navbar-left">
                    <img src="/logo.png" alt="Logo" className="logo" /> {/* Replace with your logo path */}
                    <div className="brand">
                        <h1 className="brand-name">ABuild Homes</h1>
                        <span className="brand-tagline">ESTATES</span>
                    </div>
                </div>


                <div className="navbar-right">

                    {!user ? (
                        <>
                            <ul className="navbar-menu">
                                <li><a href="/" onClick={(e) => handleClick(e, "/")} className="menu-item active">Home</a></li>
                                <li><a href="/all-properties" onClick={(e) => handleClick(e, "/all-properties")} className="menu-item">All Properties</a></li>
                                <li><a href="/" className="menu-item" onClick={(e) => handleClick(e, "/")}>About Us</a></li>
                            </ul>
                            <button className="login-btn" onClick={() => setModalOpenLogin(true)}>Login</button>
                            <button className="login-btn" onClick={() => setModalOpen(true)}>Sign Up</button>
                        </>

                    ) : (
                        <>
                            <ul className="navbar-menu">
                                <li><a href="/" onClick={(e) => handleClick(e, "/")} className="menu-item active">Home</a></li>
                                <li><a href="/all-properties" onClick={(e) => handleClick(e, "/all-properties")} className="menu-item">All Properties</a></li>
                                <li><a href="/" onClick={(e) => handleClick(e, "/")} className="menu-item">About Us</a></li>
                                {user.role === "buyer" ?
                                    (
                                        <li><a href="/dashboard" onClick={(e) => handleClick(e, "/dashboard")} className="menu-item">Dashboard</a></li>
                                    ) : (
                                        <li><a href="/admin" onClick={(e) => handleClick(e, "/admin")} className="menu-item">Dashboard</a></li>
                                    )}
                                <li><img src={user.avatar || "/user.jpg"} alt="User Avatar" className="menu-avatar" /></li>
                            </ul>
                            <button className="login-btn" onClick={handleLogOut}>Log Out</button>
                        </>

                    )}
                </div>
            </nav>
            <Toaster />
            <LoginModal isOpen={isModalOpenLogin} onClose={() => setModalOpenLogin(false)} />
            <RegistrationModal isOpen={isModalOpen} onClose={() => setModalOpen(false)} />
        </div>

    );
}
