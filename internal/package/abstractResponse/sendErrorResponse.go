package abstractResponse

import (
	json "encoding/json"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, code int, msg string) {
	error := Error{Msg: msg}
	err := json.NewEncoder(w).Encode(HttpResponse{Err: &error, Code: code})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
