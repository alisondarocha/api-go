package db

import (
	"api-go/configs"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	cfg := configs.GetDB()

	stringConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmodde=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.Database)

	connection, err := sql.Open("postgres", stringConnection)
	if err != nil {
		panic(err)
	}

	err = connection.Ping()

	return connection, err
}
