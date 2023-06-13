package initializers

import "github.com/uguremirmustafa/go-gin-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
