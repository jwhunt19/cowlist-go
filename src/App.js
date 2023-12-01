import "./App.css";
import axios from "axios";
import Cow from "./components/Cow.js";

function App() {
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
      <p>um, can I have a pizza?</p>
      <Cow></Cow>
      <button onClick={test}>Test API</button>
    </div>
  );
}

export default App;
