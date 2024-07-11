import React, { useState } from 'react';

function FormularioRegister() {
  const [nombre, setNombre] = useState('');
  const [apellido, setApellido] = useState('');
  const [username, setUsername] = useState('');
  const [rol, setRol] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleRegister = async (e) => {
    e.preventDefault();

    // Crear el objeto de datos a enviar en la solicitud POST
    const userData = {
      nombre: nombre,
      apellido: apellido,
      username: username,
      rol: rol,
      password: password,
    };

    try {
      const response = await fetch('http://localhost:8090/registrar_usuario', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(userData),
      });

      if (response.ok) {
        alert('El usuario ha sido registrado exitosamente');
        setNombre('');
        setApellido('');
        setUsername('');
        setRol('');
        setPassword('');
      } else {
        setError('Error en el registro');
      }
    } catch (error) {
      setError('Error en la solicitud');
      console.log(error);
    }
  };

  return (
    <form onSubmit={handleRegister}>
      <h3>Registrarse</h3>
      <p>Nombre:</p>
      <input type="text" className='campoNombre' value={nombre} onChange={(e) => setNombre(e.target.value)}/>
      <p>Apellido:</p>
      <input type="text" className='campoApellido' value={apellido} onChange={(e) => setApellido(e.target.value)}/>
      <p>Username:</p>
      <input type="text" className='campoUsername' value={username} onChange={(e) => setUsername(e.target.value)}/>
      <p>Rol Estudiante o Administrador?:</p>
      <input type="text" className='campoRol' value={rol} onChange={(e) => setRol(e.target.value)}/>
      <p>Contraseña:</p>
      <input type="password" className='campoContraseña' value={password} onChange={(e) => setPassword(e.target.value)}/>
      <button type="submit">Registrarse</button>
      {error && <p>{error}</p>}
    </form>
  );
}

export default FormularioRegister;
