CREATE ROLE login LOGIN PASSWORD 'password123';

CREATE DATABASE books_shop WITH OWNER = login;

CREATE TABLE books (
  isbn    char(14) NOT NULL,
  title   varchar(255) NOT NULL,
  author  varchar(255) NOT NULL,
  price   decimal(5,2) NOT NULL
);

INSERT INTO books (isbn, title, author, price) VALUES
('001-0000001', 'Iron heel', 'Lack London', 15.99),
('001-0000002', 'The Time Machine', 'H. G. Wells', 10.99),
('001-0000003', 'Memories od terrorist', 'Boris Savinkov', 15.50),
('001-0000004', 'Fahrenheit 451', 'Ray Bradbury', 20.46),
('001-0000005', '1984', 'George Orwell', 18.50),
('001-0000006', 'I, Robot', 'Issac Asimov', 25.00);

ALTER TABLE books ADD PRIMARY KEY (isbn);
