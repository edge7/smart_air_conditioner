// src/components/ImageDisplay.js

import React, { useEffect, useState } from 'react';

function ImageDisplay() {
  const imageUrl = '/image';

  return (
    <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
      <img src={imageUrl} alt="Air Conditioner" style={{ maxWidth: '100%', maxHeight: '90vh' }}
           onError={(e) => { e.target.onerror = null; e.target.src="https://via.placeholder.com/500"; }} />
    </div>
  );
}

export default ImageDisplay;