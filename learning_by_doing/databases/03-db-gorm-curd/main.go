package main

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *zap.SugaredLogger
	err    error
)

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"unique"`
}

func ConnectDB() {
	dsn := "root:password@tcp(127.0.0.1:3306)/mysql?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Error Connecting to db %v", err)
	}
	logger.Info("Database Connected Successfully")
}

func main() {

	// Logging
	loggers, _ := zap.NewProduction()
	defer loggers.Sync()     // Flushes buffer if any
	logger = loggers.Sugar() // This is new logger where error can be given in json format. Refer Readme for more info

	// Lets Connect to DB
	ConnectDB()

	// We are Pinging DB
	sqlDB, err := db.DB()

	pingErr := sqlDB.Ping()
	if pingErr != nil {
		logger.Fatalf("Error Connecting to db %v", err)
	}
	logger.Info("Ping successful to DB")

	db.AutoMigrate(&User{}) // Creates empty table with User struct as schema

	// Create Operation

}
