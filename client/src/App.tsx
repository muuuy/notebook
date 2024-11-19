import { Navbar } from "./components/Navbar/Navbar";
import { Folders } from "./components/Folders/Folders";
import { Home } from "./components/Home/Home";

import { useState } from "react";
import { Route, Routes } from "react-router-dom";

function App() {
  const [count, setCount] = useState(0);

  return (
    <>
      <div>
        <Navbar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/folders" element={<Folders />} />
        </Routes>
      </div>
    </>
  );
}

export default App;
