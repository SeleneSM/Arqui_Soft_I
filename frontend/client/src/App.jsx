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


async function fetchMaterias(cursos) {
  let endpointBase = "http://host.docker.internal:8090/materia";
  try {
    // Mapeamos los cursos a sus respectivos fetch basados en `materia_id`
    const fetchPromises = cursos.map(curso =>
      fetch(`${endpointBase}/${curso.materia_id}`)
        .then(response => {
          if (!response.ok) {
            throw new Error(`Error en materia_id ${curso.materia_id}`);
          }
          return response.json(); // Parseamos la respuesta a JSON
        })
        .then(materia => ({ ...curso, materia })) // Agregamos la materia al curso
    );

    // Esperamos a que todos los fetch se completen
    const cursosConMaterias = await Promise.all(fetchPromises);
    return cursosConMaterias;

  } catch (error) {
    console.error("Error al obtener materias:", error);
    throw error;
  }
}


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
    // Asegurarse de que el token se guarde en localStorage durante el login
    // Esto debería estar en tu componente de login
  };


  
  const toggleInscripciones = () => {
    if (isAdmin) {/*

      */
    }
    setMostrarInscripciones(!mostrarInscripciones);
  };

  useEffect(() => {
    const obtenerCursos = async () => {
      try {
        const response = await fetch('http://host.docker.internal:8090/cursos');
        const data = await response.json();
        const cs = await fetchMaterias(data);
        setCursos(cs);
      } catch (error) {
        console.error("Error al obtener cursos o materias:", error);
      }
    }

    obtenerCursos();
  }, []);

  return (
    <div className="App">
      <div>
        <h1>CURSIFY - Encontra el mejor curso para vos</h1>
        {isLoggedIn && isAdmin  && <h1>Bienvenido administrador</h1> && <BotonCrearCurso />}
        {!isLoggedIn && <BotonLogin handleLogin={handleLogin} />}
        {!isLoggedIn && <BotonRegister />}
        {!isAdmin && isLoggedIn && (
          <div>
            <div className="boton-insc">
              <button onClick={toggleInscripciones}>Mis Inscripciones</button>
            </div>
            {mostrarInscripciones && (
              <Inscripciones 
                userId={userId}
                // No necesitamos pasar el token por ahora
              />
            )}
          </div>
        )}
        {/*!isAdmin &&
          cursos.map((curso) => (
            <Cursos
              key={curso.id}
              curso={curso}
              isLoggedIn={isLoggedIn}
              userId={userId}
            />
          ))*/}
      </div>
    </div>
  );
}
export default App;

