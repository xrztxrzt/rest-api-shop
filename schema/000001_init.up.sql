CREATE TABLE Products
(
    id SERIAL  NOT NULL UNIQUE ,
    title VARCHAR(255),
    price MONEY NOT NULL,
    description VARCHAR(255),
    category VARCHAR(55) NOT NULL,
    image VARCHAR(55),
    ratingId SERIAL,
    Foreign Key (ratingId) REFERENCES (Rating_list) ON DELETE CASCADE
);

CREATE TABLE Rating_list
(   
    id SERIAL NOT NULL UNIQUE,
    count int,
    rate float64,
);