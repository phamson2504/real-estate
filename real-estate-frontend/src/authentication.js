import axios from 'axios';
import {jwtDecode} from 'jwt-decode';

const TOKEN_KEY = 'token';

// Function to set the token in local storage
export const setToken = (token) => {
    localStorage.setItem(TOKEN_KEY, token);
};

// Function to get the token from local storage
export const getToken = () => {
    return localStorage.getItem(TOKEN_KEY);
};

// Function to remove the token from local storage (for logout)
export const removeToken = () => {
    localStorage.removeItem(TOKEN_KEY);
};

// Function to check if the token is expired
const isTokenExpired = (token) => {
    try {
        const { exp } = jwtDecode(token);
        if (exp * 1000 < Date.now()) {
            return true; // Token is expired
        }
        return false;
    } catch (error) {
        console.error("Failed to decode token", error);
        return true; // Treat as expired if decode fails
    }
};

// Function to check if user is authenticated
export const isAuthenticated = () => {
    const token = getToken();
    if (!token || isTokenExpired(token)) {
        removeToken(); // Clear the token if expired
        return false;
    }
    return true;
};

// Login function to authenticate user and store token
export const login = async (user) => {
    try {
        const response = await axios.post('/authentication/login', user);
        setToken(response.data.data.token);
        return response.data.data;
    } catch (error) {
        console.error("Login failed", error);
        throw error;
    }
};

// Logout function to remove the token and clear session
export const logout = () => {
    removeToken();
    window.location.reload();
};

// Helper to attach token to request headers (for axios)
export const authAxios = axios.create();

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
