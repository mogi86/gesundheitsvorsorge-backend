CREATE TABLE temporary_registrations (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    token VARCHAR(255) NOT NULL,
    expire_date DATE NOT NULL,
    created_at DATE NOT NULL,
    updated_at DATE NOT NULL,
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);
