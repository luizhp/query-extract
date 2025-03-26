-- Table creation

-- Clients Table
CREATE TABLE clients (
  client_id INT AUTO_INCREMENT PRIMARY KEY,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  email VARCHAR(100) UNIQUE NOT NULL,
  phone VARCHAR(20),
  address VARCHAR(255),
  city VARCHAR(50),
  state VARCHAR(50),
  zip_code VARCHAR(10),
  country VARCHAR(50)
);

-- Suppliers Table
CREATE TABLE suppliers (
  supplier_id INT AUTO_INCREMENT PRIMARY KEY,
  company_name VARCHAR(100) NOT NULL,
  contact_name VARCHAR(100),
  contact_title VARCHAR(50),
  address VARCHAR(255),
  city VARCHAR(50),
  state VARCHAR(50),
  zip_code VARCHAR(10),
  country VARCHAR(50),
  phone VARCHAR(20),
  fax VARCHAR(20)
);

-- Sellers Table
CREATE TABLE sellers (
  seller_id INT AUTO_INCREMENT PRIMARY KEY,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  email VARCHAR(100) UNIQUE NOT NULL,
  phone VARCHAR(20),
  hire_date DATE,
  commission_rate DECIMAL(5,2)
);

-- Products Table
CREATE TABLE products (
  product_id INT AUTO_INCREMENT PRIMARY KEY,
  product_name VARCHAR(100) NOT NULL,
  supplier_id INT,
  category VARCHAR(50),
  unit_price DECIMAL(10,2),
  units_in_stock INT,
  units_on_order INT,
  reorder_level INT,
  discontinued TINYINT(1) DEFAULT 0,
  FOREIGN KEY (supplier_id) REFERENCES suppliers(supplier_id)
);

-- Orders Table
CREATE TABLE orders (
  order_id INT AUTO_INCREMENT PRIMARY KEY,
  client_id INT NOT NULL,
  seller_id INT,
  order_date DATE NOT NULL,
  required_date DATE,
  shipped_date DATE,
  ship_via INT,
  freight DECIMAL(10,2),
  ship_name VARCHAR(100),
  ship_address VARCHAR(255),
  ship_city VARCHAR(50),
  ship_state VARCHAR(50),
  ship_zip_code VARCHAR(10),
  ship_country VARCHAR(50),
  FOREIGN KEY (client_id) REFERENCES clients(client_id),
  FOREIGN KEY (seller_id) REFERENCES sellers(seller_id)
);

-- Order Lines Table
CREATE TABLE order_lines (
  order_line_id INT AUTO_INCREMENT PRIMARY KEY,
  order_id INT NOT NULL,
  product_id INT NOT NULL,
  quantity INT NOT NULL,
  unit_price DECIMAL(10,2) NOT NULL,
  discount DECIMAL(4,2) DEFAULT 0.00,
  FOREIGN KEY (order_id) REFERENCES orders(order_id),
  FOREIGN KEY (product_id) REFERENCES products(product_id)
);

-- Index creation
CREATE INDEX idx_supplier_id ON products(supplier_id);
CREATE INDEX idx_client_id ON orders(client_id);
CREATE INDEX idx_seller_id ON orders(seller_id);
CREATE INDEX idx_order_id ON order_lines(order_id);
CREATE INDEX idx_product_id ON order_lines(product_id);

-- Sample Data Insertion

-- Data for Clients
INSERT INTO clients (first_name, last_name, email, phone, address, city, state, zip_code, country)
VALUES
('John', 'Doe', 'john.doe@example.com', '123-456-7890', '123 Main St', 'Anytown', 'CA', '12345', 'USA'),
('Jane', 'Smith', 'jane.smith@example.com', '987-654-3210', '456 Elm St', 'Othertown', 'NY', '67890', 'USA'),
('Alice', 'Johnson', 'alice.johnson@example.com', '555-555-5555', '789 Oak St', 'Sometown', 'TX', '54321', 'USA');

-- Data for Suppliers
INSERT INTO suppliers (company_name, contact_name, contact_title, address, city, state, zip_code, country, phone, fax)
VALUES
('Supplier A Inc.', 'Anna Smith', 'CEO', '100 Supplier St', 'Supplier City', 'CA', '11111', 'USA', '111-111-1111', '111-111-1112'),
('Supplier B Ltd.', 'Bob Brown', 'Manager', '200 Supplier Ave', 'Another City', 'NY', '22222', 'USA', '222-222-2222', '222-222-2223');

-- Data for Sellers
INSERT INTO sellers (first_name, last_name, email, phone, hire_date, commission_rate)
VALUES
('Mike', 'Davis', 'mike.davis@example.com', '333-333-3333', '2020-01-01', 0.05),
('Emily', 'Wilson', 'emily.wilson@example.com', '444-444-4444', '2021-06-15', 0.07);

-- Data for Products
INSERT INTO products (product_name, supplier_id, category, unit_price, units_in_stock, units_on_order, reorder_level, discontinued)
VALUES
('Product X', 1, 'Category 1', 19.99, 100, 0, 10, 0),
('Product Y', 1, 'Category 2', 29.99, 50, 20, 15, 0),
('Product Z', 2, 'Category 1', 9.99, 200, 0, 20, 0);

-- Data for Orders
INSERT INTO orders (client_id, seller_id, order_date, required_date, shipped_date, ship_via, freight, ship_name, ship_address, ship_city, ship_state, ship_zip_code, ship_country)
VALUES
(1, 1, '2023-01-01', '2023-01-10', '2023-01-05', 1, 10.00, 'John Doe', '123 Main St', 'Anytown', 'CA', '12345', 'USA'),
(2, 2, '2023-02-15', '2023-02-25', '2023-02-20', 2, 15.00, 'Jane Smith', '456 Elm St', 'Othertown', 'NY', '67890', 'USA'),
(3, 1, '2023-03-01', '2023-03-10', NULL, 3, 0.00, 'Alice Johnson', '789 Oak St', 'Sometown', 'TX', '54321', 'USA');

-- Data for Order Lines
INSERT INTO order_lines (order_id, product_id, quantity, unit_price, discount)
VALUES
(1, 1, 2, 19.99, 0.00),
(1, 2, 1, 29.99, 0.00),
(2, 3, 5, 9.99, 0.00),
(3, 1, 3, 19.99, 0.00),
(3, 3, 10, 9.99, 0.00);