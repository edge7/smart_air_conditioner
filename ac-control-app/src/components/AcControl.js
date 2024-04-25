import React, { useEffect, useState } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPowerOff, faSnowflake } from '@fortawesome/free-solid-svg-icons';
import './ACControl.css';

function ACControl({ onToggle }) {
  const [isOn, setIsOn] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [feedback, setFeedback] = useState('');

  useEffect(() => {
    fetch('/status')
      .then(response => response.json())
      .then(data => setIsOn(data.isOn))
      .catch(error => {
        console.error("Failed to fetch status:", error);
        setFeedback("Failed to load status.");
      });
  }, []);

  const toggleAC = () => {
    setIsLoading(true); // Show loading indicator
    setFeedback(''); // Reset feedback message
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
        setIsLoading(false);
        setFeedback('Air conditioner status updated.');
        onToggle(); // Notify the parent
      })
      .catch(error => {
        console.error("Error toggling AC:", error);
        setIsLoading(false);
        setFeedback("Failed to update AC status.");
      });
  };

  return (
    <div className="ac-status">
      {isLoading && <p>Loading...</p>}
      <p>The air conditioner is currently {isOn ? 'ON' : 'OFF'}
        <FontAwesomeIcon icon={isOn ? faSnowflake : faPowerOff} className="mx-2" />
      </p>
      <button className={`ac-button ${isOn ? 'ac-button-on' : 'ac-button-off'}`} onClick={toggleAC} disabled={isLoading}>
        {isOn ? 'Turn Off' : 'Turn On'}
      </button>
      {feedback && <p>{feedback}</p>}
    </div>
  );
}

export default ACControl;
