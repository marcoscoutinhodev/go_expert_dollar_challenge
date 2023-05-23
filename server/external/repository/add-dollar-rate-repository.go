package repository

import (
	"context"
	"time"

	"github.com/marcoscoutinhodev/go_expert_dollar_challenge/external/repository/database"
	"github.com/marcoscoutinhodev/go_expert_dollar_challenge/internal/entity"
)

func AddDollarRateRepository(dollarRateEntity *entity.DollarRateEntity) {
	db := database.GetSqliteConnection()

	defer db.Close()

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*10)

	defer cancel()

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
		panic(err)
	}
}
