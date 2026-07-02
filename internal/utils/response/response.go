package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOk    = "OK"
	StatusError = "ERROR"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}
func GenError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}
func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field is required: %s", err.Field()))
		case "default":
			errMsgs = append(errMsgs, fmt.Sprintf("field is invalid: %s", err.Field()))
		}
	}
	return Response{
		Status: StatusError,
		Error:  "validation failed",
	}
}
