package config

import (
	"log"
	"vse-bank/models" // Import your models package

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=2468 dbname=vse-bank port=5434 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	} else {
		log.Println("✅ Database connection successful.")
	}

	// ✅ Auto-migrate your models here
	err = database.AutoMigrate(&models.User{}, &models.Bank{}, &models.Loan{})
	if err != nil {
		log.Fatal("❌ Failed to auto-migrate:", err)
	} else {
		log.Println("✅ Auto-migration successful.")
	}

	DB = database
}
