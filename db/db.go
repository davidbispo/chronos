package db

import (
	"log"

	"chronos-scheduler.com/api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var DB *gorm.DB

func InitDB() {
	DB, err := gorm.Open("sqlite3", "./appointments.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB.AutoMigrate(
		&models.Appointment{},
		&models.Attendee{})
}
