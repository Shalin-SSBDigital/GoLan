// ============================================================
//
//	VIEW LAYER -- how data is rendered to the client
//
// ============================================================
//
//	The "V" in MVC. In a JSON API the "view" is the JSON
//	serializer -- there is no HTML page, the client renders data.
//
//	Python equivalent:  json.dumps(obj) + setting response headers
//
//	NOTE: we set the Content-Type header BEFORE WriteHeader --
//	once the response starts flowing to the client, headers are
//	locked in.
//
// ============================================================
package views

import (
	"encoding/json"
	"net/http"
)

// WriteJSON turns any value into a JSON response with the right
// Content-Type header and status code (200, 201, 404, ...).
func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)        // e.g. 200 OK, 201 Created
	json.NewEncoder(w).Encode(v) // stream the JSON to the client
}

// WriteError sends a uniform JSON error body:
//
//	{"error": "some message"}
//
// Keeping one shape for every error makes clients simple.
func WriteError(w http.ResponseWriter, status int, msg string) {
	WriteJSON(w, status, map[string]string{"error": msg})
}
