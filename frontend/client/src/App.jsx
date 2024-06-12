import React, { useState, useEffect } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import BotonLogin from './Componentes/BotonLogin';
import Cursos from './Componentes/Cursos.jsx';
import Inscripciones from './Componentes/Inscripciones.jsx';



function App(){

  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [isAdmin, setIsAdmin] = useState(false);
  const [cursos, setCursos] = useState([]);
  const [userId, setUserId] = useState(null);
  const [mostrarInscripciones, setMostrarInscripciones] = useState(false);
  const [inscripciones, setInscripciones] = useState([]);
  /*const [nuevoCurso, setNuevoCurso] = useState({
    materia: '',
    fechaInicio: '',
    fechaFin: ''
    });*/
  const [inscripcionesTotales, setInscripcionesTotales] = useState([]);
/*
  const handleNuevoCursoChange = (event) => {
    const {materia, fechaInicio, fechaFin} = event.target;
    setNuevoCurso((prevNuevoCurso) => ({
      ...prevNuevoCurso,
      [materia]: value
    }));
  };

  const crearNuevoCurso = () => {
    const fechaInicio = date(nuevoCurso.fechaInicio);
    const fechaFin = date(nuevoCurso.fechaFin);

    fetch('http://localhost:8090/cursos', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        materia: nuevoCurso.materia,
        fechaInicio: nuevoCurso.fechaInicio,
        fechaFin: nuevoCurso.fechaFin
      })
    })
      .then((response) => response.json())
      .then((data) => {
        if (data.error) {
          console.error(data.error);
        } else {
          setCursos((prevCursos) => [...prevCursos, data]);
          alert('Se ha creado un nuevo curso con éxito');
          setNuevoCurso({
            materia: '',
            fechaInicio: '',
            fechaFin: ''
          });
        }
      })
      .catch((error) => console.error(error));
  };
*/
  const handleLogin = (rol, userId) => {
    setIsLoggedIn(true);
    setIsAdmin(rol === 1);
    setUserId(userId);
  };


  const toggleInscripciones = () => {

    if (isAdmin) {
      fetch('http://localhost:8090/inscripciones', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json'
        }
      })
        .then((response) => response.json())
          .then((data) => {
            console.log(data); // Verificar los datos recibidos
            setInscripcionesTotales(data); // Asignar los datos a la variable de estado
          })
        .catch((error) => console.error(error));
    }else{
      fetch(`http://localhost:8090/inscripciones/inscripcionuser/${userId}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json'
        }
      })
          .then((response) => response.json())
          .then((data) => setInscripciones(data))
          .catch((error) => console.error(error));
    }
      setMostrarInscripciones(!mostrarInscripciones);
  };


  useEffect(() => {
    fetch('http://localhost:8090/cursos')
      .then((response) => response.json())
      .then((data) => setCursos(data))
      .catch((error) => console.error(error));

    if (isLoggedIn) {
      fetch(`http://localhost:8090/inscripciones/inscripcionuser/${userId}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json'
        }
      })
        .then((response) => response.json())
        .then((data) => setInscripciones(data))
        .catch((error) => console.error(error));
    }
  });
/*
    if (isAdmin) {
      fetch('http://localhost:8090/inscripciones', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json'
        }
      })
        .then((response) => response.json())
        .then((data) => setInscripcionesTotales(data))
        .catch((error) => console.error(error));
    }
  }, [isLoggedIn, userId, isAdmin]);
*/
  return (
    <div className="App">
      <div>
        <h1>ENCONTRA EL MEJOR CURSO PARA VOS</h1>
        {!isLoggedIn && <BotonLogin handleLogin={handleLogin} />}
        {isLoggedIn && (
          <div>
            <button onClick={toggleInscripciones}>Inscripciones</button>
            {mostrarInscripciones && (
              <Inscripciones inscripciones={inscripciones} inscripcionesTotales={inscripcionesTotales} />
            )}
          </div>
        )}
      </div>
    </div>
  );
}
  export default App;

  /*{isAdmin && (
          <div>
            <h2>Agregar nuevo curso</h2>
            <input
              type="text"
              name="materia"
              value={nuevoCurso.materia}
              onChange={handleNuevoCursoChange}
              placeholder="Nombre de la materia del curso"
            />
            <input
              type="date"
              name="fechainicio"
              date={nuevoCurso.fechaInicio}
              onChange={handleNuevoCursoChange}
              placeholder="Fecha inicio"
            />
            <input
              type="date"
              name="fechafin"
              date={nuevoCurso.fechaFin}
              onChange={handleNuevoCursoChange}
              placeholder="Fecha fin"
            />
            <button onClick={crearNuevoCurso}>Crear Curso</button>
            <div>
              <h2>Listado de Cursos</h2>
              {cursos.length > 0 ? (
                <ul>
                  {cursos.map((curso) => (
                    <li key={curso.id}>
                      <p>Materia: {curso.materia}</p>
                      <p>Fecha Inicio: {curso.fechaInicio}</p>
                      <p>Fecha Fin: {curso.fechaFin}</p>
                    </li>
                  ))}
                </ul>
              ) : (
                <p>No se han creado curso todavía.</p>
              )}
            </div>
          </div>
        )}
        {!isAdmin &&
          cursos.map((curso) => (
            <Cursos
              key={curso.id}
              nombreMateria={curso.materia}
              fechaInicio={curso.fechaInicio}
              fechaFin={curso.fechaFin}
              isLoggedIn={isLoggedIn}
              cursoId={curso.id}
              userId={userId}
            />
          ))}*/ 