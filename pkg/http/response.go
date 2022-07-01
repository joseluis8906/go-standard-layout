package http

import (
	"encoding/json"
	"net/http"
)

type (
	response struct {
		Data   interface{} `json:"data,omitempty"`
		Error  string      `json:"error,omitempty"`
		Status int         `json:"status,omitempty"`
	}
)

func JSON(w http.ResponseWriter, status int, data interface{}, err error) error {
	res := response{
		Status: status,
		Data:   data,
		Error: func() string {
			if err != nil {
				return err.Error()
			}

			return ""
		}(),
	}

	w.WriteHeader(status)
	err = json.NewEncoder(w).Encode(res)
	return err
}
