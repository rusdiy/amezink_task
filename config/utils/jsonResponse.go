package utils

/*
	Utility functions for sending back response in a formmatted way.
*/

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Message string      `json:"msg"`
	Code    int         `json:"code"`
	Records interface{} `json:"records"`
}

func (j JSONResponse) SendResponse(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&j)
}

func SendError(w http.ResponseWriter, err error, statusCode int) {
	jsonResponse := JSONResponse{
		Message: err.Error(),
		Code:    statusCode,
		Records: nil,
	}
	jsonResponse.SendResponse(w, statusCode)
}
