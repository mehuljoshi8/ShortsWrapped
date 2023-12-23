/* Author: Mehul Joshi
 * File: relations.go
 * Description: relations.go is a file to store all the relations
 * of basey. Currently there is a User, Link, and Recipe struct in this file.
 */
package basey

/*
CREATE TABLE Document (

	id SERIAL PRIMARY KEY,
	identifer character(11) NOT NULL UNIQUE,
	title	character varying,
	body	character varying

);
*/
type Document struct {
	Id         uint64
	Identifier string
	Title      string
	Body       string
}
