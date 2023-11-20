"use client";
import React from "react";
import ReactImageUploading, { ImageListType } from "react-images-uploading";

export function ImageUpload() {
  const [images, setImages] = React.useState([]);
  const maxNumber = 69;

  const onChange = (
    imageList: ImageListType,
    addUpdateIndex: number[] | undefined
  ) => {
    // data for submit
    console.log(imageList, addUpdateIndex);
    setImages(imageList as never[]);
  };

  const convertToNewFormat = async (imageData, desiredFormat, addUpdateIndex) => {
    try {
      const apiEndpoint = process.env.NEXT_PUBLIC_API_ENDPOINT || 'http://image-wizard.local/api';
      console.log(apiEndpoint)

      const formData = new FormData();
      formData.append('image', imageData);


      const response = await fetch(`${apiEndpoint}/convert?format=${desiredFormat}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'multipart/form-data',
        },
        body: formData,
      });

      if (!response.ok) {
        throw new Error('Failed to convert image');
      }

      // Assuming the API returns the image bytes
      const newImageBytes = await response.arrayBuffer();

      // Convert the image bytes to a Data URL
      const newDataURL = `data:image/${desiredFormat};base64,${btoa(
        String.fromCharCode.apply(null, new Uint8Array(newImageBytes))
      )}`;

      // Update the state with the new image data
      setImages(prevImages => {
        const updatedImages = [...prevImages];
        updatedImages[addUpdateIndex].dataURL = newDataURL;
        return updatedImages;
      });
    } catch (error) {
      console.error('Error converting image:', error.message);
    }
  };


  return (
    <div className="image-upload">
      <ReactImageUploading
        multiple
        value={images}
        onChange={onChange}
        maxNumber={maxNumber}
      >
        {({
          imageList,
          onImageUpload,
          onImageRemoveAll,
          onImageUpdate,
          onImageRemove,
          isDragging,
          dragProps
        }) => (
          <div className="upload__image-wrapper">
            <button
              style={isDragging ? { color: "red" } : undefined}
              onClick={onImageUpload}
              {...dragProps}
            >
              Click or Drop here
            </button>
            &nbsp;
            <button onClick={onImageRemoveAll}>Remove all images</button>
            {imageList.map((image, index) => (
              <div key={index} className="image-item">
                <img src={image.dataURL} alt="" width={100} />
                <span>{image.file?.name}</span>
                <div className="image-item__btn-wrapper">
                  <button onClick={() => convertToNewFormat(image.dataURL, 'png', 0)}>
                    Convert to new format
                  </button>
                  <button onClick={() => onImageRemove(index)}>Remove</button>
                  <button onClick={() => onImageRemove(index)}>Download</button>
                </div>
              </div>
            ))}
          </div>
        )}
      </ReactImageUploading>
    </div>
  );
}
