import type { Component } from "solid-js";
import './header.css'

const Header: Component = () => {
  return (
    <header>
       <div class='logo'>Emstaht</div> 
       <ul>
        <li>CV</li>
        <li>Log in</li>
        <li>Log out</li>
       </ul>
    </header>
  );
};

export default Header;