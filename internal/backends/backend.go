package backends

// Backend is the database backend interface that must be implemented by all database backend.
type Backend interface {
	Engine() string
	DSN() string
}
