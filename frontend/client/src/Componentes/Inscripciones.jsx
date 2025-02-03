import React, { useState, useEffect, useCallback } from "react";
import '../Stylesheet/theme.css';
import '../Stylesheet/Inscripciones.css';
import Cursos from '../Componentes/Cursos';  
import BuscadorMaterias from '../Componentes/BuscadorMaterias'; 

function Inscripcion({ userId, token }) {
    const [inscripciones, setInscripciones] = useState([]);
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState(null);
    const [cursosDisponibles, setCursosDisponibles] = useState([]);
    const [materias, setMaterias] = useState([]);
    const [userData, setUserData] = useState(null);
    const [activeSection, setActiveSection] = useState('buscar');

    // Agregar función formatDate
    const formatDate = (dateString) => {
        if (!dateString) return 'Fecha no disponible';
        try {
            const date = new Date(dateString);
            const day = date.getDate().toString().padStart(2, '0');
            const month = (date.getMonth() + 1).toString().padStart(2, '0');
            const year = date.getFullYear();
            return `${day}/${month}/${year}`;
        } catch (error) {
            console.error('Error al formatear fecha:', error);
            return dateString;
        }
    };

    // Definir todas las funciones antes del useEffect
    const obtenerDetalles = async (inscripcionesData) => {
        try {
            // Primero obtenemos todas las materias
            const materiasResponse = await fetch('http://host.docker.internal:8090/materias');
            if (!materiasResponse.ok) {
                throw new Error('Error al obtener las materias');
            }
            const materiasData = await materiasResponse.json();

            // Crear un mapa de materias para búsqueda rápida
            const materiasMap = materiasData.reduce((acc, materia) => {
                acc[materia.id] = materia;
                return acc;
            }, {});

            const inscripcionesConDetalles = await Promise.all(
                inscripcionesData.map(async (inscripcion) => {
                    if (!inscripcion.id_curso) {
                        console.error('Inscripción sin id_curso:', inscripcion);
                        return null;
                    }

                    // Obtener detalles del curso
                    const cursoResponse = await fetch(
                        `http://host.docker.internal:8090/cursos/${inscripcion.id_curso}`,
                        {
                            headers: { "Content-Type": "application/json" }
                        }
                    );
                    
                    if (!cursoResponse.ok) {
                        throw new Error(`Error al obtener detalles del curso ${inscripcion.id_curso}`);
                    }
                    
                    const cursoData = await cursoResponse.json();
                    
                    // Obtener la materia del mapa local
                    const materiaData = materiasMap[cursoData.materia_id] || {
                        nombre: 'Materia no encontrada',
                        duracion: 0,
                        descripcion: 'Información no disponible',
                        palabras_clave: ''
                    };
                    
                    return {
                        id: inscripcion.id,
                        curso_id: inscripcion.id_curso,
                        fecha_inscripcion: inscripcion.fecha_inscripcion,
                        curso: {
                            instructor: cursoData.instructor,
                            fecha_inicio: cursoData.fecha_inicio,
                            fecha_fin: cursoData.fecha_fin,
                            requisitos: cursoData.requisitos
                        },
                        materia: materiaData
                    };
                })
            );

            const resultadosFiltrados = inscripcionesConDetalles.filter(item => item !== null);
            console.log('Inscripciones procesadas:', resultadosFiltrados);
            return resultadosFiltrados;

        } catch (error) {
            console.error('Error al obtener detalles:', error);
            return [];
        }
    };

    const fetchTodasLasInscripciones = useCallback(async () => {
        setIsLoading(true);
        setError(null);
        try {
            console.log('Fetching inscripciones para usuario:', userId);
            const response = await fetch(
                `http://host.docker.internal:8090/inscripciones_por_usuario/${userId}`,
                {
                    method: "GET",
                    headers: {
                        "Content-Type": "application/json"
                    }
                }
            );

            if (!response.ok) {
                throw new Error(`Error HTTP: ${response.status}`);
            }

            const inscripcionesData = await response.json();
            console.log('Datos de inscripciones recibidos:', inscripcionesData);

            if (inscripcionesData && inscripcionesData.length > 0) {
                const inscripcionesConDetalles = await obtenerDetalles(inscripcionesData);
                setInscripciones(inscripcionesConDetalles);
            } else {
                setInscripciones([]);
            }
        } catch (error) {
            console.error('Error en fetchTodasLasInscripciones:', error);
            setError(error.message);
        } finally {
            setIsLoading(false);
        }
    }, [userId]);

    const buscarCursosDisponibles = useCallback(async () => {
        try {
            // Primero obtenemos todos los cursos
            const cursosResponse = await fetch("http://host.docker.internal:8090/cursos");
            const cursosData = await cursosResponse.json();

            // Obtenemos todas las materias
            const materiasResponse = await fetch("http://host.docker.internal:8090/materias");
            const materiasData = await materiasResponse.json();

            // Crear un mapa de materias para búsqueda rápida
            const materiasMap = materiasData.reduce((acc, materia) => {
                acc[materia.id] = materia;
                return acc;
            }, {});

            // Combinar cursos con sus materias correspondientes
            const cursosConMaterias = cursosData.map(curso => ({
                ...curso,
                materia: materiasMap[curso.materia_id] || { 
                    nombre: 'Materia no encontrada',
                    duracion: 0,
                    descripcion: 'Información no disponible',
                    palabras_clave: ''
                }
            }));

            console.log('Cursos con materias:', cursosConMaterias);
            setCursosDisponibles(cursosConMaterias);
        } catch (error) {
            console.error('Error al cargar cursos:', error);
            setError('Error al cargar los cursos disponibles');
        }
    }, []);

    const fetchUserData = useCallback(async () => {
        try {
            const response = await fetch(`http://host.docker.internal:8090/users/${userId}`, {
                headers: {
                    "Content-Type": "application/json"
                }
            });

            if (!response.ok) {
                throw new Error('Error al obtener datos del usuario');
            }

            const data = await response.json();
            setUserData(data);
        } catch (error) {
            console.error('Error:', error);
        }
    }, [userId]);

    useEffect(() => {
        console.log('UseEffect ejecutándose con userId:', userId);
        if (userId) {
            const cargarDatos = async () => {
                setIsLoading(true);
                try {
                    await Promise.all([
                        fetchTodasLasInscripciones(),
                        buscarCursosDisponibles(),
                        fetchUserData()
                    ]);
                } catch (error) {
                    console.error('Error al cargar datos:', error);
                    setError('Error al cargar los datos');
                } finally {
                    setIsLoading(false);
                }
            };
            cargarDatos();
        }
    }, [userId, fetchTodasLasInscripciones, buscarCursosDisponibles, fetchUserData]); // Agregamos las dependencias correctas

    const buscarMaterias = (palabrasClave) => {
        if (!palabrasClave) {
            // Si no hay palabras clave, limpiamos los resultados
            setMaterias([]);
            return;
        }

        console.log('Buscando materias con palabras clave:', palabrasClave);
        
        fetch(`http://host.docker.internal:8090/materia/search/${palabrasClave}`, {
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
            setMaterias([]); // En caso de error, también limpiamos los resultados
        });
    };

    // Modificar el renderizado de los cursos disponibles
    const renderCursosDisponibles = () => {
        return cursosDisponibles.map((curso) => {
            // Verificar que el curso tiene todos los datos necesarios
            if (!curso || !curso.id) {
                console.error('Curso inválido:', curso);
                return null;
            }

            return (
                <Cursos
                    key={curso.id}
                    curso={{
                        ...curso,
                        materia: { nombre: 'Curso', ...curso.materia } // Proporcionar un valor por defecto
                    }}
                    isLoggedIn={true}
                    userId={userId}
                />
            );
        }).filter(Boolean); // Filtrar elementos nulos
    };

    const handleNavClick = (sectionId) => {
        setActiveSection(sectionId);
    };

    return (
        <div className="learning-platform">
            <div className="side-panel">
                <div className="profile-section">
                    <div className="profile-avatar">👤</div>
                    {userData && (
                        <div className="user-info">
                            <p className="user-name">{userData.nombre} {userData.apellido}</p>
                            <div className="user-details">
                                <div className="detail-item">
                                    <span className="detail-label">ID: {userData.id}</span>
                                    <span className="detail-value">{userData.dni}</span>
                                </div>
                                <div className="detail-item">
                                    <span className="detail-label">Rol:</span>
                                    <span className="detail-value">{userData.rol}</span>
                                </div>
                            </div>
                        </div>
                    )}
                </div>
                
                <nav className="nav-menu">
                    <a 
                        href="#buscar" 
                        className={`nav-item ${activeSection === 'buscar' ? 'active' : ''}`}
                        onClick={() => handleNavClick('buscar')}
                    >
                        <span className="nav-icon">🔍</span>
                        Buscar Materias
                    </a>
                    <a 
                        href="#inscripciones" 
                        className={`nav-item ${activeSection === 'inscripciones' ? 'active' : ''}`}
                        onClick={() => handleNavClick('inscripciones')}
                    >
                        <span className="nav-icon">📚</span>
                        Mis Inscripciones
                    </a>
                    <a 
                        href="#cursos" 
                        className={`nav-item ${activeSection === 'cursos' ? 'active' : ''}`}
                        onClick={() => handleNavClick('cursos')}
                    >
                        <span className="nav-icon">🎓</span>
                        Cursos Disponibles
                    </a>
                </nav>
            </div>

            <main className="main-content">
                <header className="main-header">
                    <div className="header-content">
                        <h1>Portal de Aprendizaje</h1>
                        <p>Bienvenido a tu espacio educativo personalizado</p>
                    </div>
                </header>

                {error && <div className="error-mensaje">{error}</div>}
                {isLoading ? (
                    <div className="loading">Cargando...</div>
                ) : (
                    <div className="content-grid">
                         <section id="buscar" className="content-section busqueda">
                            <div className="section-header">
                                <h2>Búsqueda de Materias</h2>
                            </div>
                            <div className="search-bar">
                                <BuscadorMaterias buscarMaterias={buscarMaterias} />
                            </div>
                            {materias.length > 0 && (
                                <div className="search-results-grid">
                                    {materias.map((materia) => (
                                        <div key={materia.id} className="search-result">
                                            <h3>{materia.nombre}</h3>
                                            <div className="info-group">
                                                <p><strong>Duración:</strong> {materia.duracion} meses</p>
                                                <p><strong>Descripción:</strong> {materia.descripcion}</p>
                                            </div>
                                            <div className="tags-container">
                                                {materia.palabras_clave.split(',').map((tag, i) => (
                                                    <span key={i} className="tag">{tag.trim()}</span>
                                                ))}
                                            </div>
                                        </div>
                                    ))}
                                </div>
                            )}
                        </section>
                        <section id="inscripciones" className="content-section inscripciones">
                            <div className="section-header">
                                <h2>Mis Inscripciones</h2>
                                <div className="stats">
                                    <span className="stat-item">
                                        Total: {inscripciones.length}
                                    </span>
                                    <span className="stat-item active">
                                        Activos: {inscripciones.length}
                                    </span>
                                </div>
                            </div>
                            <div className="cards-container">
                                {inscripciones.map((inscripcion) => (
                                    <div key={inscripcion.id} className="inscripcion-card">
                                        <div className="card-banner">
                                            <h3>{inscripcion.materia.nombre}</h3>
                                            <span className="status-badge">Activo</span>
                                        </div>
                                        <div className="card-content">
                                            <div className="info-group">
                                                <div className="info-item">
                                                    <span className="info-label">Instructor:</span>
                                                    <span>{inscripcion.curso.instructor}</span>
                                                </div>
                                                <div className="info-item">
                                                    <span className="info-label">Inicio:</span>
                                                    <span>{formatDate(inscripcion.curso.fecha_inicio)}</span>
                                                </div>
                                                <div className="info-item">
                                                    <span className="info-label">Fin:</span>
                                                    <span>{formatDate(inscripcion.curso.fecha_fin)}</span>
                                                </div>
                                                <div className="info-item">
                                                    <span className="info-label">Duración:</span>
                                                    <span>{inscripcion.materia.duracion} meses</span>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                ))}
                            </div>
                        </section>

                        <section id="cursos" className="content-section cursos">
                            <div className="section-header">
                                <h2>Cursos Disponibles</h2>
                            </div>
                            <div className="cards-container">
                                {renderCursosDisponibles()}
                            </div>
                        </section>
                    </div>
                )}
            </main>
        </div>
    );
}

export default Inscripcion;
