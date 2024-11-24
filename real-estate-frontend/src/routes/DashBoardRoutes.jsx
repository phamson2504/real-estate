import React from 'react'
import { Routes, Route } from 'react-router-dom';
import DashboardLayout from '../layouts/DashboardLayout';
import Profile from '../pages/profile/Profile';
import WishList from '../pages/wishList/WishList';

import Review from '../pages/review/Review';
import PropertiesBought from '../pages/PropertiesBought/PropertiesBought';
import RequestedProperties from '../pages/requestedProperties/RequestedProperties';
import PropertiesOffered from '../pages/propertiesOffered/PropertiesOffered';
import PropertiesSold from '../pages/propertySold/PropertiesSold';
import PropertiesManager from '../pages/propertiesManager/PropertiesManager';
import PropertyCreate from '../pages/PropertyCreate/PropertyCreate';

export default function DashBoardRoutes() {
    return (
        <Routes>
            <Route element={<DashboardLayout />}>
                <Route path='/' element={<Profile />} />
                <Route path='/wishList' element={<WishList />} />
                <Route path='/propertiesBought' element={<PropertiesBought />} />
                <Route path='/requestedProperties' element={<RequestedProperties />} />
                <Route path='/propertiesOffered' element={<PropertiesOffered />} />
                <Route path='/soldProperties' element={<PropertiesSold />} />
                <Route path='/propertiesManager' element={<PropertiesManager />} />\
                <Route path='/addProperties' element={<PropertyCreate />} />
                <Route path='/reviews' element={<Review />} />
            </Route>
        </Routes>
    )
}
