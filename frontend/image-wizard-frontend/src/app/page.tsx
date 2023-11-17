"use client";
import { useState, useEffect } from "react";

export default function Home() {
  const [res, setRes] = useState<string>();

  useEffect(() => {
    async function fetchData() {
      const response = await fetch("http://image-wizard-api-service:5000");
      const text = await response.text();
      setRes(text);
    }
    fetchData();
  }, []);

  return (
    <main>
      <div>{res}</div>
    </main>
  );
}
