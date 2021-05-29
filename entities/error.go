package entities

import (
	"fmt"
	"net/http"
	"runtime"
)

type StatusError struct {
	Code   int
	Err    error
	Caller string
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}

func newError(err error, code int) StatusError {
	pc, _, line, _ := runtime.Caller(2)
	details := runtime.FuncForPC(pc)

	return StatusError{
		Code:   code,
		Err:    err,
		Caller: fmt.Sprintf("%s#%d", details.Name(), line),
	}
}

func ErrorBadRequest(err error) StatusError {
	return newError(err, http.StatusBadRequest)
}

func ErrorNotFound(err error) StatusError {
	return newError(err, http.StatusNotFound)
}

func ErrorInternal(err error) StatusError {
	return newError(err, http.StatusInternalServerError)
}
