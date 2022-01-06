package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

//Retorna resposta JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

//Retorna um erro em formato JSON
func Erro(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Error string
	}{
		Error: err.Error(),
	})
}
