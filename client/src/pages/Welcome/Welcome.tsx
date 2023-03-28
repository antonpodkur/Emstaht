import type { Component } from "solid-js";
import './welcome.css'

const Welcome : Component = () => {
    return (
        <main>
            <div class="welcome-text">Hello, let's share your skills!</div>
            <button>Create CV</button>
        </main>
    )
}

export default Welcome;