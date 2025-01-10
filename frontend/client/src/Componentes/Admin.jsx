import React, { useState, useEffect } from "react";
import '../Stylesheet/Admin.css';
import Header from '../Componentes/Header';
import { useParams } from 'react-router-dom';

function Admin() {
  const [cursos, setCursos] = useState([]);
  const [nuevoCurso, setNuevoCurso] = useState({
    fecha_Inicio: '',
    fecha_Fin: '',
    requisitos: '',
    instructor: '',
    materia_id: null,
  });
  const [materias, setMaterias] = useState([]);
  const { token } = useParams();

  // Fetch materias and cursos
  useEffect(() => {
    // Fetch materias list
    fetch('http://host.docker.internal:8090/materias')
      .then((response) => response.json())
      .then((data) => setMaterias(data))
      .catch((error) => console.error(error));

    // Fetch cursos list
    fetch('http://host.docker.internal:8090/cursos')
      .then((response) => response.json())
      .then((data) => setCursos(data))
      .catch((error) => console.error(error));
  }, []);

  const crearNuevoCurso = async (event) => {
    event.preventDefault();

    // Validate dates format
    const fechaInicioValida = /^\d{4}-\d{2}-\d{2}$/.test(nuevoCurso.fecha_Inicio);
    const fechaFinValida = /^\d{4}-\d{2}-\d{2}$/.test(nuevoCurso.fecha_Fin);

    if (!fechaInicioValida || !fechaFinValida) {
      alert('Por favor, ingrese las fechas en el formato YYYY-MM-DD.');
      return;
    }

    const cursoData = {
      fecha_inicio: formatDate(nuevoCurso.fecha_Inicio),
      fecha_fin: formatDate(nuevoCurso.fecha_Fin),
      requisitos: nuevoCurso.requisitos,
      instructor: nuevoCurso.instructor,
      materia_id: Number(nuevoCurso.materia_id),
    };

    try {
      const cursoResponse = await fetch("http://host.docker.internal:8090/crear_curso", {
        method: "POST",
        headers: {
          Authorization: `${token}`,
          "Content-Type": "application/json",
        },
        body: JSON.stringify(cursoData),
      });

      const cursoResult = await cursoResponse.json();

      if (cursoResponse.ok) {
        setCursos((prevCursos) => [...prevCursos, cursoResult]);
        alert("Se ha creado un nuevo curso con éxito");
        setNuevoCurso({
          fecha_Inicio: '',
          fecha_Fin: '',
          materia_id: null,
          requisitos: '',
          instructor: '',
        });
      } else {
        console.error(cursoResult.error);
      }
    } catch (error) {
      console.error("Error al crear el curso:", error);
    }
  };

  const handleInputChange = (event) => {
    const { name, value } = event.target;
    setNuevoCurso((prevNuevoCurso) => ({
      ...prevNuevoCurso,
      [name]: value,
    }));
  };


    const formatDate = (dateString) => {
      // Create a new Date object from the string
      const isoDate = new Date(dateString);
    
      // Extract day, month, and year
      const day = isoDate.getDate().toString().padStart(2, '0'); // Ensure two digits for day
      const month = (isoDate.getMonth() + 1).toString().padStart(2, '0'); // Ensure two digits for month
      const year = isoDate.getFullYear();
    
      // Return in DD/MM/YYYY format
      return `${day}/${month}/${year}`;
    };
    

  return (
    <div>
      <Header />
      <h2>Agregar nuevo curso</h2>
      <div className="contenedor-crear-curso">
        <form onSubmit={crearNuevoCurso}>
          <input
            type="text"
            name="fecha_Inicio"
            value={nuevoCurso.fecha_Inicio}
            onChange={handleInputChange}
            placeholder="Fecha de Inicio (YYYY-MM-DD)"
            required
          />
          <input
            type="text"
            name="fecha_Fin"
            value={nuevoCurso.fecha_Fin}
            onChange={handleInputChange}
            placeholder="Fecha de Fin (YYYY-MM-DD)"
            required
          />
          <textarea
            name="requisitos"
            value={nuevoCurso.requisitos}
            onChange={handleInputChange}
            placeholder="Requisitos"
            required
          />
          <input
            type="text"
            name="instructor"
            value={nuevoCurso.instructor}
            onChange={handleInputChange}
            placeholder="Instructor"
            required
          />
          <select
            name="materia_id"
            value={nuevoCurso.materia_id}
            onChange={handleInputChange}
            required
          >
            <option value="">Selecciona una materia</option>
            {materias.map((materia) => (
              <option key={materia.id} value={materia.id}>
                {materia.nombre} 
              </option>
            ))}
          </select>
          <button type="submit">Crear Curso</button>
        </form>
      </div>
      <div className="contenedor-cursos">
        <h2>Listado de Cursos</h2>
        <ul>
          {cursos.map((curso) => (
            <li key={curso.id}>
              <h2>Curso ID: {curso.id}</h2>
              <p>Fecha de Inicio: {formatDate(curso.fecha_inicio)}</p>
              <p>Fecha de Fin: {formatDate(curso.fecha_fin)}</p>
              <p>Requisitos: {curso.requisitos}</p>
              <p>Instructor: {curso.instructor}</p>
              <p>ID de Materia: {curso.materia_id}</p>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}

export default Admin;
