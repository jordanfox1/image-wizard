"use client";
import Image from "next/image";
import React, { useState } from "react";
import ReactImageUploading, { ImageListType } from "react-images-uploading";
import { AspectRatio, Button, Chip, ChipGroup } from '@mantine/core';
import './imageUpload.css'

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

  const [desiredFormat, setDesiredFormat] = useState('webp');

  return (
    <div className="image-upload-wrapper-container">
      <ReactImageUploading
        multiple
        value={images}
        onChange={onChange}
        maxNumber={maxNumber}
      >
        {({
          imageList,
          onImageUpload,
          onImageRemove,
          isDragging,
          dragProps
        }) => (
          <>
            <div className="drop-zone" style={isDragging ? { color: "red" } : undefined} {...dragProps}>
              <Button className="btn-large" onClick={onImageUpload} >
                Select Files
              </Button>
              Click or drop files here
            </div>

            {imageList.map((image, index) => (
              <>
                <div className="format-select-btn-container">
                  <Chip.Group multiple={false} value={desiredFormat} onChange={setDesiredFormat}>
                    <Chip value='webp'>WEBP</Chip>
                    <Chip value='png'>PNG</Chip>
                    <Chip value='jpeg'>JPEG</Chip>
                    <Chip value='gif'>GIF</Chip>
                    <Chip value='bmp'>BMP</Chip>
                  </Chip.Group>
                </div>

                <div key={index} className="image-item">
                  <figure className="image-figure">
                    <Image className="image" src={image.dataURL} alt="your uploaded image" width={280} height={160} />
                    <figcaption>{image.file?.name}</figcaption>
                  </figure>


                  <div className="image-item__btn-wrapper">
                    <Button className="btn-med" onClick={() => convertToNewFormat(image.dataURL, image.file.name, desiredFormat, 0)}>
                      Convert to {desiredFormat}
                    </Button>

                    <Button className="btn-med" onClick={() => onImageRemove(index)}>Remove</Button>

                    <a className="btn-med" href={image.dataURL} download={image.file.name}>
                      Download
                    </a>
                  </div>

                </div>
              </>
            ))}
          </>
        )}
      </ReactImageUploading>
    </div>
  );
}
