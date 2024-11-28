import { useState } from "react";
import { NoteData } from "./NoteFormInterface";

const NoteForm = () => {
  const [formInfo, setFormInfo] = useState<NoteData>({
    title: "",
    folder: "",
    body: "",
  });

  const updateFormInfo = (event: React.ChangeEvent<HTMLInputElement>) => {
    event.preventDefault();

    const { name, value } = event.target;

    setFormInfo((prev) => ({ ...prev, [name]: value }));
  };

  return (
    <div>
      <h1>CREATE FORM</h1>
      <form>
        <label htmlFor="title">Title</label>
        <input
          name="title"
          value={formInfo.title}
          onChange={(event) => updateFormInfo(event)}
        />
        <label htmlFor="folder">Folder</label>
        <input
          name="folder"
          value={formInfo.folder}
          onChange={(event) => updateFormInfo(event)}
        />
        <label htmlFor="body">Body</label>
        <input
          name="body"
          value={formInfo.body}
          onChange={(event) => updateFormInfo(event)}
        />
      </form>
    </div>
  );
};

export { NoteForm };
