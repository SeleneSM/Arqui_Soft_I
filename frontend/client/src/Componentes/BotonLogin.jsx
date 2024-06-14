import React, { useState } from 'react';
import FormularioLogin from './FormularioLogin';
import '../Stylesheet/BotonLogin.css';

function BotonLogin({ handleLogin }) {
  const [isOpen, setIsOpen] = useState(false);

  const openPopup = () => {
    setIsOpen(true);
  };

  const closePopup = () => {
    setIsOpen(false);
  };

  return (
    <>
    <div className="iniciarsesion">
      <button onClick={openPopup}>Iniciar sesión</button>
      </div>
      {isOpen && (
        <div className="popup">
          <FormularioLogin handleLogin={handleLogin}/>
          <button onClick={closePopup}>Cerrar</button>
        </div>
      )}
    </>
  );
}

export default BotonLogin;