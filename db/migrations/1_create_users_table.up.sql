CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    password VARCHAR(255) DEFAULT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    mail VARCHAR(255) NOT NULL,
    sex INT(1) NOT NULL,
    birthday DATE NOT NULL,
    height DECIMAL(5, 2) DEFAULT NULL,
    weight DECIMAL(5, 2) DEFAULT NULL,
    status TINYINT(1) DEFAULT 0,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);
