CREATE DATABASE coffee_break;
USE coffee_break;

CREATE TABLE sample
(
    id     INT(11) AUTO_INCREMENT NOT NULL,
    value1 VARCHAR(64)            NOT NULL,
    value2 VARCHAR(64)            NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO sample (value1, value2)
VALUES ('value1', 'value2');