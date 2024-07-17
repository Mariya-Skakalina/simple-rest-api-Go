package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateTable, downCreateTable)
}

func upCreateTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`CREATE TABLE IF NOT EXISTS books(
		id SERIAL PRIMARY KEY,
        author VARCHAR(100) NOT NULL,
        title VARCHAR(300) NOT NULL
	)`)
	if err != nil {
		return err
	}
	return nil
}

func downCreateTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec("DROP TABLE users")
	if err != nil {
		return err
	}
	return nil
}
