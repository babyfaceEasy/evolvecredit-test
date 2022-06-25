package db

import (
	"database/sql"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func GetDBConnection() *bun.DB {
	// open database connection
	dsn := os.Getenv("DATABASE_URL")

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// create bun object
	bun := bun.NewDB(sqldb, pgdialect.New())

	return bun
}