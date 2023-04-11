CREATE TABLE Products
(
    id   SERIAL  NOT NULL PRIMARY KEY ,
    title VARCHAR(255),
    price REAL NOT NULL,
    description VARCHAR(255),
    category VARCHAR(55) NOT NULL,
    image VARCHAR(55)

);

