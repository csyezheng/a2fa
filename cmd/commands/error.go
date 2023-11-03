package commands

import (
	"errors"
	"fmt"
)

var errHelp = errors.New("help requested")

// CommandError is returned when a command fails because of a user error (unknown command, invalid flag etc.).
// All other errors comes from the execution of the command.
type CommandError struct {
	Err error
}

// Error implements error.
func (e *CommandError) Error() string {
	return fmt.Sprintf("command error: %v", e.Err)
}

// Is reports whether e is of type *CommandError.
func (*CommandError) Is(e error) bool {
	_, ok := e.(*CommandError)
	return ok
}

// IsCommandError  reports whether any error in err's tree matches CommandError.
func IsCommandError(err error) bool {
	return errors.Is(err, &CommandError{})
}
