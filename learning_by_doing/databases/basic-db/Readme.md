# Database Basics with GO
## Objective
This section covers
- how go handles databases. 
- What is the native way I can use to connect to Database
- What are the key takeaway

## How Go Handles Database 
- We know that database generally falls in two types SQL and NOSQL
### SQL
- In order to interact with any objects go uses Library, we need to import library.
- Similarly Go have Build-In Library called "database/sql" to interact with SQL DB. So we will import this Library at the beginning then we can interact with DB via GO.
### No SQL
- There is no Build-In Library for NOSQL DB when I'm writing this. However there are multiple third party library we can use to connect and interact with NOSQL library. 
- One Such example is "go.mongodb.org/mongo-driver" for MongoDB


## What is the native way I can use to connect to Database
- As discussed in the above section we can import "database/sql" at the beginning and then we can start interacting with database. Refer the main.go for example 
- However "database/sql" is more generic library which is created to interact with all SQL DB with Common interfaces. But there are many flavours in SQL db like postgres, Mysql, sqlite etc.
- So in-order to interact with specific flavour of DB we will import additional Drivers along with that. Refer the main.go file where we have imported '_ "github.com/go-sql-driver/mysql"' driver which works with database/sql native drivers and provide all the functionality. 


## What are the key takeaway
- Along with the native library we use to connect to DB we have to import corresponding DB Drivers to make full use of the library. (We can get the drivers of corresponding db with a simple google search)
- To Connect to DB we need to pass **DSN(DataSourceName)** input. The DSN contains all the information about your database such as UserName, Password, Server IP address, Port, DatabaseName. Along with that it also have some basic Config settings
- Example DSN
`dsn := "root:password@tcp(127.0.0.1:3306)/mysql?parseTime=true"`
    - Here parseTime=true >> Ensures date/time fields are returned as time.Time instead of string
- Example of Connecting to DB
    ```
    dsn := "root:password@tcp(127.0.0.1:3306)/mysql?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err) // Stop the program immediately with give the error message
	}
    ```

## How to Run this Program
- First initiate go module with go init `go mod init example.com/go-basic`
- Then download and add all dependencies using tidy `go mod tidy ` 
- Then you can directly run the program using `go run main.go`