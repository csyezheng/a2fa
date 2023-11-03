package backends

import "fmt"

type Mysql struct {
	// the database engine, always should be mysql
	engine string
	// the name of the database to use
	name     string
	user     string
	password string
	host     string
	port     int
	sslMode  string
}

func (db Mysql) Engine() string {
	return db.engine
}

func (db Mysql) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=%s",
		db.user, db.password, db.host, db.port, db.name, db.sslMode)
}
