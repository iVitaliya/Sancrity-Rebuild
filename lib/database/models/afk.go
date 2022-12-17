package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type IAFKDatabase struct {
	CreateTable func()
	Get         func(key string, default_value any) *IGetData
	Set         func(key string, value bool)
}

type IGetData struct {
	guild_id string
	id       string
	message  string
	afk      int
}

func AFKDatabase(db *sql.DB) *IAFKDatabase {
	var (
		err  error
		rows *sql.Rows
	)

	return &IAFKDatabase{
		CreateTable: func() {
			CreateTable(db, "afk", `(
				guild_id TEXT,
				id TEXT,
				message TEXT,
				afk INTEGER
			)`)
		},
		Get: func(key string, default_value any) *IGetData {
			var data *IGetData
			rows, err = db.Query("SELECT value FROM afk WHERE key = ?", key)
			if err != nil {
				LogDatabaseError("Retrieving Data", err.Error())
			}
			defer rows.Close()

			for rows.Next() {
				var (
					guild_id string
					id       string
					message  string
					afk      int
				)

				err = rows.Scan(&guild_id, &id, &message, &afk)
				if err != nil {
					LogDatabaseError("Scanning For Data", err.Error())
				}

				data = &IGetData{
					guild_id: guild_id,
					id:       id,
					message:  message,
					afk:      afk,
				}
			}

			err = rows.Err()
			if err != nil {
				LogDatabaseError("After Retrieve Iteration", err.Error())
			}

			return data
		},
		// TODO: SET, CLEAR, DELETE
	}
}
