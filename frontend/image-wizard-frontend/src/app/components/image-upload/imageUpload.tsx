"use client";
import { useState } from "react";

export default function ImageUpload() {
  const [image, setImage] = useState(null);

  const handleUpload = async () => {
    const formData = new FormData();
    formData.append("image", image);

    const response = await fetch("http://localhost:5000/api/upload", {
      method: "POST",
      body: formData,
    });

    const data = await response.json();
    console.log(data);
  };

  return (
    <div>
      <input type="file" onChange={(e) => setImage(e.target.files[0])} />
      <button onClick={handleUpload}>Upload</button>
    </div>
  );
}
