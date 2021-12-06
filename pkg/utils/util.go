package utils

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

var log = logrus.New()

type response struct {
	OK bool `json:"ok"`
}

type SuccessResponse struct {
	response
	Message string `json:"message"`
}

func newSuccessResponse(message string) SuccessResponse {
	return SuccessResponse{response: response{OK: true}, Message: message}
}

type errorResponse struct {
	response
	Message string `json:"error"`
}

func newErrorResponse(err error) errorResponse {
	return errorResponse{response: response{OK: false}, Message: err.Error()}
}

func WriteError(w http.ResponseWriter, err error) (int, error) {
	body := newErrorResponse(err)
	b, e := json.Marshal(body)
	if e != nil {
		return w.Write([]byte{})
	}
	return w.Write(b)
}

func WriteSuccess(w http.ResponseWriter, message string) (int, error) {
	body := newSuccessResponse(message)
	b, e := json.Marshal(body)
	if e != nil {
		return w.Write([]byte{})
	}
	return w.Write(b)
}
