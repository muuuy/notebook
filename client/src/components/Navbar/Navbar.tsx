import styles from "./Navbar.module.scss";

const Navbar = () => {
  return (
    <div className={styles.navbar__container}>
      <h1>NOTEBOOK</h1>
      <ul className={styles.navbar__menu}>
        <li>Browse Folders</li>
        <li>Create Form</li>
        <li>Misc.</li>
      </ul>
    </div>
  );
};

export { Navbar };
