package v1

import (
	"fmt"
	"image/jpeg"
	"net/http"
	"strconv"

	"github.com/MorselShogiew/Users-service-rest/errs"
	"github.com/MorselShogiew/Users-service-rest/middleware"
	// _ "image/jpeg"
)

func (h *Handlers) GetResizePhoto(w http.ResponseWriter, r *http.Request) {
	reqID := middleware.GetReqID(r)

	heightStr := r.URL.Query().Get("height")
	widthStr := r.URL.Query().Get("width")
	url := r.URL.Query().Get("url")

	if heightStr == "" || widthStr == "" || url == "" {
		err := errs.New(nil, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, reqID)
		return
	}
	height, err := strconv.ParseUint(heightStr, 10, 32)
	if err != nil {
		err := errs.New(nil, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, reqID)
		return
	}

	width, err := strconv.ParseUint(widthStr, 10, 32)
	if err != nil {
		err := errs.New(nil, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, reqID)
		return
	}

	res, err := h.u.ResizePhoto(reqID, height, width, url)
	fmt.Println(err)
	h.CheckErrWriteResp(err, w, reqID)
	// Encode uses a Writer, use a Buffer if you need the raw []byte

	if err = jpeg.Encode(w, res, nil); err != nil {
		err := errs.New(err, errs.ErrBadRequest, false, 500)
		h.CheckErrWriteResp(err, w, reqID)
		return
	}

}
