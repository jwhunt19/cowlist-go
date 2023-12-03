import "./App.css";
import axios from "axios";
import { useEffect, useState } from "react";
import Cowlist from "./components/Cowlist.js";

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
      console.log(data); // todo delete
    } catch (error) {
      console.log("Error!!: ", error);
    }
  }

  // test api
  const test = async () => {
    const { data } = await axios.get("http://localhost:8080/getallcows", {
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
    });

    console.log(data);
    // axios
    //   .post(
    //     "http://localhost:8080/addcow",
    //     {
    //       name: "Miltank",
    //       age: 24,
    //       color: "Pink",
    //       healthy: true,
    //     },
    //     {
    //       headers: {
    //         "Content-Type": "application/x-www-form-urlencoded",
    //       },
    //     }
    //   )
    //   .then(function (response) {
    //     console.log(response.data);
    //   })
    //   .catch(function (error) {
    //     console.log(error);
    //   });
  };

  return (
    <div className="App">
      <h1>Cowlist</h1>
      <Cowlist updateCow={updateCow} deleteCow={deleteCow} cows={cows} />
      <button onClick={test}>Test API</button>
    </div>
  );
}

export default App;
