package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	testQueries *Queries
	testDB      *pgxpool.Pool
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = pgxpool.New(context.Background(), "postgresql://root:secret@localhost:5432/money_management?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	defer testDB.Close()

	testQueries = New(testDB)

	code := m.Run()

	os.Exit(code)

}
