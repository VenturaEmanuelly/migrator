package migrator

import (
	"context"
	"database/sql"
	"io/ioutil"
	"os"
)

type Migrator struct {
	transaction *sql.Tx
}

func (m Migrator) Migrate(ctx context.Context, path string) error {
	directory, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, d := range directory {
		data, err := ioutil.ReadFile(path + d.Name())
		if err != nil {
			return err
		}

		_, err = m.transaction.ExecContext(ctx, string(data))
		if err != nil {
			return err
		}
	}

	return nil
}

func NewMigrator(transaction *sql.Tx) Migrator{
	return Migrator{transaction: transaction}
}
