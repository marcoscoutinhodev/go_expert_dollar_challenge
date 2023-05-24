package adapter

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

func DollarRateAdapter() ([]byte, error) {
	channel := make(chan []byte, 1)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)

	defer cancel()

	go func() {
		req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)

		if err != nil {
			channel <- nil
			return
		}

		res, err := http.DefaultClient.Do(req)

		if err != nil || res.StatusCode != 200 {
			channel <- nil
			return
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)

		if err != nil {
			channel <- nil
			return
		}

		channel <- body
	}()

	select {
	case <-ctx.Done():
		fmt.Println("error on DollarRateAdapter: context timeout reached")
		return nil, errors.New("context timeout reached")
	case res := <-channel:
		if res == nil {
			return nil, errors.New("unexpected error, please try again")
		}

		return res, nil
	}
}
