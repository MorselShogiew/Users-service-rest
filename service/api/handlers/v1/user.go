package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"MorselShogiew/Users-service-rest/errs"
)

func (h *Handlers) AddUser(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	mail := r.URL.Query().Get("mail")

	err := h.u.AddUser(name, mail)

	h.CheckErrWriteResp(err, w, "")

}
func (h *Handlers) DeleteUser(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		err := errs.New(nil, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, "")
		return
	}
	if id < 0 {
		err := errs.New(nil, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, "")
		return
	}

	err = h.u.DeleteUser(id)

	h.CheckErrWriteResp(err, w, "")

}
func (h *Handlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	res, err := h.u.GetUsers()

	h.CheckErrWriteResp(err, w, "")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		err := errs.New(nil, errs.ErrJSONEncode, false, 500)
		h.CheckErrWriteResp(err, w, "")
		return
	}

}
