CREATE DATABASE IF NOT EXISTS devbook;

Use devbook;


DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(


    id int auto_increment primary key, 
    cpf varchar(11) not null unique,
    nome varchar(60),
    nick varchar(60) not null unique,
    email varchar(60) not null unique, 
    senha varchar(20)  not null, 
    criadoem timestamp default current_timestamp()
) ENGINE=INNODB; 