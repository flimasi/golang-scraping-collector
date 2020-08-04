package main

import (
	"database/sql"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
)

func main() {
getData()
store()
}

func store(){
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/scraping")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO scraping VALUES ( 2, 'tese', 'source', now() )")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}

func getData(){
	fmt.Println("Scanning Folder")

	var dirname = "dataBuffer/"

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
}