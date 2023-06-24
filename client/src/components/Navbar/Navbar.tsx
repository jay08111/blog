import React from "react";
import s from "./Navbar.module.css";
import { MENU } from "./const";

const NavBar: React.FC = () => {
  function renderMenu() {
    return (
      <ul className={s.blogMenu}>
        {MENU.map((el) => (
          <li key={el.id}>{el.menu}</li>
        ))}
      </ul>
    );
  }

  return (
    <nav className={s.navHeader}>
      <div className={s.navInnerContainer}>
        {/* <h1 className={s.navName}>HOYEOUN</h1> */}
        {renderMenu()}
      </div>
    </nav>
  );
};

export default NavBar;
