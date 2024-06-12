import React, { useState, useEffect } from "react";
//import '../Stylesheet/Inscripciones.css';
//import Header from '../Componentes/Header';
import { useParams } from 'react-router-dom';
import Cursos from '../Componentes/Cursos';  

function inscripcion() {
    const [cursosDisponibles, setCursosDisponibles] = useState([]);
    const [inscripciones, setInscripciones] = useState([]);
    //const [contadorInscripcion, setContadorInscripcion] = useState(1);
    const { token, user_id } = useParams();

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

    const fetchTodasLasInscripciones = () => {
        fetch(`http://localhost:8090/inscripciones_por_usuario/${user_id}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `${token}`,
            },
        })
        .then((response) => response.json())
        .then((data) => {
            if (data && data.length > 0) {
                setInscripciones(data);
            } else {
                setInscripciones([]);
            }
        })
        .catch((error) => console.error(error));
    };

    useEffect(() => {
        fetchTodasLasInscripciones();
        buscarCursosDisponibles();
    }, []);

    return (
        <div className="contenedor-principal">
            <Header />
            <h4 className="titulo-inscripciones">Mis inscripciones:</h4>
            <div className="contenedor-inscripciones-usuario">
                {inscripciones.length > 0 ? (
                    inscripciones.map((inscripcion, index) => (
                        <div key={inscripcion.id} className="inscripcion-item">
                            <p className="subtitulo-inscripcion">Datos de la inscripción {contadorInscripcion + index}:</p>
                            <div className="detalle-inscripcion">
                                <p>Curso: {inscripcion.nombreCurso}</p>
                                <p>Fecha inicio: {formatDate(inscripcion.fechaInicio)}</p>
                                <p>Fecha fin: {formatDate(inscripcion.fechaFin)}</p>
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
                            cursoId={curso.id}
                            nombreMateria={curso.nombre}
                            fechaInicio={curso.fechaInicio}
                            fechaFin={curso.fechaFin}
                            descripcion={curso.descripcion}
                            informacionInstructor={curso.informacionInstructor}
                            requisitos={curso.requisitos}
                            userId={user_id}
                            isLoggedIn={true}
                        />
                    ))
                ) : (
                    <p>No hay cursos disponibles en este momento.</p>
                )}
            </div>
        </div>
    );
}

export default inscripcion;
