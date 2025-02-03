import React, { useState } from 'react';

function BuscadorMaterias({ buscarMaterias }) {
    const [palabrasClave, setPalabrasClave] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault();
        if (palabrasClave.trim()) {
            buscarMaterias(palabrasClave);
        }
    };

    const handleChange = (e) => {
        const valor = e.target.value;
        setPalabrasClave(valor);
        
        // Si el input está vacío, limpiamos los resultados
        if (!valor.trim()) {
            buscarMaterias(''); // Esto disparará la limpieza de resultados
        }
    };

    return (
        <form onSubmit={handleSubmit} className="buscador-form">
            <input
                type="text"
                value={palabrasClave}
                onChange={handleChange}
                placeholder="Buscar materias..."
                className="buscador-input"
            />
            <button type="submit" disabled={!palabrasClave.trim()}>
                Buscar
            </button>
        </form>
    );
}

export default BuscadorMaterias;
