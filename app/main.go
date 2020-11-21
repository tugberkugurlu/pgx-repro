package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {
	pgxURL := "postgresql://postgres-dev:s3cr3tp4ssw0rd@db:5432/dev"
	dbConfig, err := pgxpool.ParseConfig(pgxURL)
	if err != nil {
		log.Fatalf("Unable to parse config: %v\n", err)
	}

	dbConfig.ConnConfig.Logger = &pgxLogger{}
	dbConfig.ConnConfig.LogLevel = pgx.LogLevelTrace

	ctx1 := context.WithValue(context.Background(), "ctx_name", "1")
	pool, err := pgxpool.ConnectConfig(ctx1, dbConfig)
	if err != nil {
		log.Fatalf("Unable to create pool: %v\n", err)
	}
	defer pool.Close()

	ctx2 := context.WithValue(context.Background(), "ctx_name", "2")
	conn, err := pool.Acquire(ctx2)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Release()

	var v0 int64
	var v1 int64
	ctx3 := context.WithValue(context.Background(), "ctx_name", "3")
	err = conn.QueryRow(ctx3, "select 1 as v0, 2 as v1;").Scan(&v0, &v1)
	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
	}

	fmt.Println("hello there!!!!", v0, v1)
}

type pgxLogger struct {
}

func (l *pgxLogger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	log.Printf("%s: [%v] %s (%v)\n", level.String(), ctx.Value("ctx_name"), msg, data)
}

func initPgxLogger(logger *logrus.Entry) pgx.Logger {
	return &pgxLogger{
	}
}