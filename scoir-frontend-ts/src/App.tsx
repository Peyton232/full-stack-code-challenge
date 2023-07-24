import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'; // Use 'Routes' instead of 'Switch'
import SignIn from './pages/signIn';
import Dashboard from './pages/dashboard'; 

function App() {
  return (
    <Router>
      <Routes> 
        <Route path="/" element={<SignIn />} />
        <Route path="/dashboard" element={<Dashboard />} />
      </Routes>
    </Router>
  );
}

export default App;
