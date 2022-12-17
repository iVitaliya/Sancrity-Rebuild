package models

import (
	"database/sql"

	"github.com/iVitaliya/colors-go"
	"github.com/iVitaliya/logger-go"
	_ "github.com/mattn/go-sqlite3"
)

func LogDatabaseError(errored_on string, message string) {
	logger.Fatal(colors.BrightRed("Database Error"), colors.BrightBlack(":"), colors.BrightRed(errored_on), colors.BrightBlack(">>"), message)
}

func CreateTable(db *sql.DB, name string, args string) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS ` + name)
	if err != nil {
		LogDatabaseError("Creating A Table", err.Error())
	}
}
