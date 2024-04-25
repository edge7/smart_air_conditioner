// src/components/ACControl.js

import React, { useEffect, useState } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPowerOff, faSnowflake } from '@fortawesome/free-solid-svg-icons';

function ACControl() {
  const [isOn, setIsOn] = useState(false);

  useEffect(() => {
    fetch('http://localhost:3030/status')
      .then(response => response.json())
      .then(data => setIsOn(data.isOn)); // Adjust according to your API response
  }, []);

  const toggleAC = () => {
    const newStatus = !isOn;
    fetch('http://localhost:3030/toggle', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ isOn: newStatus }),
    })
      .then(response => response.json())
      .then(() => setIsOn(newStatus));
  };

  return (
    <div className="ac-status">
      <p>The air conditioner is currently {isOn ? 'ON' : 'OFF'}
        <FontAwesomeIcon icon={isOn ? faSnowflake : faPowerOff} className="mx-2" />
      </p>
      <button className={`ac-button ${isOn ? 'ac-button-on' : 'ac-button-off'}`} onClick={toggleAC}>
        {isOn ? 'Turn Off' : 'Turn On'}
      </button>
    </div>
  );
}

export default ACControl;
