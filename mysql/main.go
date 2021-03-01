package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	ID   int    `json:id`
	Name string `json:name`
}

var db *sql.DB

func main() {
	fmt.Println("Go MySQL Tutorial")

	// opens the db conn
	db = openConnDB()

	//insert(db)
	// selected(db)
	singleRow(db)

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
}

func openConnDB() *sql.DB {
	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:example@tcp(127.0.0.1:3306)/testDb")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	return db
}

func insert(db *sql.DB) {
	// peform a db.Query
	insert, err := db.Query("INSERT INTO test VALUES ( 27, 'Ricardo' )")
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// be careful deferring queries if you are using transactions
	defer insert.Close()
}

func selected(db *sql.DB) {
	results, err := db.Query("SELECT id, name FROM test")

	for results.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object

		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		fmt.Println(tag.ID, tag.Name)
	}
}

func singleRow(db *sql.DB) {
	var tag Tag

	err := db.QueryRow("SELECT id, name FROM test where id = ?", 27).Scan(&tag.ID, &tag.Name)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(tag.ID, tag.Name)
}
