import React, { useState, useEffect } from 'react';
import parseArray from './helper/result_helper';
import './App.css';


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

  const [wowClass, setWowClass] = useState('');
  const [wowSpec, setWowSpec] = useState('');
  const [result, setResult] = useState('');
  const [currentClass, setCurrentClass] = useState('');
  const [classesInApiResponse, setClassesInApiResponse] = useState([]);


  const handleGetTeamData = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch(`http://localhost:8080/getBestTeam?class=${wowClass}&spec=${wowSpec}`);

      if (!response.ok) {
        throw new Error('Network response was not ok');
      }

      const data = await response.json();
      const formattedData = JSON.stringify(data, null, 2);
      setResult(formattedData);
      console.log(formattedData);
      const classNames = data.members.map((member) => member.class);
      setClassesInApiResponse(classNames);
      console.log(classImages);
    } catch (error) {
      console.error('Fetch error:', error);
    }
  };

   useEffect(() => {
    try {
      const jsonData = JSON.parse(result);
      setCurrentClass(jsonData.members[0].class); 
    } catch (error) {
      console.error('JSON parsing error:', error);
    }
  }, [result]);

  return (
    <div className="banner">
      <p className="title">Get Best Team</p>
      <form className="label">
        <label htmlFor="class">Class:</label>
        <input
          type="text"
          id="class"
          name="class"
          required
          value={wowClass}
          onChange={(e) => setWowClass(e.target.value)}
        /><br /><br />
        <label htmlFor="spec">Spec:</label>
        <input
          type="text"
          id="spec"
          name="spec"
          required
          value={wowSpec}
          onChange={(e) => setWowSpec(e.target.value)}
        /><br /><br />
        
        <input
          type="submit"
          id="getTeamButton"
          value="Submit"
          onClick={handleGetTeamData}
        />
      </form>
      <div className="class-images-container">
        {classesInApiResponse.map((className) => (
          <img
            key={className}
            src={classImages[className]}
            alt={className}
            className="class-image"
          />
        ))}
      </div>
    </div>
  );
}

export default App;
