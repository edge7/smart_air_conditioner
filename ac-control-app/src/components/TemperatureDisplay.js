import React, { useEffect, useState } from 'react';

function TemperatureDisplay() {
  const [temperature, setTemperature] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch('http://localhost:3030/temperature') // Adjust this URL to where your temperature data is served
      .then(response => response.json())
      .then(data => {
        setTemperature(data.temperature); // Adjust "temperature" based on how data is structured in your response
        setLoading(false);
      })
      .catch(error => {
        console.error('Error fetching temperature:', error);
        setLoading(false);
      });
  }, []);

  if (loading) {
    return <p>Loading temperature...</p>;
  }

  return (
    <div style={{ textAlign: 'center', padding: '20px', fontSize: '24px' }}>
      Current Temperature: {temperature}°C
    </div>
  );
}

export default TemperatureDisplay;
