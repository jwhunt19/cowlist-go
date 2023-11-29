import './App.css';
import axios from 'axios';
import Cow from "./components/Cow.js"

function App() {

  const test = () => {
    axios.get('/').then((data) => console.log(data))
  }
  

  return (
    <div className="App">
      <p>um, can I have a pizza?</p>
      <Cow></Cow>
      <button onClick={test}>Test API</button>
    </div>
  );
}

export default App;
