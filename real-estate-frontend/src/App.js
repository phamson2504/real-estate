import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import UserRouter from './routes/UserRoute';
import AdminRoutes from './routes/AdminRoutes';
import DashBoardRoutes from './routes/DashBoardRoutes';

function App() {
  return (
    <Router future={{ v7_startTransition: true, v7_relativeSplatPath: true }}>
      <Routes>
        <Route path="/*" element={<UserRouter />} />
        <Route path="/admin/*" element={<AdminRoutes />} />
        <Route path='/dashboard/*' element={<DashBoardRoutes />} />
      </Routes>
    </Router>
  );
}

export default App;
