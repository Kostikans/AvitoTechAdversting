package customError

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

type CustomError struct {
	Code int
	File string
	Line int
	Err  error
}

func NewCustomError(err error, code int, skip int) error {
	_, fn, line, _ := runtime.Caller(skip)
	return fmt.Errorf("code:[%d] file:[%s]  line:[%d]  error:[%w]", code, fn, line, err)
}

func ParseCode(err error) int {
	index := strings.Index(err.Error(), "]")
	code, err := strconv.Atoi(err.Error()[6:index])
	if err != nil {
		return http.StatusInternalServerError
	}
	return code
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("code:[%d] file:[%s]  line:[%d]  error:[%s]", e.Code, e.File, e.Line, e.Err)
}
