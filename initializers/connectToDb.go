package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dbstring := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(dbstring), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	}

}
