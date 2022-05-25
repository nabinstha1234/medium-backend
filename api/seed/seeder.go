package seed

import (
	"github.com/jinzhu/gorm"
	"log"
	"medium-clone-backend/api/models"
)

var users = []models.User{
	models.User{
		Username: "nabin",
		Email:    "nabin@nabin.com",
		Password: "password",
	},
	models.User{
		Username: "shrestha",
		Email:    "shrestha@shrestha.com",
		Password: "password",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
