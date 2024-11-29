import { Navbar } from "./components/Navbar/Navbar";
import { Folders } from "./components/Folders/Folders";
import { Home } from "./components/Home/Home";
import { NoteForm } from "./components/NoteForm/NoteForm";

import { Route, Routes } from "react-router-dom";

function App() {

  return (
    <>
      <div>
        <Navbar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/folders" element={<Folders />} />
          <Route path="/create-note" element={<NoteForm />} />
        </Routes>
      </div>
    </>
  );
}

export default App;
