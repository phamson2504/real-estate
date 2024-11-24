import React, { createContext, useContext, useState, useEffect } from 'react';
import axios from 'axios';
import { jwtDecode } from 'jwt-decode';

export const AuthContext = createContext();

const TOKEN_KEY = 'token';

const getToken = () => localStorage.getItem(TOKEN_KEY);
const setToken = (token) => localStorage.setItem(TOKEN_KEY, token);
const removeToken = () => localStorage.removeItem(TOKEN_KEY);

const isTokenExpired = (token) => {
    try {
        const { exp } = jwtDecode(token);
        return exp * 1000 < Date.now();
    } catch (error) {
        console.error("Failed to decode token", error);
        return true;
    }
};

export const AuthProvider = ({ children }) => {
    const [user, setUser] = useState(null);

    const login = async (userData) => {
        try {
            const response = await axios.post('/authentication/login', userData);
            const token = response.data.data.token;
            setToken(token);

            const userResponse = await axios.get('/authentication/user', {
                headers: { Authorization: `Bearer ${token}` }
            });
            setUser(userResponse.data.data);
            console.log(userResponse.data.data)
            return response.data.data;
        } catch (error) {
            console.error("Login failed", error);
            throw error;
        }
    };
    useEffect(() => {
        const token = getToken();
        if (token && !isTokenExpired(token) && !user) {
            // Nếu có token nhưng user chưa được thiết lập, khôi phục user từ API hoặc từ localStorage
            const fetchUser = async () => {
                try {
                    const response = await axios.get('/authentication/user', {
                        headers: { Authorization: `Bearer ${token}` }
                    });
                    setUser(response.data.data);
                } catch (error) {
                    console.error("Failed to fetch user", error);
                    logout();
                }
            };
            fetchUser();
        }
    }, [user]);

    const logout = () => {
        removeToken();
        setUser(null);
        return
    };

    const authAxios = axios.create();
    authAxios.interceptors.request.use(
        (config) => {
            const token = getToken();
            if (token && !isTokenExpired(token)) {
                config.headers.Authorization = `Bearer ${token}`;
            }
            return config;
        },
        (error) => Promise.reject(error)
    );

    return (
        <AuthContext.Provider value={{ user, login, logout, authAxios }}>
            {children}
        </AuthContext.Provider>
    );
};

export const useAuth = () => useContext(AuthContext);
