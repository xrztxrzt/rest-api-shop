CREATE TABLE Rating_list
(   
    id SERIAL PRIMARY KEY,
    count int,
    rate real
);

CREATE TABLE Products
(
    id SERIAL  NOT NULL UNIQUE ,
    title VARCHAR(255),
    price MONEY NOT NULL,
    description VARCHAR(255),
    category VARCHAR(55) NOT NULL,
    image VARCHAR(55),
    rating_id int references Rating_list(id)
);

