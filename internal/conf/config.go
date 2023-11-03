package conf

import (
	backends2 "github.com/csyezheng/a2fa/internal/backends"
	"github.com/csyezheng/a2fa/internal/fs"
)

type Config struct {
	DatabaseBackend backends2.Backend
}

const (
	sqlite3FileName = "a2fa.db"
)

func DefaultConfig() *Config {
	return &Config{
		DatabaseBackend: backends2.NewSqlite3(fs.MakeFilenamePath(sqlite3FileName)),
	}
}
