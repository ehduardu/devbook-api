package main

import (
	"api/src/router"
	"api/src/utils/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadVariables()
	r := router.Generate()

	fmt.Printf("Servidor rodando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", r))

}
