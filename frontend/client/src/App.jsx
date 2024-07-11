import React, { useState, useEffect } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter as Router, Routes, Route, Form } from 
'react-router-dom'
import BotonLogin from './Componentes/BotonLogin';
import BotonRegister from './Componentes/BotonRegister';
import BotonCrearCurso from './Componentes/BotonCrearCurso';
import FormularioRegister from './Componentes/FormularioRegister.jsx';
import Cursos from './Componentes/Cursos.jsx';
import Inscripciones from './Componentes/Inscripciones.jsx';

function App() {

  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [isAdmin, setIsAdmin] = useState(false);
  const [cursos, setCursos] = useState([]);
  const [userId, setUserId] = useState(null);
  const [mostrarInscripciones, setMostrarInscripciones] = useState(false);
  const [inscripciones, setInscripciones] = useState([]);
  const [inscripcionesTotales, setInscripcionesTotales] = useState([]);
  

  const handleLogin = (rol, userId) => {
    setIsLoggedIn(true);
    setIsAdmin(rol === "Administrador");
    setUserId(userId);
  };


  
  const toggleInscripciones = () => {
    if (isAdmin) {/*
      
      */
    }
    setMostrarInscripciones(!mostrarInscripciones);
  };

  useEffect(() => {
    fetch('http://localhost:8090/cursos')
      .then((response) => response.json())
      .then((data) => setCursos(data))
      .catch((error) => console.error(error));

    if (isLoggedIn) {
      fetch(`http://localhost:8090/inscripciones_por_usuario/${userId}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json'
        }
      })
        .then((response) => response.json())
        .then((data) => setInscripciones(data))
        .catch((error) => console.error(error));
    }
  }, [isLoggedIn, userId]);

  return (
    <div className="App">
      <div>
        <h1>CURSIFY - Encontra el mejor curso para vos</h1>
        {isLoggedIn && isAdmin  && <h1>Bienvenido administrador</h1>&&
         <BotonCrearCurso />}
        {!isLoggedIn && <BotonLogin handleLogin={handleLogin} />}
        {!isLoggedIn && <BotonRegister />}
        {!isAdmin && isLoggedIn && (
          <div>
            <div className="boton-insc">
            <button onClick={toggleInscripciones}>Mis Inscripciones</button>
            </div>
            {mostrarInscripciones && (
              <Inscripciones inscripciones={inscripciones} inscripcionesTotales={inscripcionesTotales} />
            )}
          </div>
        )}

        {!isAdmin &&
          cursos.map((curso) => (
            <Cursos
              key={curso.id}
              curso={curso}
              isLoggedIn={isLoggedIn}
              userId={userId}
            />
          ))}
      </div>
    </div>
  );
}
export default App;

