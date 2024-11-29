import styles from "./Navbar.module.scss";

import { Link } from "react-router-dom";

const Navbar = () => {
  return (
    <div className={styles.navbar__container}>
      <h1>NOTEBOOK</h1>
      <ul className={styles.navbar__menu}>
        <li>
          <Link to={"/folders"}>Browse Folders</Link>
        </li>
        <li>
          <Link to={"/create-note"}>Create a note</Link>
        </li>
        <li>Misc.</li>
      </ul>
    </div>
  );
};

export { Navbar };
