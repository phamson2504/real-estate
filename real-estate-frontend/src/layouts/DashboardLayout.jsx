import React from 'react'
import NavDashboard from"../components/Navbar/NavDashboard"
import { Outlet } from 'react-router-dom';
import "./DashboardLayout.css"

export default function DashboardLayout() {
  return (
    <div className="db-home">
      <NavDashboard />
      <div className="db-content">
        <Outlet />
      </div>
    </div>
  )
}
