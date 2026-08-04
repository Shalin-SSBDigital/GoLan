// ============================================================
//
//	MAIN -- composition root (wiring) + server startup
//
// ============================================================
//
//	main.go does NOT contain business logic. It just BUILDS the
//	three MVC layers and glues them together -- the "composition
//	root". In a bigger app, dependency injection would happen
//	here (or a framework would do it for you).
//
//	Build order (bottom-up, each layer depends only on the one
//	below it):
//
//	 data.NewTaskStore()       -> MODEL   (persistence)
//	 controllers.NewTaskController(store)  -> CONTROLLER
//	 controller.Routes(mux)    -> register every endpoint
//	 middleware.Logger(mux)    -> wrap with middleware
//	 http.ListenAndServe(...)  -> serve it
//
// ============================================================
package main

import (
	"log"
	"net/http"
	"os"

	"golan/rest-api/controllers"
	"golan/rest-api/data"
	"golan/rest-api/middleware"
)

func main() {
	// 1. Build the MODEL layer.
	store := data.NewTaskStore()

	// 2. Build the CONTROLLER layer and register its routes.
	controller := controllers.NewTaskController(store)
	mux := http.NewServeMux()
	controller.Routes(mux)

	// 3. Wrap everything in middleware (runs on every request).
	handler := middleware.Logger(mux)

	// 4. Pick a port (override with: PORT=9999 go run .)
	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	// 5. Serve. http.ListenAndServe BLOCKS forever. If it ever
	//    returns, the server failed -- so log fatal.
	log.Printf("API listening on http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
