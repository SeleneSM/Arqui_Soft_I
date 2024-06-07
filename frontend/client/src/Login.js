import React, { useState, useEffect } from 'react';
import axios from "axios";
import Cookies from "js-cookie";

import '../Stylesheet/Login.css'; 
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import FloatingLabel from "react-bootstrap/FloatingLabel";
import { Alert } from "react-bootstrap";

import Header from '../Componentes/Header';
import { useNavigate } from 'react-router-dom';

function LoginRegister() {
/*  const [nombre, setNombre] = useState(''); //Almacenar nombre
  const [apellido, setApellido] = useState(''); //Almacenar apellido*/
  const [username, setUsername] = useState('');
  const [contraseña, setContraseña] = useState('');
  const [token, setToken] = useState('');
  const [error, setError] = useState('');
  const [isAdmin, setIsAdmin] = useState(false);
  const navigate = useNavigate();

  const handleLoginSubmit = async (event) => {
    event.preventDefault();
    if (username.trim() == '' || password.trim() == ''){
      alert("Uno de los campos se encuentra vacío, por favor completelos e intente nuevamente.");
      setUsername('');
      setPassword('');
    } else{
      const userData = {
        username: username,
        password: contraseña,
      };
  
      try {
        const authResponse = await fetch('http://localhost:8090/user/auth', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(userData),
        });
  
        if (authResponse.ok) {
          const responseJson = await authResponse.json();
          const tokenRecibido = responseJson.token; 
          setToken(tokenRecibido);
          const isAuthenticated = responseJson.autenticacion;
          console.log(isAuthenticated);
          const isAdmin = responseJson.tipo == 1; 
          setIsAdmin(isAdmin); 
  
          if (isAuthenticated == 'true') {
            if (isAdmin) {
              navigate(`/admin/${tokenRecibido}/${responseJson.user_id}`); 
            } else {
              navigate(`/reservacursos/${tokenRecibido}/${responseJson.user_id}`);
            }

          } else {
            alert("Credenciales inválidas.");
            window.location.reload();
          }
        } else {
          setError('Error en la autenticación');
        }
      } catch (error) {
        setError('Error en la solicitud');
        console.log(error);
      }

    }
  };

  const handleRegisterSubmit = async (event) => {
    event.preventDefault()
    if (username.trim() === '' || password.trim() === '' || nombre.trim() === '' || apellido.trim() === ''){
      alert("Uno de los campos se encuentra vacío, por favor completelos e intente nuevamente.");
      setNombre('');
      setApellido('');
      setUsername('');
      setPassword('');
    } else{
      const userData = {
        nombre: nombre,
        apellido: apellido,
        username:username,
        password: password,
      };
  
      try {
        const response = await fetch('http://localhost:8090/user', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(userData),
        });
  
        if (response.ok) {
          alert('El usuario ha sido registrado exitosamente');
          window.location.reload();
        } else {
          setError('Error en el registro');
        }
      } catch (error) {
        setError('Error en la solicitud');
        console.log(error);
      }
    }
  }

  return (
    <div className="main-loginregister-container">
      <Header />
      <div className='container-principal'>
        <div className="main">
          <input type="checkbox" id="chk" aria-hidden="true" />
          <div className="signup">
            <form onSubmit={handleRegisterSubmit}>
              <label for="chk" aria-hidden="true">Sign up</label>
              <input type="text" name="txt" placeholder="Nombre" required="" value={nombre} onChange={(e) => setNombre(e.target.value)} />
              <input type="text" name="txt" placeholder="Apellido" required="" value={apellido} onChange={(e) => setApellido(e.target.value)}/>
              <input type="username" name="username" placeholder="Username" required="" value={username}  onChange={(e) => setUsername(e.target.value)}/>
              <input type="password" name="pswd" placeholder="Contraseña" required="" value={password} onChange={(e) => setPassword(e.target.value)}/>
              <button className='sign' type="submit">Sign up</button>
            </form>
          </div>
          <div className="login">
            <form onSubmit={handleLoginSubmit}>
              <label for="chk" aria-hidden="true">Login</label>
              <input type="username" name="username" placeholder="Username" required="" value={username} onChange={(e) => setUsername(e.target.value)}/>
              <input type="password" name="pswd" placeholder="Password" required="" value={password} onChange={(e) => setPassword(e.target.value)}/>
              <button className='log' type="submit">Login</button>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
}

export default LoginRegister;
