import React from 'react'
import "./Footer.css"
import { FaArrowAltCircleRight, FaFacebook, FaInstagram, FaLinkedin, FaTwitter } from "react-icons/fa";

export default function Footer() {
    return (
        <footer className="footer">
            <div className="footer-section company-info">
                <img src="/logo.png" alt="ABuild Homes Estates Logo" className="footer-logo" />
                <p>Where your dreams find their address, and every door opens to endless possibilities.</p>
                <p>Mail: <a href="mailto:abuild@estate.com">abuild@estate.com</a></p>
                <p>Phone: +880 1622 3121</p>
                <div className="social-icons">
                    <FaFacebook style={{ color: "#3b5998" }} />
                    <FaInstagram style={{ color: "#e4405f" }} />
                    <FaLinkedin style={{ color: "#0077b5" }} />
                    <FaTwitter style={{ color: "#1da1f2" }} />
                </div>
            </div>

            <div className="footer-section">
                <h3>OUR COMPANY</h3>
                <ul>
                    <li><a href="#">About us</a></li>
                    <li><a href="#">Contact</a></li>
                    <li><a href="#">Jobs</a></li>
                    <li><a href="#">Press kit</a></li>
                </ul>
            </div>

            <div className="footer-section">
                <h3>LEGAL</h3>
                <ul>
                    <li><a href="#">Terms of use</a></li>
                    <li><a href="#">Privacy policy</a></li>
                    <li><a href="#">Cookie policy</a></li>
                </ul>
            </div>

            <div className="footer-section subscribe">
                <h3>SUBSCRIBE</h3>
                <p>Enter your email address</p>
                <div className="subscribe-form">
                    <input type="email" placeholder="username@email.com" />
                    <button type="submit"><FaArrowAltCircleRight /></button>
                </div>
            </div>
        </footer>
    )
}
