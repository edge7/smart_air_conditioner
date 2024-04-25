// src/components/ACControl.js

import React, { useEffect, useState } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPowerOff, faSnowflake } from '@fortawesome/free-solid-svg-icons';

function ACControl({ onToggle }) {
  const [isOn, setIsOn] = useState(false);

  useEffect(() => {
    fetch('/status')
      .then(response => response.json())
      .then(data => setIsOn(data.isOn));
  }, []);

  const toggleAC = () => {
    const newStatus = !isOn;
    fetch('/toggle', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ isOn: newStatus }),
    })
      .then(response => response.json())
      .then(() => {
        setIsOn(newStatus);
        onToggle(); // Call the prop function to notify the parent
      });
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
