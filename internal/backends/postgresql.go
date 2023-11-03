package backends

import "fmt"

type Postgresql struct {
	// the database engine, always should be postgresql
	engine string
	// the name of the database to use
	name     string
	user     string
	password string
	host     string
	port     int
	sslMode  string
}

func (db Postgresql) Engine() string {
	return db.engine
}

func (db Postgresql) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Shanghai",
		db.host, db.user, db.password, db.name, db.port, db.sslMode)
}
