package databases

import "github.com/arueljust/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Product{})
}
