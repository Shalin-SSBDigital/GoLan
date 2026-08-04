// ============================================================
//
//	CONTROLLER LAYER -- one handler per endpoint
//
// ============================================================
//
//	The "C" in MVC. Controllers are the ONLY layer that touches
//	net/http directly. Each controller method:
//	  1. parses the request (JSON body, path vars)
//	  2. calls the data layer (the model)
//	  3. hands a status code + data to the views layer
//
//	Controllers NEVER store data themselves -- they delegate to
//	*data.TaskStore. That split is what makes the layers swap-able
//	(e.g. swap the store for Postgres without touching handlers).
//
// ============================================================
package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golan/rest-api/data"
	"golan/rest-api/views"
)

// TaskController is our controller: think of it as the "class"
// whose methods are the REST endpoints. It holds the store (the
// model) so every handler can reach it.
type TaskController struct {
	store *data.TaskStore
}

// NewTaskController wires the controller to its store.
func NewTaskController(store *data.TaskStore) *TaskController {
	return &TaskController{store: store}
}

// Routes registers every endpoint on the mux. The router here is
// Go 1.22+'s improved http.ServeMux, which understands
// "METHOD /pattern" strings:
//
//	"GET /tasks"       -> list tasks
//	"POST /tasks"      -> create a task
//	"GET /tasks/{id}"  -> get one task   ({id} = path wildcard)
//	"PUT /tasks/{id}"  -> full update
//	"PATCH /tasks/{id}"-> toggle Done on/off
//	"DELETE /tasks/{id}"-> delete task
func (c *TaskController) Routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /tasks", c.ListTasks)
	mux.HandleFunc("POST /tasks", c.CreateTask)
	mux.HandleFunc("GET /tasks/{id}", c.GetTask)
	mux.HandleFunc("PUT /tasks/{id}", c.UpdateTask)
	mux.HandleFunc("PATCH /tasks/{id}", c.ToggleTask)
	mux.HandleFunc("DELETE /tasks/{id}", c.DeleteTask)
}

// ---- endpoint methods ------------------------------------------

// ListTasks: GET /tasks      -> 200 + JSON array
func (c *TaskController) ListTasks(w http.ResponseWriter, r *http.Request) {
	views.WriteJSON(w, http.StatusOK, c.store.List())
}

// CreateTask: POST /tasks    -> 201 + the new task
//
//	Body: {"title": "..."}
func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	// 1. Decode the JSON body into a struct.
	var input struct {
		Title string `json:"title"`
	}
	// A Decoder errors if the JSON is malformed.
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		views.WriteError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	// 2. Validate.
	if input.Title == "" {
		views.WriteError(w, http.StatusBadRequest, "title is required")
		return
	}

	// 3. Delegate to the data layer, then render the result.
	task := c.store.Add(input.Title)
	views.WriteJSON(w, http.StatusCreated, task) // 201 Created
}

// GetTask: GET /tasks/{id}   -> 200 + task, or 404
func (c *TaskController) GetTask(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	task, found := c.store.Find(id)
	if !found {
		views.WriteError(w, http.StatusNotFound, "task not found")
		return
	}
	views.WriteJSON(w, http.StatusOK, task)
}

// UpdateTask: PUT /tasks/{id}-> 200 + updated task, or 404
//
//	Body: {"title": "...", "done": true}
func (c *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	var input struct {
		Title string `json:"title"`
		Done  bool   `json:"done"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		views.WriteError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if input.Title == "" {
		views.WriteError(w, http.StatusBadRequest, "title is required")
		return
	}
	task, found := c.store.Update(id, input.Title, input.Done)
	if !found {
		views.WriteError(w, http.StatusNotFound, "task not found")
		return
	}
	views.WriteJSON(w, http.StatusOK, task)
}

// ToggleTask: PATCH /tasks/{id} -> flips Done, then 200 + task
// A convenience so clients don't need a full PUT to tick a box.
func (c *TaskController) ToggleTask(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	current, found := c.store.Find(id)
	if !found {
		views.WriteError(w, http.StatusNotFound, "task not found")
		return
	}
	task, _ := c.store.Update(id, current.Title, !current.Done)
	views.WriteJSON(w, http.StatusOK, task)
}

// DeleteTask: DELETE /tasks/{id} -> 204 No Content, or 404
func (c *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	if !c.store.Remove(id) {
		views.WriteError(w, http.StatusNotFound, "task not found")
		return
	}
	w.WriteHeader(http.StatusNoContent) // 204 -- no body needed
}

// ---- controller helpers ------------------------------------------

// parseID reads the {id} path value and converts it to an int.
// It already wrote the error response on failure, so the caller
// can just `return` when ok is false.
func parseID(w http.ResponseWriter, r *http.Request) (int, bool) {
	raw := r.PathValue("id") // Go 1.22+: value of {id} in the pattern
	id, err := strconv.Atoi(raw)
	if err != nil || id <= 0 {
		views.WriteError(w, http.StatusBadRequest, "invalid task id")
		return 0, false
	}
	return id, true
}
