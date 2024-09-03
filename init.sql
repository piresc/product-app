CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    variety TEXT,
    rating NUMERIC(2, 1),
    stock INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO products (name, description, price, variety, rating, stock) VALUES
('Product 1', 'Description for Product 1', 19.99, 'Variety A', 4.5, 100),
('Product 2', 'Description for Product 2', 29.99, 'Variety B', 4.0, 50),
('Product 3', 'Description for Product 3', 9.99, 'Variety C', 3.5, 200),
('Product 4', 'Description for Product 4', 49.99, 'Variety D', 5.0, 10),
('Product 5', 'Description for Product 5', 15.99, 'Variety E', 4.2, 75);