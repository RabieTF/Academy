DROP TABLE IF EXISTS Products;
DROP TABLE IF EXISTS Shops;
DROP TABLE IF EXISTS Categories;
DROP TABLE IF EXISTS Users;


CREATE TABLE Users (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE Categories (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(255) NOT NULL UNIQUE,
    PRIMARY KEY (`id`)
);

CREATE TABLE Shops (
  id         INT AUTO_INCREMENT NOT NULL,
  name      VARCHAR(255) NOT NULL UNIQUE,
  address     VARCHAR(255) NOT NULL UNIQUE,
  owned_by      INT,
  PRIMARY KEY (id),
  FOREIGN KEY (`owned_by`) REFERENCES Users(`id`)
);

CREATE TABLE Products (
    id INT AUTO_INCREMENT NOT NULL,
    shop_id INT,
    name VARCHAR(255) NOT NULL UNIQUE,
    description VARCHAR(255),
    categories VARCHAR(255),
    PRIMARY KEY (`id`),
    FOREIGN KEY (`shop_id`) REFERENCES Shops(`id`)
);




INSERT INTO Categories (name) VALUES ('Food'),('Electronics'),('Cleaning');