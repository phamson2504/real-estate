import React from 'react';
import { Outlet } from 'react-router-dom';
import NavBar from '../components/Navbar/NavBar';
import Footer from '../components/Footer/Footer';

const MainLayout = () => {
  return (
    <>
      <NavBar />
      <Outlet />
      <Footer /> 
    </>
  );
};

export default MainLayout;