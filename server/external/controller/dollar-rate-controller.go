package controller

import (
	"encoding/json"
	"net/http"

	use_case "github.com/marcoscoutinhodev/go_expert_dollar_challenge/internal/use-case"
)

type ErrorHelper struct {
	Error string `json:"error"`
}

func DollarRateController(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if recover() != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&ErrorHelper{Error: "internal server error"})
			return
		}
	}()

	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&ErrorHelper{Error: "404 NOT FOUNDD"})
		return
	}

	dollarRateResponseEntity, err := use_case.DollarRateResolverUseCase()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&ErrorHelper{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dollarRateResponseEntity)
}
