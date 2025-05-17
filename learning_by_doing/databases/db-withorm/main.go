package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	err error
	DB  *gorm.DB // Global DB Instance. REFER README FOR MORE INFO ON WHY THIS IS DECLARED LIKE THIS
)

// Connect to DB

func ConnectDB() {
	dsn := "root:password@tcp(127.0.0.1:3306)/mysql?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB: \n", err)
	}
	log.Println("Database connected Successfully")
}

func main() {
	// Let connect to DB
	ConnectDB()

	// We are trying to using Ping method which is of gorm.io/driver/mysql, that needs DB() method so we first call DB method with our GormDB Pointer vairable which contains our db values and then we are using Ping method with the variable
	sqlDB, err := DB.DB()

	// Testing the connection
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Error Connecting to db.\n", err)
	}
	log.Println("Ping to DB successfully. App is Ready")

}
