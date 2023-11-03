package backends

import (
	"fmt"
)

type Sqlite3 struct {
	// the database engine, always should be sqlite3
	engine string
	// the name is the full path to the database file for sqlite3
	name string
}

func (db Sqlite3) Engine() string {
	return db.engine
}

func (db Sqlite3) DSN() string {
	return fmt.Sprintf("file:%s?_journal=WAL&_vacuum=incremental", db.name)
}

func NewSqlite3(filePath string) *Sqlite3 {
	return &Sqlite3{
		engine: "sqlite3",
		name:   filePath,
	}
}
