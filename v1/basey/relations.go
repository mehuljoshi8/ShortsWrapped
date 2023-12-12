/* Author: Mehul Joshi
 * File: relations.go
 * Description: relations.go is a file to store all the relations
 * of basey. Currently there is a User, Link, and Recipe struct in this file.
 */
package basey

/*
CREATE TABLE USERS (

	id SERIAL PRIMARY KEY,
	number varchar(12) UNIQUE NON NULL

);
*/
type User struct {
	id     uint64
	number string
}

/*
CREATE TABLE LINKS (

	id SERIAL PRIMARY KEY,
	hyperlink varchar NON NULL,
	user_id INTEGER REFERENCES users (id)

);
*/
type Link struct {
	Id        uint64
	UserId    uint64
	Identifer string
}

/*
CREATE TABLE RECIPES (

	id SERIAL PRIMARY KEY,
	identifer character(11) NOT NULL UNIQUE,
	title	character varying,
	body	character varying

);
*/
type Recipe struct {
	Id         uint64
	Identifier string
	Title      string
	Body       string
}
