import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { IoLocationOutline } from "react-icons/io5";
import { MdMoney } from "react-icons/md";
import { useNavigate } from "react-router-dom";

import "./Property.css"

export default function Property() {
    const navigate = useNavigate()
    const [data, setData] = useState([]);
    const [page, setPage] = useState(1);
    const [totalPages, setTotalPages] = useState(1);
    const limit = 10;

    const fetchData = async (page) => {
        try {
            const response = await axios.get(`/properties?page=${page}&limit=${limit}`);
            console.log(response.data.properties)
            setData(response.data.properties)
            setTotalPages(response.data.totalPages)

        } catch (error) {
            console.error("Error fetching data:", error);
        }
    };

    useEffect(() => {
        fetchData(page);
    }, [page])

    const handleDetailsClick = (id) => {
        navigate(`/property-details/${id}`);
    };

    return (
        <>
            <div className='property-list-grid'>
                {data && data.map(p => (
                    <div className="property-list-card" key={p.id}>
                        <div className="property-list-verified"><span>{p.status}</span></div>
                        <div className="property-list-card-content">
                            <img src={p.image[0].url} alt="Property" className="property-list-property-image" />
                            <div className="property-list-details">
                                <h2>{p.title}</h2>
                                <p className="property-list-address"><span><IoLocationOutline className='property-list-local' /> </span>{p.location}</p>
                                <p className="property-list-price"> <MdMoney className='property-list-money' /> <p className='property-list-price-range'>Price Range: <span>{p.minPrice} $ - {p.maxPrice} $</span></p></p>
                                <div className="property-list-agent-info">
                                    <img src={p.agent.avatarAgent || "/user.jpg"} alt="Agent Profile" className="property-list-agent-photo" />
                                    <p>{p.agent.name}</p>
                                </div>
                                <button className="property-list-details-button" onClick={() => handleDetailsClick(p.id)} >Details</button>
                            </div>
                        </div>
                    </div>
                ))}
            </div>
            <div className='property-list-pagination'>
                <button onClick={() => setPage(page - 1)} disabled={page === 1}>
                    Pev
                </button>
                <span>{page} / {totalPages}</span>
                <button onClick={() => setPage(page + 1)} disabled={page === totalPages}>
                    Next
                </button>
            </div>
        </>


    )
}
