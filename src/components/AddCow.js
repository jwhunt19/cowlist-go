import React, { useState } from "react";

const AddCow = ({ addCow }) => {
  const [addingCow, setAddingCow] = useState(false);
  const [warning, setWarning] = useState(false);

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
      Age: parseInt(age.value),
      Color: color.value,
      Healthy: healthy.value === "true" ? true : false,
    };
    addCow(newCow);
    setAddingCow(false);
  };

  return (
    <div>
      {addingCow ? (
        <form
          style={{
            display: "flex",
            flexDirection: "column",
            width: "200px",
            margin: "auto",
          }}
          onSubmit={handleSubmit}
        >
          <label htmlFor="name">Name:</label>
          <input type="text" id="name" name="name" required />
          <label htmlFor="age">Age:</label>
          <input type="number" id="age" name="age" required />
          <label htmlFor="color">Color:</label>
          <input type="text" id="color" name="color" required />
          <label htmlFor="healthy">Healthy:</label>
          <input type="text" id="healthy" name="healthy" required />
          <button type="submit">Submit</button>
          <button onClick={() => setAddingCow(false)}>Cancel</button>
          {warning && (
            <p style={{ color: "red" }}>Healthy must be true or false</p>
          )}
        </form>
      ) : (
        <button onClick={() => setAddingCow(true)}>Add Cow</button>
      )}
    </div>
  );
};

export default AddCow;
