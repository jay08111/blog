import React from "react";
import s from "./Header.module.css";
import NavBar from "@/components/Navbar/Navbar";

const Header: React.FC = () => {
  return (
    <header className={s.blogHeader}>
      <NavBar />
      <div className={s.mainBlogContainer}>
        <h1>THE BLOG </h1>
      </div>
    </header>
  );
};

export default Header;
