package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePool() *pgxpool.Pool {
	log.SetPrefix("CreateDatabase: ")
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	conn, err := dbpool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error acquiring connection: ", err)
	}
	conn.Release()
	return dbpool
}
