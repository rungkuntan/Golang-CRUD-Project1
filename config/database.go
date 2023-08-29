package config

import (
	"database/sql"
	"fmt"
	"go_project/errorhelper"

	_ "github.com/lib/pq" // Postgres golang driver
	"github.com/rs/zerolog/log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = ""
	password = ""
	dbName   = ""
)

func DatabaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", sqlInfo)
	errorhelper.PanicIfErr(err)

	err = db.Ping()
	errorhelper.PanicIfErr(err)

	log.Info().Msg("Connected to database!!")

	return db
}