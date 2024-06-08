import React, { useEffect, useRef, useState } from 'react';
import '../Stylesheet/Cursos.css'; 

function Cursos(props) {
    const [cursos, setCursos] = useState([]);
    //const [materias, setMaterias] = useState([]);
    const cursoIdRef = useRef(props.cursoId);

    useEffect(() => {
        // Obtener los IDs de los hoteles mediante una solicitud fetch
        fetch('http://localhost:8090/cursos')
            .then(response => response.json())
            .then(data => {
                setCursos(data);
            })
            .catch(error => {
                console.error('Error:', error);
            });
    }, []);

    useEffect(() => {
        cursoIdRef.current = props.cursoId;
    }, [props.cursoId]);

    const handleInscripcionClick = () => {
        if (!props.isLoggedIn) {
          alert("Debes iniciar sesión para realizar la reserva.");
        } 
    };
      
    const realizarInscripcion = () => {
        fetch("http://localhost:8090/inscripciones", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            curso_id: props.cursoId, // Agrega el ID del curso correspondiente
            user_id: props.userId, // Agrega el ID del usuario correspondiente
          }),
        })
            .then((response) => response.json())
            .then((data) => {
              alert("Inscripcion realizada con éxito!");
            })
            .catch((error) => {
              console.error("Error:", error);
            });
      };


    return (
        <div className="contenedor-cursos">
            <p className="nombre-materia1">
                <strong>{props.nombreMateria}</strong>
            </p>
            <p className="fecha-inicio">
                Fecha inicio: {props.fechaInicio}
            </p>
            <p className="fecha-fin">
                Fecha fin: {props.fechaFin}
            </p>
            <p className="descripcion-materia">
                Descripción: {props.descripcion}
            </p>
            <p className="informacion-instructor">
                Sobre el instructor: {props.informacionInstructor}
            </p>
            <p className="requisitos">
                Requisitos: {props.requisitos}
            </p> 
            <div className="boton-inscripcion">
                <button onClick={handleInscripcionClick}>Inscribirme</button>
            </div>
        </div>
        
    );
}

export default Cursos;
