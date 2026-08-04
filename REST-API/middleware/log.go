// ============================================================
//
//	MIDDLEWARE -- code that wraps every request
//
// ============================================================
//
//	Go handlers are just func(w, r). A middleware is a function
//	that takes a handler and RETURNS a new handler that does
//	something "around" the original:
//
//	   BEFORE handler  ->  call next  ->  AFTER handler
//
//	Example realization of the connection:
//	  the handler is the entry point of a program. Calling
//	  next.ServeHTTP(w, r) is like invoking a function that
//	  the server ("main loop") then completes.
//
// ============================================================
package middleware

import (
	"log"
	"net/http"
)

// Logger logs each request line to the console, e.g.:
//
//	2024/01/01 12:00:00 GET /tasks
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// BEFORE the real handler runs:
		log.Printf("%s %s", r.Method, r.URL.Path)
		// Call the wrapped handler (the controller/router).
		next.ServeHTTP(w, r)
	})
}
