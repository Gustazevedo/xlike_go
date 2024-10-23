package database

import (
    "fmt"
    "github.com/jackc/pgx/v4/pgxpool"
    "os"
    "context"
)

func NewPostgresConnection() (*pgxpool.Pool, error) {
    dbUrl := os.Getenv("postgres://postgres:opq198lt7m@localhost:5432/xlike_go") // Ensure this is set in your environment variables
    config, err := pgxpool.ParseConfig(dbUrl)
    if err != nil {
        return nil, fmt.Errorf("error configuring the database: %v", err)
    }

    pool, err := pgxpool.ConnectConfig(context.Background(), config)
    if err != nil {
        return nil, fmt.Errorf("error connecting to the database: %v", err)
    }

    return pool, nil
}
