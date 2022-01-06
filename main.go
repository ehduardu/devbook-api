package main

import (
	"api/src/database"
	"api/src/router"
	"api/src/utils/config"
	"fmt"
	"log"
	"net/http"
)

func initializeRouter() {
	r := router.Generate()

	fmt.Printf("Servidor rodando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}

func main() {
	config.LoadVariables()
	database.Setup()

	initializeRouter()
}
