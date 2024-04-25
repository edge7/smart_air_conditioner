import React, { useState } from 'react';
import ACControl from './AcControl';
import ImageDisplay from './ImageDisplay';

function ParentComponent() {
  const [acToggle, setAcToggle] = useState(0); // Use a counter to track toggles

  const handleToggle = () => {
    setAcToggle(prev => prev + 1); // Increment to trigger updates
  };

  return (
    <div>
      <ACControl onToggle={handleToggle} />
      <ImageDisplay key={acToggle} />
    </div>
  );
}

export default ParentComponent;