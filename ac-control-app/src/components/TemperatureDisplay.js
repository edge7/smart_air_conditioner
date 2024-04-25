import React, { useEffect, useState } from 'react';
import './TemperatureDisplay.css'; // Ensure the path is correct based on your project structure

function TemperatureDisplay() {
  const [temperature, setTemperature] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch('/temperature') // Adjust this URL to where your temperature data is served
      .then(response => response.json())
      .then(data => {
        const roundedTemperature = parseFloat(data.temperature.toFixed(2));
        setTemperature(roundedTemperature);
        setLoading(false);
      })
      .catch(error => {
        console.error('Error fetching temperature:', error);
        setLoading(false);
      });
  }, []);

  if (loading) {
    return <p className="loading-text">Loading temperature...</p>;
  }

  return (
    <div className="temperature-display">
      Current Temperature: {temperature}°C
    </div>
  );
}

export default TemperatureDisplay;
