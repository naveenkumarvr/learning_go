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
	gorm.Model
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"unique"`
}

func ConnectDB() {
	dsn := "root:password@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Error Connecting to db %v", err)
	}
	logger.Info("Database Connected Successfully")
}

func AddUser(user []User) {
	logger.Infof("Adding User %v", user)
	result := db.Create(&user)
	if result.Error != nil {
		logger.Errorf("Error Creating User %v", user)
	} else {
		logger.Infof("User %v created successfully", user)
	}
}

func GetUserByID(id int) *User {
	var user User
	logger.Infof("Getting User with the ID %v", id)
	result := db.First(&user, id)
	if result.Error != nil {
		logger.Errorf("Error Finding user with ID %v", id)
	} else {
		logger.Infof("User Found with id %v", id)
	}
	return &user
}

func DeleteUser(id int) *User {
	var user User
	result := db.Unscoped().Delete(&user, id)
	if result.Error != nil {
		logger.Errorf("Error Deleting user %v", id)
	} else {
		logger.Infof("Deleted User %v", id)
	}
	return &user
}

func GetAllUser() []User {
	logger.Infof("Showing all user")
	var user []User
	result := db.Find(&user)
	if result.Error != nil {
		logger.Errorf("Error Getting Data from Database")
	} else {
		logger.Infof("Successfully show the database")
	}
	return user
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
	// user := []User{
	// 	{
	// 		Name:  "Test02",
	// 		Email: "Test02@gmail.com",
	// 	},
	// }
	// AddUser(user)
	// result := GetAllUser()
	// result := GetUserByID(4)
	// logger.Infoln(result)
	DeleteUser(4)
}
