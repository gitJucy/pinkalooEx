import React, { useState } from 'react';
import './App.css';
import AppList from './components/AppList';
import SearchComp from './components/SearchComp';
const App = () => {
  let appArray = [];
  let [apps, setApps] = useState([]);
  function fetchApps() {
    if (!apps.length) {
      fetch(`/api/v1/apps`)
        .then(res => res.json())
        .then(data => {
          data.map(app => {
            appArray.push(app);
            return appArray;
          });
          setApps((apps = appArray));
        })
        .catch(err => {
          console.log(
            'Backend does not contain any app meta data.',
            err.message
          );
        });
    }
  }
  function clearApps() {
    setApps((apps = []));
  }
  return (
    <div className='App'>
      <h1>Pinkaloo App Search</h1>
      <button onClick={fetchApps}>Show All Apps</button>
      <button onClick={clearApps}>Clear App List</button>
      <SearchComp />
      <AppList apps={apps} />
    </div>
  );
};

export default App;
