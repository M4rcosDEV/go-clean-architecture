package app

import (
	"log"

	"github.com/M4rcosDEV/learningGO/internal/config"
	"github.com/M4rcosDEV/learningGO/internal/router"
)

func Initialize() {
	db, err := config.ConnectDatabase()

	if err != nil {
		log.Fatal(err)
	}	
	
	router := router.SetupRouter(db)
	
	//Started server
	err = router.Run(":8080")

	if err != nil {
		log.Fatal(err)
	}	
	
	log.Print("Application started")
}


