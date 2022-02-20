package migrate

import (
	"github.com/jinzhu/gorm"
	"production_service/models"
	"production_service/utils"
)

var (
	teams = []models.Team{
		{ID: utils.GenShortUUID(), Name: "Sales"},
		{ID: utils.GenShortUUID(), Name: "Operations"},
		{ID: utils.GenShortUUID(), Name: "Integration"},
	}

	hubs = []models.Hub{
		{ID: utils.GenShortUUID(), Name: "Tokyo", GeoLocation: "Tundra"},
		{ID: utils.GenShortUUID(), Name: "Oulu", GeoLocation: "Accord"},
		{ID: utils.GenShortUUID(), Name: "Oslo", GeoLocation: "Sentra"},
		{ID: utils.GenShortUUID(), Name: "London", GeoLocation: "F-150"},
	}

	users = []models.User{
		{ID: utils.GenShortUUID(), Name: "Joonas", Role: "Admin"},
		{ID: utils.GenShortUUID(), Name: "Mark", Role: "Manager"},
		{ID: utils.GenShortUUID(), Name: "Catherine", Role: "Manager"},
		{ID: utils.GenShortUUID(), Name: "Sarah", Role: "Manager"},
	}
)

func MigrateDB() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=productionService sslmode=disable password=postgres")
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Team{})
    for index := range teams {
        db.Create(&teams[index])
    }
 
	db.AutoMigrate(&models.User{})
    for index := range users {
        db.Create(&users[index])
    }

	db.AutoMigrate(&models.Hub{})
    for index := range hubs {
        db.Create(&hubs[index])
    }
}
