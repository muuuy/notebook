import styles from "./Navbar.module.scss";

import { Link } from "react-router-dom";
import { NavbarItemsData } from "./NavbarInterface";

// Navbar items to populate the navbar
const navbarItems: NavbarItemsData[] = [
  { title: "Folders", address: "/folders" },
  { title: "Notes", address: "/create-note" },
  { title: "Gallery", address: "/gallery" },
  { title: "Misc.", address: "/misc" },
];

const Navbar = () => {
  return (
    <div className={styles.navbar__container}>
      <h1>NOTEBOOK</h1>
      <ul className={styles.navbar__menu}>
        {navbarItems.map((item) => {
          return (
            <li>
              <Link to={item.address}>{item.title}</Link>
            </li>
          );
        })}
      </ul>
    </div>
  );
};

export { Navbar };
