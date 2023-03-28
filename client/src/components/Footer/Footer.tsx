import type { Component } from "solid-js";
import './footer.css'

const Footer: Component = () => {
  //TODO: this text has to be animated.
  const terminalText = 'yourmachine ~ $'
  const terminalInput = 'by Anton Podkur'
  return (
    <footer>
        <div class='terminal-text'>{terminalText}</div>
        <div class='terminal-input'>{terminalInput}</div>
    </footer>
  );
};

export default Footer;