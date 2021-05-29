package interfaces

import "net/http"

type Presenters interface {
	JSON(w http.ResponseWriter, r *http.Request, v interface{})
	Error(w http.ResponseWriter, r *http.Request, err error)
}
