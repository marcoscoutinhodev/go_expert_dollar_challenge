package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

type ServerResponse struct {
	Dollar string `json:"dollar"`
}

type Convertion struct {
	Coin      string
	CoinValue string
	Dollar    float64
	Real      float64
}
type ErrorHelper struct {
	Error string `json:"error"`
}

func main() {
	http.HandleFunc("/home", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
		return
	}

	templates := []string{
		"template/header.html",
		"template/body.html",
		"template/footer.html",
	}

	t := template.New("body.html")
	t = template.Must(t.ParseFiles(templates...))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*300)

	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		"http://localhost:4001/dollar_value",
		nil,
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintln("internal server error, please try again..")))
		return
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintln("internal server error, please try again..")))
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintln("internal server error, please try again..")))
		return
	}

	if res.StatusCode != 200 {
		var errorHelper ErrorHelper
		json.Unmarshal(body, &errorHelper)
		fmt.Printf("error on request to http://localhost:4001/dollar_value | Status code: %d  | Message: %s\n", res.StatusCode, errorHelper.Error)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintln("internal server error, please try again..")))
		return
	}

	var serverResponse ServerResponse

	err = json.Unmarshal(body, &serverResponse)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintln("internal server error, please try again..")))
		return
	}

	dollar, err := strconv.ParseFloat(fmt.Sprintf(serverResponse.Dollar), 64)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintln("internal server error, please try again..")))
		return
	}

	file, err := os.OpenFile("history.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintln("internal server error, please try again..")))
		return
	}

	defer file.Close()

	file.WriteString(fmt.Sprintf("Dollar: %.2f\n", dollar))

	convertion := Convertion{
		Coin:      fmt.Sprintf("%.2f Brazilian Reais", 1.0),
		CoinValue: fmt.Sprintf("%.2f US Dollar", dollar),
		Dollar:    dollar,
		Real:      1,
	}

	err = t.Execute(w, convertion)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("internal server error: %v", err)))
		return
	}
}

func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(num*output)) / output
}
