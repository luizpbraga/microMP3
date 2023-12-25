-- auth_user é o usuario responsavel for salvar os dados de usuario
CREATE USER 'auth_user'@'localhost' IDENTIFIED BY 'Aauth123';

CREATE DATABASE IF NOT EXISTS auth;

GRANT ALL PRIVILEGES ON auth.* TO 'auth_user'@'localhost';

USE auth;

CREATE TABLE IF NOT EXISTS user (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL
);

INSERT INTO user (email, password) VALUES ('luiz@gmail.com', '123');
