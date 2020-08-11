import React, { useState } from 'react';
import logo from './logo.svg';
import './App.css';

interface MeetingsState {
  readonly started: boolean
  readonly meeting?: Meeting?
}

interface Meeting {

}

function defaultState(): MeetingsState {
  return { started: false };
}

function A() {
  const [state, setState] = useState(defaultState());
  return (

  )
}


function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
