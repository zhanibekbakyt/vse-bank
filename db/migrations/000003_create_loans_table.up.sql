CREATE TABLE loans (
    id SERIAL PRIMARY KEY,
    amount FLOAT NOT NULL,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    bank_id INT REFERENCES banks(id) ON DELETE CASCADE,
    loan_type_id INT
);
