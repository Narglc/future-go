package main

// MYSQL 建表
/*
CREATE DATABASE gin_restful_api;

USE gin_restful_api;

CREATE TABLE user_info(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	name CHAR(32) NOT NULL
);
*/

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
