import "./App.css";
import axios from "axios";
import { useEffect, useState } from "react";
import Cowlist from "./components/Cowlist.js";
import AddCow from "./components/AddCow.js";

function App() {
  const [cows, setCows] = useState([]);

  // get all cows from database
  const getCows = async () => {
    const { data } = await axios.get("http://localhost:8080/getallcows");

    if (data === null) {
      setCows([]);
      return;
    }

    setCows(data);
  };

  // get all cows from database on page load
  useEffect(() => {
    getCows();
  }, []);

  // update cow in database
  const updateCow = async (cow) => {
    axios
      .put("http://localhost:8080/updatecow", cow)
      .then(() => {
        getCows();
      })
      .catch((err) => {
        console.log("Error!!: ", err);
      });
  };

  // delete cow from database
  const deleteCow = async (id) => {
    axios
      .delete(`http://localhost:8080/deletecow/${id}`)
      .then(() => {
        getCows();
      })
      .catch((err) => {
        console.log("Error!!: ", err);
      });
  };

  // add cow to database
  const addCow = (cow) => {
    axios
      .post("http://localhost:8080/addcow", {
        Name: cow.Name,
        Age: cow.Age,
        Color: cow.Color,
        Healthy: cow.Healthy,
      })
      .then(() => {
        getCows();
      })
      .catch((err) => {
        console.log("Error!!: ", err);
      });
  };

  return (
    <div className="App">
      <h1>Cowlist</h1>
      <AddCow addCow={addCow} />
      <Cowlist updateCow={updateCow} deleteCow={deleteCow} cows={cows} />
    </div>
  );
}

export default App;
