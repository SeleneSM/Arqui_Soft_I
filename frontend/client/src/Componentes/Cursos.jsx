import React from 'react';
import '../Stylesheet/Cursos.css';
import pythonImg from '../imagenes/python.png'
import cImg from '../imagenes/c.png'
import rImg from '../imagenes/r.png'

function Cursos(props) {
  const { curso, isLoggedIn, userId } = props;
  const imageMap = {
    'Python': pythonImg,
    'C++': cImg,
    'R': rImg,
    // agrega más materias e imágenes según sea necesario
  };

  const formatDate = (dateString) => {
    // Parsea la cadena de fecha ISO 8601 a un objeto de fecha
    const isoDate = new Date(dateString);
    // Obtiene las partes de la fecha (año, mes, día) utilizando métodos de Date
    const year = isoDate.getFullYear();
    const month = isoDate.getMonth() + 1; // Los meses en JavaScript van de 0 a 11
    const day = isoDate.getDate();
    // Formatea la fecha según tus preferencias
    const formattedDate = `${day}/${month}/${year}`;
    return formattedDate;
  };

  const handleInscripcionClick = () => {
    if (!isLoggedIn) {
      alert("Debes iniciar sesión para realizar la reserva.");
    } else {
      realizarInscripcion();
    }
  };

  const realizarInscripcion = () => {
    console.log('Iniciando inscripción para el curso:', curso.id);
    console.log('User ID:', userId);
    fetch("http://host.docker.internal:8090/inscribir", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        id_curso: curso.id,
        id_usuario: userId,
        fecha_inscripcion: new Date().toISOString(),
      }),
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error('Error en la solicitud de inscripción');
        }
        return response.json();
      })
      .then((data) => {
        console.log('Respuesta de la inscripción:', data);
        if (data.id) {
          alert("Inscripción realizada con éxito!");
        } else if (data.error) {
          alert("No se puede realizar la inscripción: " + data.error);
        } else {
          alert("Hubo un problema al procesar la inscripción");
        }
      })
      .catch((error) => {
        console.error("Error en la inscripción:", error);
        alert("Hubo un error al realizar la inscripción");
      });
  };
  
  

  return (
    <div className="nombre-materia1">
      <h2>{curso.materia.nombre}</h2>
      <img src={imageMap[curso.materia.nombre]} alt={`Imagen de ${curso.materia.nombre}`} className="imagen-curso" />
      <p className="fecha-inicio">
      Fecha inicio: {formatDate(curso.fecha_inicio)}
      </p>
      <p className="fecha-fin">
      Fecha fin: {formatDate(curso.fecha_fin)}
      </p>
      <p className="informacion-instructor">
        Sobre el instructor: {curso.instructor}
      </p>
      <p className="requisitos">
        Requisitos: {curso.requisitos}
      </p>
      {curso.materia ? (
        <div className="materia">
          <p className="duracion">
            Duración en meses: {curso.materia.duracion}
          </p>
          <p className="descripcion">
            Descripción: {curso.materia.descripcion}
          </p>
          <p className="palabras-clave">
            Palabras clave: {curso.materia.palabras_clave}
          </p>
          <div className="boton-inscripcion">
            <button onClick={handleInscripcionClick}>Inscribirme</button>
          </div>
        </div>
      ) : (
        <p>No hay materia disponible</p>
      )}
    </div>
  );
}

export default Cursos;
