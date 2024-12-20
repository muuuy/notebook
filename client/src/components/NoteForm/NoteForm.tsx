import { useState } from "react";
import { NoteData } from "./NoteFormInterface";

import styles from "./NoteForm.module.scss";

const NoteForm = () => {
  const [formInfo, setFormInfo] = useState<NoteData>({
    title: "",
    folder: "",
    body: "",
  });

  // Update form info when the user types
  const updateFormInfo = (
    event:
      | React.ChangeEvent<HTMLInputElement>
      | React.ChangeEvent<HTMLTextAreaElement>
  ) => {
    event.preventDefault();

    const { name, value } = event.target;

    setFormInfo((prev) => ({ ...prev, [name]: value }));
  };

  // User submits the form
  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    try {
      const res = await fetch("http://localhost:3000/users", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          name: formInfo.title,
          email: formInfo.folder,
        }),
      });

      if (!res.ok) {
        const errorMessage = await res.text();
        console.log("Error creating user: ", errorMessage);
        return;
      }

      console.log("submitted");
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <div className={styles.note_form__container}>
      <h1>CREATE FORM</h1>
      <form className={styles.form__container} onSubmit={handleSubmit}>
        <label htmlFor="title">Title</label>
        <input
          name="title"
          id="title"
          value={formInfo.title}
          onChange={(event) => updateFormInfo(event)}
        />
        <label htmlFor="folder">Folder</label>
        <input
          name="folder"
          id="folder"
          value={formInfo.folder}
          onChange={(event) => updateFormInfo(event)}
        />
        <label htmlFor="body">Body</label>
        <textarea
          name="body"
          id="body"
          value={formInfo.body}
          onChange={(event) => updateFormInfo(event)}
        />
        <button type="submit">Submit</button>
      </form>
    </div>
  );
};

export { NoteForm };
