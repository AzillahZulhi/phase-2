CREATE DATABASE IF NOT EXISTS superhero_database;

CREATE TABLE IF NOT EXISTS Heroes (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(100) NOT NULL,
    Universe VARCHAR(50) NOT NULL,
    Skill VARCHAR(255) NOT NULL,
    ImageURL VARCHAR(255) NOT NULL
);

INSERT INTO Heroes (Name, Universe, Skill, ImageURL) VALUES
('Superman', 'DC', 'Super strength, flight, heat vision', 'https://example.com/superman.jpg'),
('Spider-Man', 'Marvel', 'Wall-crawling, web-slinging, spider sense', 'https://example.com/spiderman.jpg'),
('Wonder Woman', 'DC', 'Super strength, flight, divine wisdom', 'https://example.com/wonderwoman.jpg');

CREATE TABLE IF NOT EXISTS Villain (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(100) NOT NULL,
    Universe VARCHAR(50) NOT NULL,
    ImageURL VARCHAR(255) NOT NULL
);

INSERT INTO Villain (Name, Universe, ImageURL) VALUES
('Joker', 'DC', 'https://example.com/joker.jpg'),
('Thanos', 'Marvel', 'https://example.com/thanos.jpg'),
('Lex Luthor', 'DC', 'https://example.com/lexluthor.jpg');