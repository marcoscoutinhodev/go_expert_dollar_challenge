package adapter

import (
	"context"
	"io"
	"net/http"
	"time"
)

func DollarRateAdapter() []byte {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)

	if err != nil {
		return nil
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil || res.StatusCode != 200 {
		return nil
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil
	}

	return body
}
