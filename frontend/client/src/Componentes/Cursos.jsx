import React from 'react';
import '../Stylesheet/Cursos.css';
import pythonImg from '../imagenes/python.png'
import cImg from '../imagenes/c.png'
import rImg from '../imagenes/r.png'

function Cursos(props) {
  const { curso, isLoggedIn, userId } = props;

  // Validación de props
  if (!curso) {
    console.error('Curso props es undefined');
    return <div>Error: Datos del curso no disponibles</div>;
  }

  // Asegurar que materia existe
  const materia = curso.materia || { nombre: 'Curso sin nombre' };

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
      {imageMap[curso.materia.nombre] && (
        <img 
          src={imageMap[curso.materia.nombre]} 
          alt={`Imagen de ${curso.materia.nombre}`} 
          className="imagen-curso" 
        />
      )}
      <div className="curso-info">
        <div className="fechas-info">
          <p className="fecha-inicio">
            <strong>Inicio:</strong> {formatDate(curso.fecha_inicio)}
          </p>
          <p className="fecha-fin">
            <strong>Fin:</strong> {formatDate(curso.fecha_fin)}
          </p>
        </div>
        
        <div className="materia-info">
          <p className="duracion">
            <strong>Duración:</strong> {curso.materia.duracion} meses
          </p>
          <p className="descripcion">
            {curso.materia.descripcion}
          </p>
          <p className="palabras-clave">
            <strong>Tags:</strong> {curso.materia.palabras_clave}
          </p>
        </div>

        <div className="curso-detalles">
          <p className="informacion-instructor">
            <strong>Instructor:</strong> {curso.instructor}
          </p>
          <p className="requisitos">
            <strong>Requisitos:</strong> {curso.requisitos}
          </p>
        </div>

        {isLoggedIn && (
          <div className="boton-inscripcion">
            <button onClick={handleInscripcionClick}>Inscribirme</button>
          </div>
        )}
      </div>
    </div>
  );
}

export default Cursos;
