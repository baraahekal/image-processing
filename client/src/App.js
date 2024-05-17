import React from 'react';
import './App.css';
import Toggles from './Toggles';

function App() {
  const fetchData = () => {
    fetch('http://localhost:8080/api/hello')
        .then(response => response.text())
        .then(data => console.log(data));
  };

  return (
      <div className="App">
         <Toggles />
      </div>
  );
}

export default App;