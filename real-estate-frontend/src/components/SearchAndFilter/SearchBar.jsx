import React from 'react';
import { HiAdjustmentsHorizontal } from "react-icons/hi2";
import { FaAngleDown } from "react-icons/fa";
import './SearchBar.css';

const SearchBar = () => {
    return (
        <div className="search-bar">
            <div className="location">
                <span>Hồ Chí Minh</span>
                <FaAngleDown />
            </div>
            <input
                type="text"
                className="search-input"
                placeholder="Enter up to 5 locations."
            />
            <button className="search-button">Search</button>
            <div className="filters">
                <button>Price</button>
                <button>Property type</button>
                <button>Bedrooms</button>
                <button> <HiAdjustmentsHorizontal /> Filters </button>
            </div>
        </div>
    );
};

export default SearchBar;
