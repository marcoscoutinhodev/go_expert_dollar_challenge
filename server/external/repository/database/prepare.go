package database

func PreparaDatabase() {
	db := GetSqliteConnection()

	defer db.Close()

	_, err := db.Exec(`
		DROP TABLE IF EXISTS dollar_rates;
		CREATE TABLE dollar_rates (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			code VARCHAR(50),
			codein VARCHAR(50),
			name VARCHAR(50),
			high VARCHAR(50),
			low VARCHAR(50),
			varBid VARCHAR(50),
			pctChange VARCHAR(50),
			bid VARCHAR(50),
			ask VARCHAR(50),
			timestamp VARCHAR(50),
			create_date DATETIME
		);
	`)

	if err != nil {
		panic(err)
	}
}
