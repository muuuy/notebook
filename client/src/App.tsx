import { Navbar } from "./components/Navbar/Navbar";
import { Folders } from "./components/Folders/Folders";
import { Home } from "./components/Home/Home";
import { NoteForm } from "./components/NoteForm/NoteForm";
import { GalleryPage } from "./components/GalleryPage/GalleryPage";

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
          <Route path="/gallery" element={<GalleryPage />} />
        </Routes>
      </div>
    </>
  );
}

export default App;
