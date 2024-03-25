CREATE DATABASE preview_week_2;

CREATE TABLE games(
	game_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    genre VARCHAR(100) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    stock INT NOT NULL
);

CREATE TABLE branches(
	branch_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    location VARCHAR(100) NOT NULL
);

CREATE TABLE sales(
	sale_id INT AUTO_INCREMENT PRIMARY KEY,
    game_id INT,
    branch_id INT,
    sale_date DATE NOT NULL,
    quantity INT NOT NULL,
    FOREIGN KEY (game_id)REFERENCES games(game_id),
    FOREIGN KEY (branch_id)REFERENCES branches(branch_id)
);

INSERT INTO games (title, genre, price, stock)
VALUES
("Final Fantasy", "RPG", 59.99, 100),
("FIFA 2024", "Sports", 49.99, 120),
("Doom Eternal", "FPS", 49.99, 80);

INSERT INTO branches (name, location)
VALUES
("Downtown branch", "123 Downtown St"),
("Uptown branch", "456 Uptown Ave");

INSERT INTO sales (game_id, branch_id, sale_date, quantity)
VALUES
(1, 1, "2024-05-01", 2),
(2, 2, "2024-04-02", 3);