package errors

import (
	"fmt"
	"runtime"
)

func ThrowDocError(what, which string) error {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	fnName := "<could not find fnName>"
	if ok && details != nil {
		fnName = getFunctionName(details.Name())
	}
	return &DocError{
		what,
		which,
		fnName,
	}
}

type DocError struct {
	What  string // why the error occured
	Which string // which document got the error
	Where string // where in the code the error occured
}

func (e DocError) Error() string {
	return fmt.Sprintf("Error: %v . Occured in %v  for Doc : %v", e.What, e.Where, e.Which)
}
