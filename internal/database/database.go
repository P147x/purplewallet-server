package database

import (
	"log"
	"purplewallet/internal/config"
	"purplewallet/internal/models"
	"strconv"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

//GetDatabase is a singleton for getting the instance of the database
func GetDatabase() *gorm.DB {
	if db == nil {
		initDatabase()
	}
	return db
}

// InitDatabase is used for the first instanciation of the database, connecting to the server and creating the missing tables.
func initDatabase() {
	cstr := config.Config.Database.User + ":" + config.Config.Database.Password + "@tcp(" + config.Config.Database.URL + ":" + strconv.Itoa(config.Config.Database.Port) + ")/" + config.Config.Database.DBName + "?charset=utf8&parseTime=True"
	var err error
	db, err = gorm.Open("mysql", cstr)
	log.Println("Connection to database started")

	if err != nil {
		log.Fatalln("Error occured: " + err.Error())
		return
	}
	createTables()
	log.Println("Success")
	//TODO: Check if debug mode
	db.LogMode(true)
	return
}

// CloseDatabase is used for safe disconnection to the database server
func CloseDatabase() {
	db.Close()
	log.Print("Database closed")
}

func createTables() {
	db.AutoMigrate(&models.Categories{})
	db.AutoMigrate(&models.Users{})
	db.AutoMigrate(&models.Wallets{})
	db.AutoMigrate(&models.Purchase{})
	db.AutoMigrate(&models.Logs{})
}

func Log(user uint, msg string, level uint) {
	db.Create(&models.Logs{
		Message: msg,
		UserID:  user,
		Level:   level,
	})
}
