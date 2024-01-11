package domain

import (
	"Kinopoisk-Parser/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"strconv"
)

func InitPgConnectionByParams(params config.DatabaseConnectionParams) *sql.DB {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", params.User,
		params.Password, params.Host, strconv.Itoa(int(params.Port)), params.Database)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	return db
}

func InitNeo4jConnectionByParams(params config.DatabaseConnectionParams) neo4j.DriverWithContext {
	connString := fmt.Sprintf("neo4j://%s:%s", params.Host, strconv.Itoa(int(params.Port)))

	driver, err := neo4j.NewDriverWithContext(connString, neo4j.BasicAuth(params.User, params.Password, ""))
	if err != nil {
		panic(err)
	}

	return driver
}
