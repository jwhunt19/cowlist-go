CREATE TABLE cows (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50),
  age INTEGER,
  color VARCHAR(50),
  healthy BOOLEAN
);

INSERT INTO cows (name, age, color, healthy) VALUES ('Miltank', 24, 'Pink', true);
