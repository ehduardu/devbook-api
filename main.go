package main

import (
	"api/src/router"
	"api/src/utils/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Servidor rodando")

	r := router.Generate()

	config.LoadVariables()

	fmt.Println(config.INTERNAL_SECRET)
	fmt.Println(config.FIREBASE_CREDENTIALS)

	log.Fatal(http.ListenAndServe(":5000", r))
}
