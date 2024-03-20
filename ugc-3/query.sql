CREATE TABLE inventories (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    stock INT NOT NULL,
    description TEXT NOT NULL,
    status ENUM('broken', 'active') NOT NULL
);

INSERT INTO inventories (name, stock, description, status) VALUES
('Weapon A', 100, 'Details for Weapon A', 'active'),
('Weapon B', 75, 'Details for Weapon B', 'active'),
('Weapon C', 50, 'Details for Weapon C', 'broken'),
('Weapon D', 200, 'Details for Weapon D', 'active'),
('Weapon E', 30, 'Details for Weapon E', 'active');