package database

import (
	"api/src/models"
	"api/src/utils/config"
	"fmt"
	"log"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup() {
	var dsn = fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable",
		config.DATABASE.InstanceName,
		config.DATABASE.User,
		config.DATABASE.Name,
		config.DATABASE.Password,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "cloudsqlpostgres",
		DSN:        dsn,
	}))

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
