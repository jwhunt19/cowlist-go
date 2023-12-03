import React, { useState } from "react";

const Cow = ({ updateCow, deleteCow, cow }) => {
  const [editing, setEditing] = useState(false);
  const [warning, setWarning] = useState(false);

  const editCow = () => {
    setEditing(!editing);
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    const { name, age, color, healthy } = e.target;

    if (healthy.value !== "true" && healthy.value !== "false") {
      setWarning(true);
      setTimeout(() => {
        setWarning(false);
      }, 5000);

      return;
    }

    const newCow = {
      Name: name.value,
      Age: age.value,
      Color: color.value,
      Healthy: healthy.value,
      Id: cow.Id,
    };
    updateCow(newCow);
    editCow();
  };

  const handleDelete = () => {
    deleteCow(cow.Id);
  };

  return (
    <div
      className="cow"
      style={{
        border: "2px solid black",
        borderRadius: "5px",
        margin: "10px",
        padding: "5px",
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
          <button onClick={handleDelete} style={{ backgroundColor: "#a00000" }}>
            Delete
          </button>
        </>
      ) : (
        <form onSubmit={(e) => handleSubmit(e)}>
          <input
            type="text"
            name="name"
            placeholder="name"
            defaultValue={cow.Name}
            required
          />
          <input
            type="number"
            name="age"
            placeholder="age"
            defaultValue={cow.Age}
            required
          />
          <input
            type="text"
            name="color"
            placeholder="color"
            defaultValue={cow.Color}
            required
          />
          <input
            type="text"
            name="healthy"
            placeholder="healthy"
            defaultValue={cow.Healthy}
            required
          />
          {warning && (
            <p style={{ color: "red" }}>Healthy must be true or false</p>
          )}
          <button type="submit">Submit</button>
        </form>
      )}
    </div>
  );
};

export default Cow;
