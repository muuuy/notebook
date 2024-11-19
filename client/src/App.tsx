import { Navbar } from "./components/Navbar/Navbar";

import { useState } from "react";

function App() {
  const [count, setCount] = useState(0);

  return (
    <>
      <div>
        <Navbar />
        <h1>NOTEBOOK</h1>
        <p>Things I learn.</p>
      </div>
    </>
  );
}

export default App;
