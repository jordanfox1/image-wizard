"use client";
import React, { useState } from "react";
import ReactImageUploading, { ImageListType } from "react-images-uploading";
import './imageUpload.css'
import { useViewportWidth } from "../../hooks/useViewportWidth";
import Dropzone from "./dropzone/dropzone";
import ImageItem from "./image-item/imageItem";

export function ImageUpload() {
  const [images, setImages] = React.useState([]);
  const [loading, setLoading] = React.useState(false);
  const [errors, setErrors] = React.useState({});
  const [desiredFormats, setDesiredFormats] = useState(["webp"]);

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

  const convertToNewFormat = async (
    imageData,
    fileName: string,
    desiredFormat: string = "webp",
    addUpdateIndex
  ) => {
    try {
      setLoading(true);
      setErrors((prevErrors) => {
        const updatedErrors = { ...prevErrors };
        delete updatedErrors[addUpdateIndex];
        return updatedErrors;
      });

      const apiEndpoint =
        process.env.NEXT_PUBLIC_API_ENDPOINT || "https://img-switch-api-production.up.railway.app/api";

      const formData = new FormData();
      formData.append("image", imageData);
      formData.append("fileName", fileName);

      const response = await fetch(
        `${apiEndpoint}/convert?format=${desiredFormat}`,
        {
          method: "POST",
          body: formData,
        }
      );

      if (!response.ok) {
        if (response.status == 400) {
          throw new Error("Invalid request - Input and output formats can not be the same - supported formats are WEBP, JPEG, GIF and BMP.");
        }

        throw new Error("Something went wrong, please check your file size and formats and try again later.");
      }

      const responseData = await response.json();

      // Access the data URL directly from the response
      const dataURL = responseData.dataURL;
      const newFileName = responseData.fileName;

      setImages((prevImages) => {
        const updatedImages = [...prevImages];
        const newFile = new File(
          [updatedImages[addUpdateIndex].file],
          newFileName
        );
        updatedImages[addUpdateIndex].dataURL = dataURL;
        updatedImages[addUpdateIndex].file = newFile;
        updatedImages[addUpdateIndex].converted = true;
        return updatedImages;
      });
    } catch (error) {

      console.error(error.message);
      if (error.message == "Failed to fetch") {
        error.message = "Something went wrong, please check your file size and formats and try again later."
      }
      setErrors((prevErrors) => ({
        ...prevErrors,
        [addUpdateIndex]: error.message,
      }));

    } finally {
      setLoading(false);
    }
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
          dragProps,
        }) => (
          <>
            <Dropzone
              isDragging={isDragging}
              dragProps={dragProps}
              onImageUpload={onImageUpload}
              viewportWidth={vw}
            />

            {imageList.map((image, index) => (
              <ImageItem
                key={index}
                image={image}
                index={index}
                errors={errors}
                onImageRemove={onImageRemove}
                convertToNewFormat={convertToNewFormat}
                desiredFormats={desiredFormats}
                vw={vw}
                setDesiredFormats={setDesiredFormats}
                loading={loading}
              />
            ))}
          </>
        )}
      </ReactImageUploading>
    </div>
  );
}
