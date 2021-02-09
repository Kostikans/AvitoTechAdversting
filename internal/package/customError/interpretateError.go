package customError

import (
	"net/http"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/serverError"
)

var convertStatusToHTTP = map[int]int{
	clientError.BadRequest:          http.StatusBadRequest,
	serverError.ServerInternalError: http.StatusInternalServerError,
	clientError.NotFound:            http.StatusNotFound,
}

func StatusCode(code int) int {
	return convertStatusToHTTP[code]
}
