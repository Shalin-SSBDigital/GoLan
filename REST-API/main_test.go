// ============================================================
//  TESTS -- httptest makes testing handlers easy
// ============================================================
//  We test through the PUBLIC HTTP interface (routes + JSON),
//  NOT by calling Go functions directly. httptest.NewRecorder
//  pretends to be a real HTTP client/server connection, so we
//  can send requests and inspect responses without opening
//  a port.
//
//  Run with:  go test ./...        (from the REST-API folder)
//
//  Note how the test only touches the exported layers
//  (data.NewTaskStore, controller.Routes, middleware) -- exactly
//  the way an external client would use the API.
// ============================================================

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golan/rest-api/controllers"
	"golan/rest-api/data"
	"golan/rest-api/middleware"
	"golan/rest-api/models"
)

// newTestAPI builds a fresh store + router for every test, so
// tests never share state. Each test gets a clean "database".
func newTestAPI() http.Handler {
	store := data.NewTaskStore()
	controller := controllers.NewTaskController(store)
	mux := http.NewServeMux()
	controller.Routes(mux)
	return middleware.Logger(mux)
}

// doJSON is a tiny test helper: it performs a request with an
// optional JSON body and returns the recorder.
func doJSON(h http.Handler, method, path, body string) *httptest.ResponseRecorder {
	var req *http.Request
	if body == "" {
		req = httptest.NewRequest(method, path, nil)
	} else {
		req = httptest.NewRequest(method, path, strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

// ---- the tests ---------------------------------------------------

func TestListTasks(t *testing.T) {
	api := newTestAPI()

	rec := doJSON(api, "GET", "/tasks", "")

	if rec.Code != http.StatusOK {
		t.Fatalf("want 200, got %d", rec.Code)
	}
	var tasks []models.Task
	if err := json.Unmarshal(rec.Body.Bytes(), &tasks); err != nil {
		t.Fatalf("bad JSON: %v", err)
	}
	if len(tasks) != 1 { // the seeded "Learn Go REST APIs" task
		t.Fatalf("want 1 task, got %d", len(tasks))
	}
}

func TestCreateTask(t *testing.T) {
	api := newTestAPI()

	rec := doJSON(api, "POST", "/tasks", `{"title":"Buy milk"}`)

	if rec.Code != http.StatusCreated { // 201, not 200!
		t.Fatalf("want 201, got %d", rec.Code)
	}
	var task models.Task
	if err := json.Unmarshal(rec.Body.Bytes(), &task); err != nil {
		t.Fatalf("bad JSON: %v", err)
	}
	if task.Title != "Buy milk" || task.ID != 2 || task.Done {
		t.Fatalf("unexpected task: %+v", task)
	}
}

func TestCreateTaskMissingTitle(t *testing.T) {
	api := newTestAPI()

	rec := doJSON(api, "POST", "/tasks", `{"title":""}`)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("want 400, got %d", rec.Code)
	}
}

func TestGetTask(t *testing.T) {
	api := newTestAPI()

	rec := doJSON(api, "GET", "/tasks/1", "")

	if rec.Code != http.StatusOK {
		t.Fatalf("want 200, got %d", rec.Code)
	}
	var task models.Task
	json.Unmarshal(rec.Body.Bytes(), &task)
	if task.ID != 1 || task.Title != "Learn Go REST APIs" {
		t.Fatalf("unexpected task: %+v", task)
	}
}

func TestGetTaskNotFound(t *testing.T) {
	api := newTestAPI()

	rec := doJSON(api, "GET", "/tasks/999", "")

	if rec.Code != http.StatusNotFound {
		t.Fatalf("want 404, got %d", rec.Code)
	}
}

func TestUpdateTask(t *testing.T) {
	api := newTestAPI()

	rec := doJSON(api, "PUT", "/tasks/1", `{"title":"Learn Gin","done":true}`)

	if rec.Code != http.StatusOK {
		t.Fatalf("want 200, got %d", rec.Code)
	}
	var task models.Task
	json.Unmarshal(rec.Body.Bytes(), &task)
	if task.Title != "Learn Gin" || !task.Done {
		t.Fatalf("unexpected task: %+v", task)
	}
}

func TestToggleTask(t *testing.T) {
	api := newTestAPI()

	// The seeded task starts as Done=false.
	rec := doJSON(api, "PATCH", "/tasks/1", "")

	if rec.Code != http.StatusOK {
		t.Fatalf("want 200, got %d", rec.Code)
	}
	var task models.Task
	json.Unmarshal(rec.Body.Bytes(), &task)
	if !task.Done {
		t.Fatal("want Done=true after toggle")
	}
}

func TestDeleteTask(t *testing.T) {
	api := newTestAPI()

	rec := doJSON(api, "DELETE", "/tasks/1", "")

	if rec.Code != http.StatusNoContent { // 204
		t.Fatalf("want 204, got %d", rec.Code)
	}
	// And now it should be gone:
	rec2 := doJSON(api, "GET", "/tasks/1", "")
	if rec2.Code != http.StatusNotFound {
		t.Fatalf("want 404 after delete, got %d", rec2.Code)
	}
}

func TestInvalidID(t *testing.T) {
	api := newTestAPI()

	rec := doJSON(api, "GET", "/tasks/abc", "")

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("want 400, got %d", rec.Code)
	}
}
