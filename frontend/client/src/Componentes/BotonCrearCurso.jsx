import React, { useState } from 'react';
import Admin from './Admin';

function BotonCrearCurso() {
  const [isOpen, setIsOpen] = useState(false);

  const openPopup = () => {
    setIsOpen(true);
  };

  const closePopup = () => {
    setIsOpen(false);
  };

  return (
    <>
      <button onClick={openPopup}>Crear Curso</button>
      {isOpen && (
        <div className="popup">
          <Admin />
          <button onClick={closePopup}>Cerrar</button>
        </div>
      )}
    </>
  );
}

export default BotonCrearCurso;