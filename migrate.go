package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	db, err := sql.Open("mysql", "root:passapp@tcp(127.0.0.1:3306)/gotest")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	}

	stmt, err := db.Prepare("CREATE TABLE person (id int NOT NULL AUTO_INCREMENT, first_name varchar(40), last_name varchar(40), PRIMARY KEY (id));")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	} else {
		fmt.Printf("Person Table successfully migrated....")
	}
}
