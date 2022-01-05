package database

import (
	"api/src/utils/config"
	"database/sql"
	"fmt"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*sql.DB, error) {

	var DSN = fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable",
		config.DATABASE.InstanceName,
		config.DATABASE.User,
		config.DATABASE.Name,
		config.DATABASE.Password,
	)

	connection, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "cloudsqlpostgres",
		DSN:        DSN,
	}))

	if err != nil {
		return nil, err
	}

	db, _ := connection.DB()

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, err

}
