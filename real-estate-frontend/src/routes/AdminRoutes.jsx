import React from 'react'
import { Routes, Route } from 'react-router-dom';
import AdminLayout from '../layouts/AdminLayout';
import Admin from '../pages/admin/Admin';

export default function AdminRoutes() {
    return (
        <Routes>
            <Route element={<AdminLayout />}>
                <Route path="/" element={<Admin />} />
            </Route>
        </Routes>
    )
}
