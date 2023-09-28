package inits

import "github.com/Fishmansky/noteflow/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
