package database

import (
	"fmt"
	"github.com/KeishiIrisa/backend-go-template/internal/models"
	"gorm.io/gorm"
)

func UserSeed(db *gorm.DB) error {
	var count int64

	db.Model(&models.User{}).Count(&count)
	if count > 0 {
		return nil
	}
	users := [2]models.User{
		{
			FirstName:    "Taro",
			LastName:     "Yamada",
			Email:        "test1@sample.com",
			PasswordHash: "<dummy>",
		},
		{
			FirstName:    "Hanako",
			LastName:     "Suzuki",
			Email:        "test2@sample.com",
			PasswordHash: "<dummy>",
		},
	}

	if err := db.Create(&users).Error; err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}
	return nil
}
