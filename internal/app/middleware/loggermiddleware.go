package apiMiddleware

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/Kostikans/AvitoTechadvertising/internal/package/logger"
	"github.com/gorilla/mux"
)

func LoggerMiddleware(log *logger.CustomLogger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if strings.Contains(req.URL.Path, "docs") {
				next.ServeHTTP(w, req)
				return
			}

			rand.Seed(time.Now().UnixNano())
			id := fmt.Sprintf("%016x", rand.Int())[:5]

			log.StartReq(*req, id)
			start := time.Now()
			next.ServeHTTP(w, req)

			respTime := time.Since(start)
			log.EndReq(respTime.Microseconds(), req.Context())

		})
	}
}
