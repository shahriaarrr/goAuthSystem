package initializers

import "goAuthSystem/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})

}
