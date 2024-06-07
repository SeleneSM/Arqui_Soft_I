import React from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import LoginRegister from './Login'
/*import Reserve from './ReservaCursos'
import Admin from './Admin.jsx'*/
import './main.css'


function App(){
  return (
      <Router>
        <Routes>
          <Route path= "/login" element={<Login/>}/>          
          <Route path="/reservacursos/:token/:user_id" element={<ReservaCursos />} />
          <Route path="/admin/:token/:user_id" element={<Admin />} />
        </Routes>
      </Router>
    );
  }
  
  export default App;