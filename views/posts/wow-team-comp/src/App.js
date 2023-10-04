import React, { useState, useEffect } from 'react';
import './App.css';
import { BrowserRouter, Routes, Route } from "react-router-dom";

import {Main} from './Main'

const classImages = {
  Druid: require('./assets/icon_Druid.png'),
  'Demon Hunter' : require('./assets/icon_Demon Hunter.png'),
  Hunter: require('./assets/icon_Hunter.png'),
  Mage: require('./assets/icon_Mage.png'),
  Monk: require('./assets/icon_Monk.png'),
  Paladin: require('./assets/icon_Paladin.png'),
  Priest: require('./assets/icon_Priest.png'),
  Rogue: require('./assets/icon_Rogue.png'),
  Shaman: require('./assets/icon_Shaman.png'),
  Warlock: require('./assets/icon_Warlock.png'),
  Warrior: require('./assets/icon_Warrior.png'),
}


function App() {


  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Main/>}/> 
      </Routes>
    </BrowserRouter>
  );
}

export default App;
