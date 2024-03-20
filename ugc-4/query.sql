CREATE TABLE criminal_report (
    id INT PRIMARY KEY AUTO_INCREMENT,
    hero_id INT,
    villain_id INT NOT NULL,
    description TEXT NOT NULL,
    date_of_incident DATE NOT NULL,
    time_of_incident TIME NOT NULL,
    FOREIGN KEY (hero_id) REFERENCES heroes(id),
    FOREIGN KEY (villain_id) REFERENCES villain(id)
);

INSERT INTO criminal_report (hero_id, villain_id, description, date_of_incident, time_of_incident) VALUES
(1, 1, 'Theft at Central Bank', '2024-03-20', '10:30:00'),
(2, 2, 'Hostage situation at City Hall', '2024-03-18', '14:45:00'),
(3, 3, 'Robbery at Jewelry Store', '2024-03-16', '12:00:00');