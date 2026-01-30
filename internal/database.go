package internal

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./my.db")
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to SQLite")

	// Test connection
	var version string
	if err := db.QueryRow("SELECT sqlite_version()").Scan(&version); err != nil {
		return nil, err
	}

	fmt.Println("SQLite version:", version)

	// Create tables
	if err := CreateTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

func CreateTables(db *sql.DB) error {
	roomTable := `
	CREATE TABLE IF NOT EXISTS rooms (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code TEXT NOT NULL UNIQUE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	questionTable := `
	CREATE TABLE IF NOT EXISTS questions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		room_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		votes INTEGER DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (room_id) REFERENCES rooms(id)
	);`

	confusionTable := `
	CREATE TABLE IF NOT EXISTS confusion_events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		room_id INTEGER NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (room_id) REFERENCES rooms(id)
	);`

	if _, err := db.Exec(roomTable); err != nil {
		return err
	}

	if _, err := db.Exec(questionTable); err != nil {
		return err
	}

	if _, err := db.Exec(confusionTable); err != nil {
		return err
	}

	return nil
}
