/*
 * CD to the src/ directory and $go get -u github.com/go-sql-driver/mysql to install the lib
 *
 * To install mysql in CentOS7
 * Add repo
 * wget http://repo.mysql.com/mysql-community-release-el7-5.noarch.rpm
 * sudo rpm -ivh mysql-community-release-el7-5.noarch.rpm
 * yum update
 *
 * Now, install mysql server
 * sudo yum install mysql-server
 * sudo systemctl start mysqld
 *
 * MySQL will bind to localhost (127.0.0.1) by default.
 *
 * mysql -u root -p
 * mysql>
 *
 * Now, we can create a Database tagsDB and attach a user to it
 * create database tagsDB;
 * CREATE USER 'asethi'@'localhost' IDENTIFIED BY 'pinewood';
 * GRANT ALL PRIVILEGES ON tags.* TO 'asethi'@'localhost';
 * exit

mysql> select host, user, password from mysql.user;
+-----------------------+--------+-------------------------------------------+
| host                  | user   | password                                  |
+-----------------------+--------+-------------------------------------------+
| localhost             | root   |                                           |
| localhost.localdomain | root   |                                           |
| 127.0.0.1             | root   |                                           |
| ::1                   | root   |                                           |
| localhost             |        |                                           |
| localhost.localdomain |        |                                           |
| localhost             | asethi | *539E51A7AFDBACD366C5CAD3965EBE9D8E275758 |
+-----------------------+--------+-------------------------------------------+

SHOW GRANTS FOR 'asethi'@'localhost';

 *
 * Now, create a table tags in the tagsDB
 * mysql -u asethi -p 
 * use testdb;
 * DROP TABLE IF EXISTS tags;
 * create table tags (customer_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, username TEXT, password TEXT);
 * exit
 */
package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func dbInit() {
	fmt.Println("Go MySql iniit")
	// sql.Open gives us an handle to the DB:
	//db, err := sql.Open("mysql", "asethi:pinewood@tcp(127.0.0.1:3306)/tagsDB")
	db, err := sql.Open("mysql", "asethi:pinewood@/tagsDB")
	//db, err := sql.Open("mysql", "root:@/tagsDB")
	if (err != nil) {
		panic(err.Error())
	}
	defer db.Close()
	//tx, err := db.Begin()
	if (err != nil) {
		panic(err.Error())
	}
	// Lets insert a username/password into the tagsDB DB
	fmt.Println("Inserting into tagsDB DB")
	insert, err := db.Query(`INSERT INTO tags(username, password) VALUES('asethi', 'pinewood')`)
	if (err != nil) {
		panic(err.Error())
	}

	fmt.Println("Inserted: ", insert)
	results, err := db.Query(`SELECT username, password FROM tags`)
	if (err != nil) {
		panic(err.Error())
	}
	fmt.Println("Reading from tags table")
	for results.Next() {
		var creds Credentials
		err = results.Scan(&creds.Username, &creds.Password)
		if (err != nil) {
			panic(err.Error())
		}
		fmt.Println("DB tagsDB:", creds.Username, creds.Password)
	}
	fmt.Println("Read all records")
}
