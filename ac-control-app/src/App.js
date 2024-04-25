import React from 'react';
import TemperatureDisplay from './components/TemperatureDisplay';
import ParentComponent from './components/ParentComponent'; // Correct path to your ParentComponent

function App() {
  return (
    <div>
      <TemperatureDisplay />
      <ParentComponent />
    </div>
  );
}

export default App;
