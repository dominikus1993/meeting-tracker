import 'fomantic-ui-css/semantic.min.css';
import React, { useState } from 'react';
import logo from './logo.svg';
import './App.css';

interface Meeting {

}

interface MeetingsState {
  readonly started: boolean
  readonly meeting?: Meeting | null | undefined
}



function defaultState(): MeetingsState {
  return { started: false };
}

function StartMeetingForm() {
  const 
}


function Meeting() {
  const [state, setState] = useState(defaultState());
  return (
    <h1>{state.started.toString()}</h1>
  )
}


function App() {
  return (
    <div className="App">
    
      <header className="App-header">
      <Meeting></Meeting>
      </header>
    </div>
  );
}

export default App;
