package presenters

import (
	"api/utils"
	"net/http"
)

func (p *presenters) JSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	err := utils.WriteJson(w, v)
	if err != nil {
		p.Error(w, r, err)
	}
}
