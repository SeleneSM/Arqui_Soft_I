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

  const formatDateForAPI = (dateString) => {
    // Convierte YYYY-MM-DD a YYYY-MM-DDT00:00:00Z para que sea aceptado por el API (basicamente para que respete el formato que Go maneja las fechas por defecto)
    return `${dateString}T00:00:00Z`;
  };

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
      fecha_inicio: formatDateForAPI(nuevoCurso.fecha_Inicio),
      fecha_fin: formatDateForAPI(nuevoCurso.fecha_Fin),
      requisitos: nuevoCurso.requisitos,
      instructor: nuevoCurso.instructor,
      materia_id: parseInt(nuevoCurso.materia_id),
    };

    console.log('Datos a enviar:', cursoData); // Para debug

    try {
      const cursoResponse = await fetch("http://host.docker.internal:8090/crear_curso", {
        method: "POST",
        headers: {
          "Authorization": `Bearer ${token}`, // Agregamos 'Bearer' para indicar que es un token JWT
          "Content-Type": "application/json",
        },
        body: JSON.stringify(cursoData),
      });

      // Imprimir la respuesta completa para debug
      const responseText = await cursoResponse.text();
      console.log('Respuesta del servidor:', responseText);

      if (!cursoResponse.ok) {
        throw new Error(`Error ${cursoResponse.status}: ${responseText}`);
      }

      const cursoResult = JSON.parse(responseText);
      setCursos((prevCursos) => [...prevCursos, cursoResult]);
      alert("Se ha creado un nuevo curso con éxito");
      setNuevoCurso({
        fecha_Inicio: '',
        fecha_Fin: '',
        materia_id: null,
        requisitos: '',
        instructor: '',
      });
    } catch (error) {
      alert(`Error al crear el curso: ${error.message}`);
      console.error("Error completo:", error);
    }
  };

  const handleInputChange = (event) => {
    const { name, value } = event.target;
    setNuevoCurso((prevNuevoCurso) => ({
      ...prevNuevoCurso,
      [name]: value,
    }));
  };

  // La función formatDate solo la usamos para mostrar las fechas, no para enviarlas
  const formatDate = (dateString) => {
    try {
      const [year, month, day] = dateString.split('-');
      return `${day}/${month}/${year}`;
    } catch {
      return dateString; // Si el formato es diferente, retornamos el original
    }
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
