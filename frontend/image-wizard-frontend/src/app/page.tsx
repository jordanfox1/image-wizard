"use server";
export default async function Home() {
  const response = await fetch("http://image-wizard-api-srvc:5000");
  const text = await response.text();

  return (
    <main>
      <div>{text}</div> 
    </main>
  );

}
