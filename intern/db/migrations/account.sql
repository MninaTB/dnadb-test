CREATE TABLE IF NOT EXISTS account (
    id VARCHAR(255),
    mail VARCHAR(255),
    password VARCHAR(255),
    nickname VARCHAR(255),
    firstname VARCHAR(255),
    lastname VARCHAR(255),
    country INT,
    street VARCHAR(255),
    streetnumber INT,
    zipcode INT,
    city VARCHAR(255),
    breeder TINYINT,
    created DATE,
    updated DATE
);