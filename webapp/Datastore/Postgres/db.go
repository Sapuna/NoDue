package Postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// db details
const (
	Postgres_host     = "db"
	Postgres_port     = 5432
	Postgres_user     = "postgres"
	Postgres_password = "postgres"
	Postgres_dbname   = "my_db"
)

// create pointer variqqble Db which points to sql driver
var Db *sql.DB

func init() {
	db_info := fmt.Sprintf("host=%s  port=%d user=%s password=%s dbname=%s sslmode=disable", Postgres_host, Postgres_port, Postgres_user, Postgres_password, Postgres_dbname)

	var err error
	Db, err = sql.Open("postgres", db_info)

	if err != nil {
		panic(err)
	} else {
		log.Println("Database successfully configured")
	}
}
