import type { Component } from 'solid-js';
import Header from './components/Header/Header'
import Footer from './components/Footer/Footer'
import Welcome from './pages/Welcome/Welcome'
import MainBackground from './components/backgrounds/Main/MainBackground';

const App: Component = () => {
  return (
    <div id='app'>
      {/*<MainBackground/>*/}
      <Header/>
        <div class="main">
            <Welcome/>
        </div>
      <Footer/>
    </div>
  );
};

export default App;
