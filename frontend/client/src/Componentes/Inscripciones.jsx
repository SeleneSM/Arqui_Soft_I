import React, { useState, useEffect } from "react";
import '../Stylesheet/Inscripciones.css';
import { useParams } from 'react-router-dom';
import Cursos from '../Componentes/Cursos';  
import BuscadorMaterias from '../Componentes/BuscadorMaterias'; 

function Inscripcion() {
    const [cursosDisponibles, setCursosDisponibles] = useState([]);
    const [cursos, setCursos] = useState([]);
    const [usuarios, setUsuarios] = useState([]);
    const [inscripciones, setInscripciones] = useState([]);
    const [contadorInscripcion, setContadorInscripcion] = useState(1);
    const [materias, setMaterias] = useState([]);
    const { token, userId } = useParams();

    const buscarCursosDisponibles = () => {
        fetch("http://localhost:8090/cursos", {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `${token}`,
            }
        })
        .then((response) => response.json())
        .then((data) => {
            setCursosDisponibles(data);
        })
        .catch((error) => console.error(error));
    };

    const buscarMaterias = (palabrasClave) => {
        console.log('Buscando materias con palabras clave:', palabrasClave);
        
        fetch(`http://localhost:8090/materia/search/${palabrasClave}`, {
          method: "GET",
          headers: {
            "Content-Type": "application/json"
          },
        })
        .then((response) => {
          if (!response.ok) {
            throw new Error('Error al buscar materias');
          }
          return response.json();
        })
        .then((materias) => {
          setMaterias(materias);
        })
        .catch((error) => {
          console.error('Error al buscar materias:', error);
        });
      };
      
      

    const fetchTodasLasInscripciones = () => {
        fetch(`http://localhost:8090/inscripciones_por_usuario/${userId}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `${token}`,
            },
        })
        .then((response) => {
            console.log('User ID:', userId);

            console.log('Response status:', response.status);
            return response.json();
          })
        .then((data) => {
            console.log('Data recibida:', data);

            if (data && data.length > 0) {
              console.log('Data de inscripciones:', data);
              obtenerDetalles(data);
            } else {
              console.log('No hay inscripciones para el usuario');
              setInscripciones([]);
            }})
        .catch((error) => console.error(error));
    };
    

    const obtenerDetalles = (inscripcionesData) => {
        Promise.all([
          fetch("http://localhost:8090/cursos").then((response) => response.json()),
          fetch("http://localhost:8090/users").then((response) => response.json()),
        ])
          .then(([cursosData, usersData]) => {
            setCursos(cursosData);
            setUsuarios(usersData);
    
            console.log('infocursos:', cursosData);
            console.log('infouser:', usersData);

            const cursosMap = {};
            cursosData.forEach((curso) => {
              cursosMap[curso.id] = curso.materia.nombre;
              console.log('cursooo:', cursosMap[curso.id]);
            });
      
            const usuariosMap = {};
            usersData.forEach((user) => {
              usuariosMap[user.id] = user.username;
              console.log('useeers:', usuariosMap[user.id]);
            });
    
            const inscripcionesActualizadas = inscripcionesData.map((inscripcion) => {
                const cursoNombre = cursosMap[inscripcion.id_curso] || "";
                const usuarioNombre = usuariosMap[inscripcion.userId] || "";
        
                console.log(`Inscripción ID: ${inscripcion.id}`);
                console.log(`Curso Nombre: ${cursoNombre}`);
                console.log(`Usuario Nombre: ${usuarioNombre}`);
        
                return {
                  ...inscripcion,
                  curso_nombre: cursoNombre,
                  usuario_nombre: usuarioNombre,
                };
              });
    
            setInscripciones(inscripcionesActualizadas); // Actualiza el estado con las inscripciones detalladas
          })
          .catch((error) => console.error(error));
      };


    useEffect(() => {
        fetchTodasLasInscripciones();
        buscarCursosDisponibles();
    }, [userId]);

   return (
        <div className="contenedor-principal">
            <h4 className="titulo-inscripciones">Mis inscripciones:</h4>
            <div className="contenedor-inscripciones-usuario">
                {inscripciones.length > 0 ? (
                    inscripciones.map((inscripcion, index) => (
                        <div key={inscripcion.id} className="inscripcion-item">
                            <p className="subtitulo-inscripcion">Datos de la inscripción {contadorInscripcion + index}:</p>
                            <div className="detalle-inscripcion">
                                <p>Curso: {inscripcion.curso_nombre}</p>
                                <p>Fecha inicio: {inscripcion.fechaInicio}</p>
                                <p>Fecha fin: {inscripcion.fechaFin}</p>
                            </div>
                        </div>
                    ))
                ) : (
                    <p>No tienes inscripciones realizadas.</p>
                )}
            </div>
            <h2>Cursos disponibles:</h2>
            <div className="contenedor-cursos">
                {cursosDisponibles.length > 0 ? (
                    cursosDisponibles.map((curso) => (
                        <Cursos
                            key={curso.id}
                            curso={curso}
                            isLoggedIn={true}
                            id_usuario={userId}
                        />
                    ))
                ) : (
                    <p>No hay cursos disponibles en este momento.</p>
                )}
            </div>
            <h2>Buscar Materias:</h2>
            <BuscadorMaterias buscarMaterias={(palabrasClave) => buscarMaterias(palabrasClave, token)} />
            {materias.length > 0 && (
                <div className="contenedor-materias">
                    {materias.map((materia) => (
                        <div key={materia.id} className="materia-item">
                            <h3>{materia.nombre}</h3>
                            <p>Duración: {materia.duracion} meses</p>
                            <p>Descripción: {materia.descripcion}</p>
                            <p>Palabras clave: {materia.palabras_clave}</p>
                        </div>
                    ))}
                </div>
            )}
        </div>
    );
}

export default Inscripcion;
