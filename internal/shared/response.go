package shared

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type APIMessage struct {
	Status  string      `json:"status"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type DataResponse struct {
	Items interface{} `json:"items"`
	Meta  interface{} `json:"meta,omitempty"`
}

type Meta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}

func httpStatusToCode(status int) string {
	return strings.ReplaceAll(
		strings.ToUpper(http.StatusText(status)),
		" ",
		"_",
	)
}

func RespondSuccess(w http.ResponseWriter, status int, message string, items interface{}, meta ...Meta) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	var data interface{}

	if len(meta) > 0 {
		data = DataResponse{
			Items: items,
			Meta:  meta[0],
		}
	} else {
		data = DataResponse{
			Items: items,
		}
	}

	json.NewEncoder(w).Encode(APIMessage{
		Status:  "success",
		Code:    httpStatusToCode(status),
		Message: message,
		Data:    data,
	})

	log.Println("success:", message)
}

func RespondError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(APIMessage{
		Status:  "error",
		Code:    httpStatusToCode(status),
		Message: message,
	})
	log.Println("Error:", message)
}
