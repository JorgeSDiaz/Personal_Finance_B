package server

import (
	"log"

	"github.com/JorgeSDiaz/Personal_Finance_B/api"
)

func Run() {
	router := api.SetUpRoutes()

	log.Println(
		"Personal Finance (Backend) is Running on http://localhost:8080",
	)

	router.Run(":8080")
}
