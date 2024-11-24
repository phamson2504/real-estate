import React from 'react';
import { Routes, Route } from 'react-router-dom';
import Home from '../pages/home/Home';
import MainLayout from '../layouts/MainLayout';
import AllProperties from '../pages/allProperties/AllProperties';
import PropertyDetails from '../pages/propertyDetail/PropertyDetails';

const UserRouter = () => {
    return (
        <Routes>
            <Route element={<MainLayout />}>
                <Route path="/" element={<Home />} />
                <Route path="/all-properties"  element={<AllProperties />} />
                <Route path="/property-details/:id"  element={<PropertyDetails />} />
            </Route>
        </Routes>
    );
};

export default UserRouter;