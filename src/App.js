import "./App.css";
import axios from "axios";
import { useEffect, useState } from "react";
import Cowlist from "./components/Cowlist.js";
import AddCow from "./components/AddCow.js";

function App() {
  const [cows, setCows] = useState([]);

  // get all cows from database
  const getCows = async () => {
    const { data } = await axios.get("http://localhost:8080/getallcows", {
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
    });

    // if cows is null, set to empty array
    if (data === null) {
      setCows([]);
      return;
    }

    // set cows to data
    setCows(data);
  };

  // get all cows from database on page load
  useEffect(() => {
    getCows();
  }, []);

  // update cow in database
  const updateCow = async (cow) => {
    try {
      const { data } = await axios.put(
        "http://localhost:8080/updatecow",
        {
          Name: cow.Name,
          Age: cow.Age,
          Color: cow.Color,
          Healthy: cow.Healthy,
          Id: cow.Id,
        },
        {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
        }
      );

      getCows();
      console.log(data); // todo delete
    } catch (error) {
      console.log("Error!!: ", error);
    }
  };

  // delete cow from database
  const deleteCow = async (id) => {
    try {
      const { data } = await axios.delete(
        `http://localhost:8080/deletecow/${id}`,
        {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
        }
      );

      getCows();
      console.log(data);
    } catch (error) {
      console.log("Error!!: ", error);
    }
  };

  // add cow to database
  const addCow = (cow) => {
    axios
      .post(
        "http://localhost:8080/addcow",
        {
          Name: cow.Name,
          Age: cow.Age,
          Color: cow.Color,
          Healthy: cow.Healthy,
        },
        {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
          },
        }
      )
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
