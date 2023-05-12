CREATE TABLE Roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(55) NOT NULL
);

CREATE TABLE Users
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    roleId  INT NOT NULL REFERENCES Roles(id)
);


CREATE TABLE Products
(
    id   SERIAL  PRIMARY KEY,
    title VARCHAR(255),
    price REAL NOT NULL,
    description VARCHAR(255),
    category VARCHAR(55) NOT NULL,
    image VARCHAR(55),
    done BOOLEAN NOT NULL DEFAULT FALSE

);

CREATE TABLE User_Product
(
    user_id INT  NOT NULL REFERENCES Users(id) ON DELETE CASCADE,
    product_id INT NOT NULL REFERENCES Products(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, product_id)

);
