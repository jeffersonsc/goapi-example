package middleware

import (
	"log"
	"net/http"
	"os"
	"time"
)

// AccessLog print and monitore handlers
func AccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log := log.New(os.Stdout, "[api] ", 0)

		next.ServeHTTP(w, r)
		log.Printf("%s: %s - response time %s", r.Method, r.URL.RequestURI(), time.Since(start).String())
	})
}
