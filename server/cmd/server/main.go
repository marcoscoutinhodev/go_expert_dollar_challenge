package main

import (
	"net/http"

	"github.com/marcoscoutinhodev/go_expert_dollar_challenge/external/controller"
	"github.com/marcoscoutinhodev/go_expert_dollar_challenge/external/repository/database"
)

func init() {
	database.PreparaDatabase()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/dollar_value", controller.DollarRateController)

	if err := http.ListenAndServe(":4001", mux); err != nil {
		panic(err)
	}
}
