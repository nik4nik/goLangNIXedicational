create database productdb;
use productdb;
create table posts (
	id int auto_increment primary key,
	user_id int not null,
	title varchar(80) not null,
	body varchar(300) not null
);
create table comments (
	id int auto_increment primary key,
	post_id int not null,
	name varchar(80) not null,
	email varchar(32) not null,
	body varchar(250) not null,
	FOREIGN KEY (post_id)  REFERENCES posts (id)
);
