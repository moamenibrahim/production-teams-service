package migrate

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	"production_service/models"
	"production_service/utils"
)

var (
	teams = []models.Team{
		{ID: "2bjzdXAhLjwdMtmZfbykqY", Name: "Sales", TeamType: "A"},
		{ID: "38DaJtarUhR6TtZTkT6oMa", Name: "Operations", TeamType: "B"},
		{ID: "4j89fm2gKMD2eLuEPssPNP", Name: "Integration", TeamType: "C"},
	}

	hubs = []models.Hub{
		{ID: "4j89fm2gKMD2eLuEPssPNR", Name: "Headquarters", GeoLocation: "Tokyo"},
		{ID: "4j89fm2gKMD2eLuEPssPNW", Name: "Engineering", GeoLocation: "Oulu"},
		{ID: "4j89fm2gKMD2eLuEPssPNZ", Name: "Sales", GeoLocation: "Oslo"},
		{ID: "3j89fm2gKMD2eLuEPssPNP", Name: "Marketing", GeoLocation: "London"},
	}

	users = []models.User{
		{ID: "38FaJtarUhR6TtZTkT6oMa", Name: "Joonas", Role: "Admin"},
		{ID: "38DaJtarUhR6TtZTkT9oMa", Name: "Mark", Role: "Manager"},
		{ID: "38DaJtarUhU6TtZTkT6oMa", Name: "Catherine", Role: "Contractor"},
		{ID: "38DaJtadUhR6TtZTkT6oMa", Name: "Sarah", Role: "Product Owner"},
	}
)

func MigrateDB() {
	HOST := os.Getenv("HOST")
	if HOST == "" {
		HOST = "localhost"
	}
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s", utils.DB_USER, utils.DB_PASSWORD, utils.DB_NAME, HOST)
	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		panic("failed to connect database")
	}

	db.DropTableIfExists(&models.Team{})
	db.AutoMigrate(&models.Team{})
	for index := range teams {
		db.Create(&teams[index])
	}

	db.DropTableIfExists(&models.User{})
	db.AutoMigrate(&models.User{})
	for index := range users {
		db.Create(&users[index])
	}

	db.DropTableIfExists(&models.Hub{})
	db.AutoMigrate(&models.Hub{})
	for index := range hubs {
		db.Create(&hubs[index])
	}
}
