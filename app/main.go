package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	pgxURL := "postgresql://postgres-dev:s3cr3tp4ssw0rd@db:5432/dev"
	conn, err := pgx.Connect(context.Background(), pgxURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var v0 int64
	var v1 int64
	err = conn.QueryRow(context.Background(), "select 1 as v0, 2 as v1;").Scan(&v0, &v1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("hello there!!!!", v0, v1)
}
