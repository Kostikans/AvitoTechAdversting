package abstractResponse

import (
	json "encoding/json"
	"net/http"

	easyjson "github.com/mailru/easyjson"
)

func SendDataResponse(w http.ResponseWriter, code int, data interface{}) {
	response := HttpResponse{Data: data, Code: code}
	_, _, err := easyjson.MarshalToHTTPResponseWriter(response, w)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
}

func SendOkResponse(w http.ResponseWriter, code int) {
	err := json.NewEncoder(w).Encode(HttpResponse{Code: code})
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
}
