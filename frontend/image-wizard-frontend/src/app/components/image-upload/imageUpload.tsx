"use client";
import Image from "next/image";
import React, { useState } from "react";
import ReactImageUploading, { ImageListType } from "react-images-uploading";
import { Button, Chip, Text } from '@mantine/core';
import { IconPhoto, IconDownload, IconUpload, IconTrash } from '@tabler/icons-react';
import './imageUpload.css'
import { useViewportWidth } from "../../hooks/useViewportWidth";

export function ImageUpload() {
  const [images, setImages] = React.useState([]);
  const [loading, setLoading] = React.useState(false);
  const [errors, setErrors] = React.useState({});
  const [desiredFormats, setDesiredFormats] = useState(['webp']);

  const maxNumber = 69;
  const vw = useViewportWidth();

  const onChange = (
    imageList: ImageListType,
    addUpdateIndex: number[] | undefined
  ) => {
    setErrors((prevErrors) => {
      const updatedErrors = { ...prevErrors };
      if (addUpdateIndex !== undefined) {
        delete updatedErrors[addUpdateIndex[0]];
      }
      return updatedErrors;
    });
    setImages(imageList as never[]);
  };

  const convertToNewFormat = async (imageData, fileName: string, desiredFormat: string, addUpdateIndex) => {
    try {
      setLoading(true);
      setErrors((prevErrors) => {
        const updatedErrors = { ...prevErrors };
        delete updatedErrors[addUpdateIndex];
        return updatedErrors;
      });

      const apiEndpoint = process.env.NEXT_PUBLIC_API_ENDPOINT || 'http://img-switch.local/api';

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
      setErrors((prevErrors) => ({ ...prevErrors, [addUpdateIndex]: 'Error converting image. Please try again.' }));
    } finally {
      setLoading(false);
    }
  };

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
            <div className="drop-zone sticky" style={isDragging ? { color: "red" } : undefined} {...dragProps}>
              <Button size={vw > 1023 ? 'xl' : 'sm'} rightSection={<IconUpload size={14} />} className="btn-large" onClick={onImageUpload} >
                Select Files
              </Button>
              <Text size="md" c="#3f51b5">Click or drop files here</Text>
            </div>

            {imageList.map((image, index) => (
              <>
                <div className="format-select-btn-container">
                  <Chip.Group defaultValue="webp" multiple={false} value={desiredFormats[index]} onChange={(value) => handleChipChange(value, index)}>
                    <Chip radius="xs" size={vw > 1023 ? 'xl' : 'sm'} className="chip" value='webp'>WEBP</Chip>
                    <Chip radius="xs" size={vw > 1023 ? 'xl' : 'sm'} className="chip" value='png'>PNG</Chip>
                    <Chip radius="xs" size={vw > 1023 ? 'xl' : 'sm'} className="chip" value='jpeg'>JPEG</Chip>
                    <Chip radius="xs" size={vw > 1023 ? 'xl' : 'sm'} className="chip" value='gif'>GIF</Chip>
                    <Chip radius="xs" size={vw > 1023 ? 'xl' : 'sm'} className="chip" value='bmp'>BMP</Chip>
                  </Chip.Group>
                </div>

                <div key={index} className="image-item">
                  <figure className="image-figure">
                    <Image className="image" src={image.dataURL} alt="your uploaded image" width={280} height={160} layout="responsive" />
                    <figcaption>
                      <Text size="md">{image.file?.name} </Text>
                    </figcaption>

                    {errors[index] && (
                      <Text size="md" c="red">
                        {errors[index]}
                      </Text>
                    )}

                  </figure>

                  <div className="image-item__btn-wrapper" >
                    <Button className="btn-med" size={vw > 1023 ? 'xl' : 'sm'} rightSection={<IconPhoto size={14} />} onClick={() => convertToNewFormat(image.dataURL, image.file.name, desiredFormats[index], index)} loading={loading}>
                      Convert to {desiredFormats[index]}
                    </Button>
                    <Button className="btn-med" size={vw > 1023 ? 'xl' : 'sm'} rightSection={<IconTrash size={14} />} onClick={() => onImageRemove(index)} loading={loading}>Remove</Button>
                    <Button className="btn-med" size={vw > 1023 ? 'xl' : 'sm'} component="a" rightSection={<IconDownload size={14} />} href={image.dataURL} download={image.file.name} loading={loading}>
                      Download
                    </Button>
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
