-- Arquivo contendo alguns comandos para execução rápida no workbench em abiente de dev


CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(20) not null unique,
    createdat timestamp default current_timestamp()
) ENGINE=INNODB;