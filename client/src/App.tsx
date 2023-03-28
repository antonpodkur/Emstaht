import type { Component } from 'solid-js';
import Header from './components/Header/Header'
import Footer from './components/Footer/Footer'
import Welcome from './pages/Welcome/Welcome'
import MainBackground from './components/backgrounds/Main/MainBackground';

const App: Component = () => {
  return (
    <div id='app'>
      <MainBackground/>
      <Header/>
      <Welcome/>
      <Footer/>
    </div>
  );
};

export default App;
