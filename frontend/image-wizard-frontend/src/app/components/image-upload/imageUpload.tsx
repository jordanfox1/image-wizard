"use client";
import Image from "next/image";
import React, { useState } from "react";
import ReactImageUploading, { ImageListType } from "react-images-uploading";
import { AspectRatio, Button, Chip, ChipGroup } from '@mantine/core';
import { IconPhoto, IconDownload, IconTrash } from '@tabler/icons-react';
import './imageUpload.css'

export function ImageUpload() {
  const [images, setImages] = React.useState([]);
  const [loading, setLoading] = React.useState(false);
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
      setLoading(true);
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
    } finally {
      setLoading(false);
    }
  };

  const [desiredFormats, setDesiredFormats] = useState(['webp']);

  const handleChipChange = (value, index) => {
    setDesiredFormats((prevFormats) => {
      const updatedFormats = [...prevFormats];
      updatedFormats[index] = value;
      return updatedFormats;
    });
  };

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
                  <Chip.Group defaultValue="webp" multiple={false} value={desiredFormats[index]} onChange={(value) => handleChipChange(value, index)}>
                    <Chip radius="0" value='webp'>WEBP</Chip>
                    <Chip radius="0" value='png'>PNG</Chip>
                    <Chip radius="0" value='jpeg'>JPEG</Chip>
                    <Chip radius="0" value='gif'>GIF</Chip>
                    <Chip radius="0" value='bmp'>BMP</Chip>
                  </Chip.Group>
                </div>

                <div key={index} className="image-item">
                  <figure className="image-figure">
                    <Image className="image" src={image.dataURL} alt="your uploaded image" width={280} height={160} />
                    <figcaption>{image.file?.name}</figcaption>
                  </figure>




                  <Button.Group className="image-item__btn-wrapper" orientation="vertical">
                    <Button rightSection={<IconPhoto size={14} />} onClick={() => convertToNewFormat(image.dataURL, image.file.name, desiredFormats[index], index)} loading={loading}>
                      Convert to {desiredFormats[index]}
                    </Button>
                    <Button rightSection={<IconTrash size={14} />} onClick={() => onImageRemove(index)} loading={loading}>Remove</Button>
                    <Button component="a" rightSection={<IconDownload size={14} />} href={image.dataURL} download={image.file.name} loading={loading}>
                      Download
                    </Button>
                  </Button.Group>

                </div>

              </>
            ))}
          </>
        )}
      </ReactImageUploading>
    </div>
  );
}
