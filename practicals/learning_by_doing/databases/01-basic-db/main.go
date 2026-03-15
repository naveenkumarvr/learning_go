package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Setting DataSourceName(DSN) Details. "?parseTime=true" Defines Go to parse the date time in time.Time format rather byte or string format.
	// By default it parse in string, by using the "go-sql-driver/mysql" in build capability we are telling go to format the time in time.Time format so that we can use it anywhere in the program
	dsn := "root:password@tcp(127.0.0.1:3306)/mysql?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err) // Stop the program immediately with give the error message
	}
	defer db.Close()                       // This make sure that all the db connection is closed once main exit no matter how it exit
	db.SetConnMaxLifetime(time.Minute * 3) // Help prevent using stale or timed-out-connection. Here we are saying 3min
	db.SetMaxOpenConns(10)                 // Prevent DB from opening too many connections. Now max is set to 10
	db.SetMaxIdleConns(10)                 // This will keep x amount of connection idle so that user can use same connection without opening a new one

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected Successfully")

	// Declaring the variable to store the result
	var Id int
	var Zone string
	// Here QueryRow is part of our "go-sql-driver/mysql" library.
	// what we are saying here is Select 1 row(which is SQL Command) and Scan will capture the output then we ask go to store result in Memory Location(&) of testValue variable. So that the actual variable get updated instead of a copy of variable
	err = db.QueryRow("SELECT * FROM time_zone_name LIMIT 1").Scan(&Zone, &Id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Test Query Result:", Id, Zone)
}
