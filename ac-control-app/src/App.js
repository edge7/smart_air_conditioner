import React from 'react';
import ImageDisplay from './components/ImageDisplay';
import ACControl from './components/AcControl';
import TemperatureDisplay from './components/TemperatureDisplay';

function App() {
  return (
    <div>
      <ImageDisplay />
      <TemperatureDisplay />
      <ACControl />
    </div>
  );
}

export default App;
