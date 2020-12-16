CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    password VARCHAR(255) DEFAULT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    mail VARCHAR(255) NOT NULL,
    sex VARCHAR(6) NOT NULL,
    birthday DATE NOT NULL,
    height DECIMAL(3, 2) DEFAULT NULL,
    weight DECIMAL(3, 2) DEFAULT NULL,
    status TINYINT(1) DEFAULT 0,
    created_at DATE NOT NULL,
    updated_at DATE NOT NULL
);
