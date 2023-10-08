go
package middleware

import (
	"log"
	"net/http"
)

func Middleware(authHandler http.Handler, errorHandler http.Handler) http.Handler {
	return AuthMiddleware(ErrorMiddleware(AdditionalMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add your custom middleware logic here
		log.Println("Executing additional middleware logic")

		// Call the authentication handler
		authHandler.ServeHTTP(w, r)
	}))))
}

func AdditionalMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add your custom middleware logic here
		log.Println("Executing additional middleware logic")

		// Pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}
