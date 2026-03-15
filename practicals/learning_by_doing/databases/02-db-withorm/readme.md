# Database Interaction with GOLANG Using ORM

## Objective
This section covers
- What is ORM, why we need it
- How to use ORM to Connect and interact with DB
- Example to connect with DB using GORM

## What is ORM Why we need it
- ORM stands for **Object Relational Mapping**
- *Why we need it ?* - By Definition, It is a technique that allows us to interact with SQL database using Go Struct and Objects instead of RAW SQL Queries. 
- i.e Lets take an example that we need to query user table in a db and we want only one row as result. 
### Example: Native SQL (Without ORM)
```go
rows, err :=db.Query("SELECT * FROM user LIMIT 1")
```
### Now: Using GORM (ORM Library in Go)
```go
var user User // Here User refers to the struct (Model in DB)

result := db.Limit(1).Find(&user)
// or
result := db.First(&user)
```
- We don't need to work on Direct SQL queries, the code is more human readable, and it avoids repeating of same SQL Code everywhere, and easy to manage. Hence we are using ORM instead of native code.

## How to use ORM to Connect and interact with DB
- First step we need to import the Go library for GORM. Go have its standard GORM library called "gorm.io/gorm". 
- We will also import the corresponding driver for based on the flavour of DB we use such as Mysql, postgres or etc
```go
import (
    "gorm.io/gorm"             // GORM core package
    "gorm.io/driver/mysql"     // MySQL driver (use a different driver for PostgreSQL, SQLite, etc.)
)
```
- We will Declare a variable called DB which will be a pointer to *gorm.DB. This holds our DB connection. More on this in later section
- Then we use Gorm syntax to connect to establish connection to the Database.

### gorm syntax to connect to db
```go
// SYNATX
gorm.Open(dialector, config)
//dialector : Tells gorm which database you are using (Mysql postrges etc)
//config: gives additional settings to customize Gorm behavior

// EXAMPLE
dsn := "root:password@tcp(127.0.0.1:3306)/mysql?charset=utf8&parseTime=True&loc=Local"
DB, err := gorm.Open(<DRIVER TYPE>.Open(dsn),&gorm.Config{})
```
- *gorm.Open - This will open the Gorm Connection with specified settings*
- *Driver type - Defines which type of driver we are using such as mysql, postgres etc*
- *\<DRIVER TYPE\>.Open(dsn) - This will open connection to the DB with the DSN settings*
- gorm.Config{} - Pointer to the gorm Basic settings struct to change the gorm default behavior based on the need

### What is in *&gorm.Config{}*, why we use them
- This is pointer to the Gorm Default config struct. This gives you control over how gorm work globally. You can change some of the gorm behavior by changing the values from the struct such as enable logging, changing naming strategies etc
- You can find more details on the settings here: https://gorm.io/docs/gorm_config.html
```go
// EXAMPLE
&gorm.Config{
    Logger: logger.Default.LogMode(logger.Info), // show SQL logs
    DisableForeignKeyConstraintWhenMigrating: true,
}
```
### Why we need var DB *\*gorm.DB* (Refer *main.go* file)
I had question why the variable DB is declared and why it is pointing to gorm.DB
- The DB variable which is a pointer to **gorm.DB struct**
- The gorm.DB struct contains the information like db connections, config and other internal states We use a pointer (*gorm.DB) because:
    - Gorm returns and excepts pointers from gorm.Open() and in its method calls
    - Using a pointer allows us to share the single db connection acrros the whole application without copying the entire struct
    - When we call methods like DB.AutoMigrate()or DB.Create() those methods needs pointer as receivers so the value will be automatically dereferenced.
- *In Simple terms it holds your connection details. We used pointer because the changes to the connection should be updated on the same variable* 

## Example Connect with DB using GORM (Refer *main.go*)
To run and test the file follow below steps
- Run Docker container with below command 
    ```
    docker run --name go-mysql \
    -e MYSQL_ROOT_PASSWORD=password \
    -e MYSQL_DATABASE=mysql \
    -p 3306:3306 \
    -d mysql:8.0
    ```
- First initiate go module with go init `go mod init example.com/go-withorm`
- Then download and add all dependencies using tidy `go mod tidy ` 
- Then you can directly run the program using `go run main.go`