import React from 'react';
import './ImageDisplay.css'; // Ensure the path is correct based on your project structure

function ImageDisplay({ updateKey }) {
  const imageUrl = `/image?time=${updateKey}`;

  return (
    <div className="image-container">
      <img src={imageUrl} alt="Air Conditioner" className="ac-image"
           onError={(e) => { e.target.onerror = null; e.target.src="https://via.placeholder.com/500"; }}
      />
    </div>
  );
}
export default ImageDisplay;
