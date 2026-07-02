package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/sanket2004-b/REST-API-GO/internal/types"
	"github.com/sanket2004-b/REST-API-GO/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		slog.Info("Welcome this is first go api")
		w.Write([]byte("Welcome this is first go api"))
		// err := json.NewDecoder(r.Body).Decode(&student)

		// if err != nil {
		// 	response.WriteJson(w, http.StatusBadRequest, response.GenError(err))
		// 	return
		// }
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GenError(fmt.Errorf("empty body")))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GenError(err))
			return
		}
	}
}
