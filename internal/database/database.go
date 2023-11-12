package database

import (
	"github.com/csyezheng/a2fa/internal/backends"
	"github.com/csyezheng/a2fa/internal/conf"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"log/slog"
	"sync"
)

// Database is a handle to the database layer
type Database struct {
	db      *gorm.DB
	dbLock  sync.Mutex
	backend backends.Backend
}

// Open creates a database connection. This should only be called once.
func (db *Database) Open() error {
	var rawDB *gorm.DB
	var err error
	switch db.backend.Engine() {
	case "sqlite3":
		dsn := db.backend.DSN()
		slog.Debug(dsn)
		rawDB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
		})
		sqlDB, _ := rawDB.DB()
		sqlDB.SetMaxOpenConns(1)
	case "mysql":
		dsn := db.backend.DSN()
		rawDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case "postgresql":
		dsn := db.backend.DSN()
		rawDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	default:
		log.Fatalf("not supported database engine: %s", db.backend.Engine())
	}
	if err != nil {
		log.Fatalf("failed to connect database:%s", err.Error())
	}
	db.db = rawDB
	return nil
}

// Close will close the database connection. Should be deferred right after Open.
func (db *Database) Close() error {
	sqlDB, err := db.db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	return sqlDB.Close()
}

// LoadDatabase initialize database instance, it does not connect to the database.
func LoadDatabase() (*Database, error) {
	config := conf.DefaultConfig()
	backend := config.DatabaseBackend
	return &Database{
		backend: backend,
	}, nil
}

func (db *Database) Engine() string {
	return db.backend.Engine()
}

func (db *Database) AutoMigrate(dst ...interface{}) error {
	var err error
	if db.Engine() == "mysql" {
		err = db.db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(dst...)
	} else {
		err = db.db.AutoMigrate(dst...)
	}
	return err
}
