CREATE DATABASE IF NOT EXISTS dev;
USE dev;

DROP TABLE IF EXISTS `users`;

CREATE TABLE users(
    id INT NOT NULL AUTO_INCREMENT primary key,
    name VARCHAR(50) NOT NULL,
    nickname VARCHAR(25) NOT NULL unique,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(20) NOT NULL,
    createAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)ENGINE=InnoDB;
