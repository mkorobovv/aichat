package controller

import (
	"encoding/json"
	"errors"
	"net/http"
)

func (ctr *Controller) handleError(w http.ResponseWriter, err error) {
	respErr := new(responseError)

	switch {
	case errors.As(err, respErr):
		w.WriteHeader(respErr.status)

		respJsonBytes, marshalErr := json.Marshal(respErr)
		if marshalErr != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		_, writeErr := w.Write(respJsonBytes)
		if writeErr != nil {
			ctr.logger.Error(writeErr.Error())
		}
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	return
}
