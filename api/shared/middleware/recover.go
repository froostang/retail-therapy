package middleware

import (
	"log"
	"net/http"
	"runtime"
)

// TODO: Try prompting for test creation
func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic occurred: %v", runtime.StartTrace().Error())
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
