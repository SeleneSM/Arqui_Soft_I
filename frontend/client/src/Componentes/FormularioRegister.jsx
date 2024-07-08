import React, { useState } from 'react';

function FormularioRegister() {
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleRegister = (e) => {
    e.preventDefault();

    // Crear el objeto de datos a enviar en la solicitud POST
    const userData = {
      name: firstName,
      last_name: lastName,
      username: username,
      password: password,
    };

    // Realizar la solicitud POST al backend
    fetch('http://localhost:8090/users', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(userData),
    })
      .then(response => response.json())
      .then(data => {
        // Aquí puedes manejar la respuesta del servidor, si es necesario
        console.log(data);
        alert('El usuario ha sido registrado exitosamente');
        setFirstName('');
        setLastName('');
        setUsername('');
        setPassword('');
      })
      .catch(error => {
        // Aquí puedes manejar el error, si ocurre alguno
        console.error('Error:', error);
      });
  };

  return (
    <form onSubmit={handleRegister}>
      <h3>Registrarse</h3>
      <p>Nombre:</p>
      <input type="text" className='campoNombre' value={firstName} onChange={(e) => setFirstName(e.target.value)}/>
      <p>Apellido:</p>
      <input type="text" className='campoApellido' value={lastName} onChange={(e) => setLastName(e.target.value)}/>
      <p>Username:</p>
      <input type="text" className='campoUsername' value={username} onChange={(e) => setUsername(e.target.value)}/>
      <p>Contraseña:</p>
      <input type="password" className='campoContraseña' value={password} onChange={(e) => setPassword(e.target.value)}/>
      <button type="submit">Registrarse</button>
    </form>
  );
}

export default FormularioRegister;