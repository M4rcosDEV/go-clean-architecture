package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {

	config, err := Init()

	if err != nil { 
		return nil, err
	}
	
	datasource := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", 
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
		config.DBMode,
		config.DBTimezone,
	)

	db, err := gorm.Open(postgres.Open(datasource), &gorm.Config{})	

	if err != nil {
		return nil, err
	}

	if err := Migrate(db); err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected!")
	return  db, nil
}