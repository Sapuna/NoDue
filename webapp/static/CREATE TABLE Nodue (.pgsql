CREATE TABLE Nodue (
     id SERIAL PRIMARY KEY,
  enrollment_number INT,
  name VARCHAR(255) NOT NULL,
  year INT NOT NULL,
  status VARCHAR(50) NOT NULL
);

INSERT INTO due (enrollment_number, name, year, status)
VALUES
  (123456, 'John Doe', 2023, 'Pending'),
  (789012, 'Jane Smith', 2022, 'Paid' );