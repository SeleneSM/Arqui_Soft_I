import React, { useState, useEffect } from 'react';
import '../Stylesheet/Login.css'; 

function FormularioLogin({handleLogin}) {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  /*const [token, setToken] = useState('');
  const [error, setError] = useState('');
  const [isAdmin, setIsAdmin] = useState(false);
  const [loggedIn, setLoggedIn] = useState(false);
  //const navigate = useNavigate();*/


  const handleSubmit = async (event) => {
    event.preventDefault();
    if (username.trim() == '' || password.trim() == ''){
      alert("Uno de los campos se encuentra vacío, por favor completelos e intente nuevamente.");
      setUsername('');
      setPassword('');
    } else{
      const userData = {
        username: username,
        password: password,
      };
  
      try {
        const response = await fetch('http://localhost:8090/users/auth', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(userData),
        });

        const data = await response.json();

        if (data.auth) {
            const tipoUsuario = data.rol;
            const userId = data.user_id;

            handleLogin(tipoUsuario, userId)
            alert('Autenticacion exitosa');
          }else{
            alert('La autenticación fue incorrecta. Ingrese sus datos nuevamente');
          }
        }catch(error) {
          console.error('Error:', error);
          alert('Ocurrió un error durante la autenticación. Por favor, inténtelo de nuevo más tarde.');
        }
      }
  };


  return (
    <form onSubmit={handleSubmit}>
      <h3>Iniciar sesión</h3>
      <p>Email:</p>
      <input type="username" className='campoUsername' value={username} onChange={(e) => setUsername(e.target.value)} />
      <p>Contraseña:</p>
      <input type="password" className='campoPassword' value={password}  onChange={(e) => setPassword(e.target.value)} />
      <button type="submit">Iniciar sesión</button>
    </form>
  );


}

export default FormularioLogin;
