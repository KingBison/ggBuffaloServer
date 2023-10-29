import { useEffect, useState } from 'react';
import './App.css';
import getServerUrl from './helpers/getServerUrl';
import Login from './components/login';
import GameSelector from './components/gameSelector';



function App() {

  // servers
  const [buffaloServer, setBuffaloServer] = useState(false)

  const [loggedIn, setLoggedIn] = useState(false)
  const [name, setName] = useState("")
  const [selectedGame, setSelectedGame] = useState("")

  useEffect(()=>{
    getServerUrl(setBuffaloServer)
  },[])

  const determineComponent = () => {
    if (!loggedIn) {
      return <Login 
          setLoggedIn={setLoggedIn} 
          setName={setName}
          name = {name}
        />
    }
    if (selectedGame === "") {
      return <GameSelector

      />
    }
  }

  return (
    <>
      {determineComponent()}
    </>
  );
}

export default App;
