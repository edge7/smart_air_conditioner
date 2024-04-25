// src/components/ImageDisplay.js

import React, { useEffect, useState } from 'react';

function ImageDisplay({ updateKey}) {
  const imageUrl = `/image?time=${updateKey}`;

  return (
    <div style={{
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        height: '100vh',
        backgroundColor: '#f0f0f0', // Soft background color
        padding: '20px' // Adds padding around the image
    }}>
      <img src={imageUrl} alt="Air Conditioner" style={{
          maxWidth: '100%',
          maxHeight: '90vh',
          border: '1px solid #ddd', // Adds a light border
          boxShadow: '0 5px 6px rgba(0, 0, 0, 0.1)', // Adds a shadow
          transition: 'transform 2.3s ease-in-out' // Smooth transition for transform
        }}
        onError={(e) => { e.target.onerror = null; e.target.src="https://via.placeholder.com/500"; }}
        onMouseOver={(e) => e.target.style.transform = 'scale(1.25)'} // Scales up image on hover
        onMouseOut={(e) => e.target.style.transform = 'scale(1)'} // Returns to original scale on mouse out
      />
    </div>
  );

}

export default ImageDisplay;