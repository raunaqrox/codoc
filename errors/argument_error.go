package errors

import (
	"fmt"
	"runtime"
	"strings"
)

// TODO : Add to utils function
func getFunctionName(nameWithModule string) string {
	return strings.Split(nameWithModule, ".")[1]
}

func ThrowArgumentError(what string) error {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	fnName := "<could not find fnName>"
	if ok && details != nil {
		fnName = getFunctionName(details.Name())
	}
	return &ArgumentError{
		what,
		fnName,
	}
}

type ArgumentError struct {
	What  string
	Where string
}

func (e ArgumentError) Error() string {
	return fmt.Sprintf("Error: %v occured in %v function", e.What, e.Where)
}
