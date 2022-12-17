package database

import (
	"database/sql"

	"github.com/iVitaliya/Sancrity-Rebuild/lib/database/models"
	"github.com/iVitaliya/colors-go"
	"github.com/iVitaliya/logger-go"
)

func logError(errored_on string, message string) {
	logger.Fatal(colors.BrightRed("Database Error"), colors.BrightBlack(":"), colors.BrightRed(errored_on), colors.BrightBlack(">>"), message)
}

func Database() {
	db, err := sql.Open("sqlite3", "../data.sqlite")
	if err != nil {
		logError("Opening Path", err.Error())
	}
	defer db.Close()

	models.AFKDatabase(db)
}
