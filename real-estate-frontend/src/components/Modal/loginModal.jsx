import React, { useState } from 'react';
import useAuth from "../../hooks/useAuth";
import "./RegistrationModal.css"

const LoginModal = ({ isOpen, onClose }) => {

    const [error, setError] = useState("");
    
    const {login} = useAuth()

    const [user, setUser] = useState({
        username: '',
        password: '',
    });

    const handleInputChange = (e) => {
        const { name, value } = e.target;
        setUser({ ...user, [name]: value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            await login(user);
            alert("Login successfully !!")
            onClose();
        } catch (error) {
            console.error("Login error:", error);
            setError("Failed to log in. Please check your credentials.");
        }
    };

    if (!isOpen) return null;

    const handleOverlayClick = (e) => {
        if (e.target.className === "modal-overlay") {
            onClose();
        }
    };
    return (
        <div className="modal-overlay" onClick={handleOverlayClick}>
            <div className="modal-content">
                <button onClick={onClose} className="close-button">X</button>
                <h2>Login</h2>
                {error && <p>{error}</p>}
                <form onSubmit={handleSubmit}>
                    <div>
                        <label>Username:</label>
                        <input
                            type="text"
                            name="username"
                            value={user.username}
                            onChange={handleInputChange}
                            required
                        />
                    </div>
                    <div>
                        <label>Password:</label>
                        <input
                            type="password"
                            name="password"
                            value={user.password}
                            onChange={handleInputChange}
                            required
                        />
                    </div>
                    <button type="submit">Login</button>
                </form>
            </div>
        </div>
    );
};

export default LoginModal;
