package config

import (
	"fmt"

	"github.com/M4rcosDEV/learningGO/internal/user"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error { 
	
	//All models are included here.
	err := db.AutoMigrate(
		&user.User{},
	)

	if err != nil {
		fmt.Print("Error running migrations.")
		return err
	}

	return nil
}
