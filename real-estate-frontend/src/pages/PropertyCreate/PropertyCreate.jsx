import React, { useState } from 'react';
import './PropertyCreate.css';
import { useNavigate } from "react-router-dom";
import useAuth from '../../hooks/useAuth';

const PropertyCreate = () => {
  const navigate = useNavigate()
  const { authAxios } = useAuth()
  const [formData, setFormData] = useState({
    title: '',
    description: '',
    maxPrice: '',
    minPrice: '',
    location: '',
    bedrooms: '',
    bathrooms: '',
    squareFeet: '',
    city:'',
    district:'',
    images: [],

  });

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const handleFileChange = (e) => {
    const files = Array.from(e.target.files);
    const existingFileNames = formData.images.map((file) => file.name);

    const newFiles = files.filter((file) => !existingFileNames.includes(file.name));

    if (newFiles.length < files.length) {
      alert('Some files were skipped because they have duplicate names.');
    }

    setFormData({
      ...formData,
      images: [...formData.images, ...newFiles],
    });

    e.target.value = '';
  };

  const handleRemoveImage = (index) => {
    const updatedImages = formData.images.filter((_, i) => i !== index);
    setFormData({
      ...formData,
      images: updatedImages,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const formDataToSubmit = new FormData();

    // Append files
    formData.images.forEach((file) => {
      formDataToSubmit.append('images', file);
    });
    if(formData.images.length === 0){
      alert("You need at least 1 image")
      return
    }

    const dataObject = {
      title: formData.title,
      description: formData.description,
      maxPrice: formData.maxPrice,
      minPrice: formData.minPrice,
      location: formData.location,
      bedrooms: formData.bedrooms,
      bathrooms: formData.bathrooms,
      squareFeet: formData.squareFeet,
    };

    formDataToSubmit.append('data', JSON.stringify(dataObject));

    try {
      const response = await authAxios.post('/properties/create', formDataToSubmit, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });

      if (response.status === 200) {
        alert('Form submitted successfully!');
      }

    } catch (error) {
      console.error('Error:', error);
    }
  };
  
  const handleClick = (e, path) => {
    e.preventDefault();
    navigate(path);
};

  return (
    <div className="property-create-page">
      <h1>Create Property</h1>
      <form onSubmit={handleSubmit} className="property-form" encType="multipart/form-data">
        <div className="form-row">
          <label>
            Title:
            <input
              type="text"
              name="title"
              value={formData.title}
              onChange={handleInputChange}
              required
            />
          </label>
        </div>
        <div className="form-row">
          <label>
            Description:
            <textarea
              name="description"
              value={formData.description}
              onChange={handleInputChange}
              required
            />
          </label>
        </div>
        <div className="form-row horizontal">
          <label>
            Max Price:
            <input
              type="number"
              name="maxPrice"
              value={formData.maxPrice}
              onChange={handleInputChange}
              required
            />
          </label>
          <label>
            Min Price:
            <input
              type="number"
              name="minPrice"
              value={formData.minPrice}
              onChange={handleInputChange}
              required
            />
          </label>
        </div>
        <div className="form-row horizontal">
          <label>
            City:
            <input
              type="text"
              name="city"
              value={formData.city}
              onChange={handleInputChange}
              required
            />
          </label>
          <label>
            District:
            <input
              type="text"
              name="district"
              value={formData.district}
              onChange={handleInputChange}
              required
            />
          </label>
        </div>
        <div className="form-row">
          <label>
            Location Details:
            <input
              type="text"
              name="location"
              value={formData.location}
              onChange={handleInputChange}
              required
            />
          </label>
        </div>
        <div className="form-row horizontal">
          <label>
            Bedrooms:
            <input
              type="number"
              name="bedrooms"
              value={formData.bedrooms}
              onChange={handleInputChange}
              required
            />
          </label>
          <label>
            Bathrooms:
            <input
              type="number"
              name="bathrooms"
              value={formData.bathrooms}
              onChange={handleInputChange}
              required
            />
          </label>
          <label>
            Square Feet:
            <input
              type="number"
              name="squareFeet"
              value={formData.squareFeet}
              onChange={handleInputChange}
              required
            />
          </label>
        </div>
        <div className="form-row">
          <label>Images:</label>
          <input
            type="file"
            multiple
            accept="image/*"
            onChange={handleFileChange}
            className="file-input"
          />
          <div className="uploaded-images">
            {formData.images.map((file, index) => (
              <div key={index} className="image-preview">
                <img
                  src={URL.createObjectURL(file)}
                  alt={`Uploaded ${index}`}
                  className="preview-img"
                />
                <button
                  type="button"
                  onClick={() => handleRemoveImage(index)}
                  className="remove-image-button"
                >
                  Remove
                </button>
              </div>
            ))}
          </div>
        </div>
        <div className="form-row-action">
          <button className="back-button" type="button" onClick={(e) => handleClick(e, "/dashboard/propertiesManager")}>
            Back
          </button>
          <button type="submit" className="submit-button">
            Submit
          </button>
        </div>
      </form>
    </div>
  );
};

export default PropertyCreate;
