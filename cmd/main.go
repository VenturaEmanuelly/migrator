package main

import (
	"autoMigrations/migrator"
	"context"
	"database/sql"
	"flag"

	_ "github.com/lib/pq"
)

func main() {

	path:=flag.String("p","path","folder path")


	ctx := context.TODO()

	db, err := sql.Open("postgres", "host=localhost port=5432 user= store password=golang dbname=store sslmode=disable")
	if err != nil {
		panic(err)
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	newMigrator := migrator.NewMigrator(tx)
	err = newMigrator.Migrate(ctx, *path)
	if err != nil{
		panic(err)
	}
}
