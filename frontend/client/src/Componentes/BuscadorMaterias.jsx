import React, { useState } from 'react';

function BuscadorMaterias({ buscarMaterias }) {
  const [palabrasClave, setPalabrasClave] = useState('');

  const handleBuscarClick = () => {
    buscarMaterias(palabrasClave);
  };

  return (
    <div>
      <input
        type="text"
        value={palabrasClave}
        onChange={(e) => setPalabrasClave(e.target.value)}
        placeholder="Buscar materias por palabras clave"
      />
      <button onClick={handleBuscarClick}>Buscar</button>
    </div>
  );
}

export default BuscadorMaterias;
