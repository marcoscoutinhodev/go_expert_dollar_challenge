package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/marcoscoutinhodev/go_expert_dollar_challenge/external/repository/database"
	"github.com/marcoscoutinhodev/go_expert_dollar_challenge/internal/entity"
)

func AddDollarRateRepository(dollarRateEntity *entity.DollarRateEntity) {
	channel := make(chan error, 1)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*10)

	defer cancel()

	go func() {
		db := database.GetSqliteConnection()

		defer db.Close()

		stmt, err := db.Prepare(`
			INSERT OR IGNORE INTO dollar_rates (
				code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date
			)
			SELECT ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
			WHERE NOT EXISTS (SELECT * FROM dollar_rates WHERE code = ?);
		`)

		if err != nil {
			panic(err)
		}
		_, err = stmt.ExecContext(
			ctx,
			&dollarRateEntity.Usdbrl.Code,
			&dollarRateEntity.Usdbrl.Codein,
			&dollarRateEntity.Usdbrl.Name,
			&dollarRateEntity.Usdbrl.High,
			&dollarRateEntity.Usdbrl.Low,
			&dollarRateEntity.Usdbrl.VarBid,
			&dollarRateEntity.Usdbrl.PctChange,
			&dollarRateEntity.Usdbrl.Bid,
			&dollarRateEntity.Usdbrl.Ask,
			&dollarRateEntity.Usdbrl.Timestamp,
			&dollarRateEntity.Usdbrl.CreateDate,
			&dollarRateEntity.Usdbrl.Code,
		)

		if err != nil {
			channel <- err
			return
		}

		channel <- nil
	}()

	select {
	case <-ctx.Done():
		fmt.Println("error on AddDollarRateRepository: context timeout reached")
		panic(errors.New("context timeout reached"))
	case err := <-channel:
		if err != nil {
			fmt.Printf("error on AddDollarRateRepository: %v\n", err)
			panic(err)
		}

		return
	}
}
