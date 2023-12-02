import React, { useState } from "react";

const Cow = ({ updateCow, cow }) => {
  const [editing, setEditing] = useState(false);

  const editCow = () => {
    setEditing(!editing);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    const { name, age, color, healthy } = e.target;
    const newCow = {
      Name: name.value,
      Age: age.value,
      Color: color.value,
      Healthy: healthy.value,
      Id: cow.Id,
    };
    updateCow(newCow);
    editCow();
  }

  return (
    <div
      style={{
        border: "1px solid black",
        margin: "20px",
        padding: "20px",
        width: "200px",
      }}
    >
      {!editing ? (
        <>
          <h2>name: {cow.Name}</h2>
          <p>age: {cow.Age}</p>
          <p>color: {cow.Color}</p>
          <p>healthy: {cow.Healthy.toString()}</p>
          <button onClick={editCow}>Edit</button>{" "}
        </>
      ) : (
        <form onSubmit={(e) => handleSubmit(e)}>
          <input type="text" name="name" placeholder="name" defaultValue={cow.Name} />
          <input type="number" name="age" placeholder="age" defaultValue={cow.Age}/>
          <input type="text" name="color" placeholder="color" defaultValue={cow.Color} />
          {/* todo - fix healthy input to only accept true/false (bool) */}
          <input type="text" name="healthy" placeholder="healthy" defaultValue={cow.Healthy} />
          <button type="submit">Submit</button>
        </form>
      )}
    </div>
  );
};

export default Cow;
