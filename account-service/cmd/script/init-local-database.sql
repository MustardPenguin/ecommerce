-- For local development only

DROP SCHEMA IF EXISTS account CASCADE;

CREATE SCHEMA account;

CREATE TABLE account.accounts (
    account_id INT PRIMARY KEY,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL
);