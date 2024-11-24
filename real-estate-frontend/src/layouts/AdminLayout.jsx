import React from 'react'
import { Outlet } from 'react-router-dom';
import "./AdminLayout.css"
import NavAdmin from '../components/Navbar/NavAdmin';

export default function AdminLayout() {
  return (
    <div className="ad-home">
      <NavAdmin />
      <div className="ad-content">
        <Outlet />
      </div>
    </div>
  )
}
