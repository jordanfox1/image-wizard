"use client";
import Image from "next/image";
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


  const convertToNewFormat = async (imageData, fileName: string, desiredFormat: string, addUpdateIndex) => {
    try {
      const apiEndpoint = process.env.NEXT_PUBLIC_API_ENDPOINT || 'http://image-wizard.local/api';
      console.log(apiEndpoint)

      const formData = new FormData();
      formData.append('image', imageData);
      formData.append('fileName', fileName);


      const response = await fetch(`${apiEndpoint}/convert?format=${desiredFormat}`, {
        method: 'POST',
        body: formData,
      });

      if (!response.ok) {
        throw new Error('Failed to convert image');
      }

      const responseData = await response.json();

      // Access the data URL directly from the response
      const dataURL = responseData.dataURL;
      const newFileName = responseData.fileName


      setImages(prevImages => {
        const updatedImages = [...prevImages];
        const newFile = new File([updatedImages[addUpdateIndex].file], newFileName);
        updatedImages[addUpdateIndex].dataURL = dataURL;
        updatedImages[addUpdateIndex].file = newFile;
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
                <Image src={image.dataURL} alt="" width={100} height={100} />
                <span>{image.file?.name}</span>
                <div className="image-item__btn-wrapper">
                  <button onClick={() => convertToNewFormat(image.dataURL, image.file.name, 'png', 0)}>
                    Convert to new format
                  </button>
                  <button onClick={() => onImageRemove(index)}>Remove</button>
                  <a href={image.dataURL} download={image.file.name}>
                    Download
                  </a>
                </div>
              </div>
            ))}
          </div>
        )}
      </ReactImageUploading>
    </div>
  );
}
