package main

import (
	"encoding/json"
	"net/http"

	"github.com/eduardovasconcelos/gostudy/internal/entity"
)

func main() {
	http.HandleFunc("/order", Order)
	http.ListenAndServe(":8080", nil)
}

func Order(response http.ResponseWriter, request *http.Request) {
	order, err := entity.NewOrder("123", 10, 1)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(err.Error())
		return
	}
	order.CalculateFinalPrice()

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(order)
}
