// ============================================================
//
//	DATA LAYER -- where Tasks are stored
//
// ============================================================
//
//	Part of the MODEL concern: persistence. A real app would use
//	Postgres or SQLite here; we use a slice.
//
//	WHY THE Mutex? Go's web server runs each request in its OWN
//	goroutine. Two requests could write to `tasks` at the SAME
//	time and corrupt the slice. sync.Mutex guarantees only ONE
//	goroutine touches the data at any instant -- the #1 thing
//	that separates a toy Go server from a production one.
//
//	NOTE: NO HTTP knowledge lives here. The store knows nothing
//	about requests or JSON -- that is what makes it a clean,
//	testable data layer.
//
// ============================================================
package data

import (
	"sync"
	"time"

	"golan/rest-api/models"
)

// TaskStore is the persistence layer for models.Task.
type TaskStore struct {
	mu     sync.Mutex    // EVERY method must Lock() ... Unlock()
	tasks  []models.Task // our "table"
	nextID int           // auto-increment primary key
}

// NewTaskStore builds a fresh store and seeds one example task
// so GET /tasks returns data immediately.
func NewTaskStore() *TaskStore {
	s := &TaskStore{nextID: 1}
	s.tasks = append(s.tasks, models.Task{
		ID:        s.nextID,
		Title:     "Learn Go REST APIs",
		CreatedAt: time.Now(),
	})
	s.nextID++ // next insert gets id = 2
	return s
}

// Add inserts a task and returns it (now with its real ID).
// Lock() + defer Unlock(): the mutex is released even if the
// function panics, because defer ALWAYS runs on the way out.
func (s *TaskStore) Add(title string) models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	t := models.Task{
		ID:        s.nextID,
		Title:     title,
		CreatedAt: time.Now(),
	}
	s.tasks = append(s.tasks, t)
	s.nextID++
	return t
}

// Find returns a COPY of the task, with bool = true if found.
// The "comma-ok" idiom (value, found) is everywhere in Go --
// never use a sentinel like -1 to mean "missing".
func (s *TaskStore) Find(id int) (models.Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, t := range s.tasks {
		if t.ID == id {
			return t, true
		}
	}
	return models.Task{}, false
}

// List returns a COPY of every task. The copy matters: a caller
// cannot accidentally mutate the store's backing array.
func (s *TaskStore) List() []models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]models.Task, len(s.tasks))
	copy(out, s.tasks)
	return out
}

// Update overwrites Title and Done inside the store and returns
// the new task plus a found flag (like Find).
func (s *TaskStore) Update(id int, title string, done bool) (models.Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks[i].Title = title
			s.tasks[i].Done = done
			return s.tasks[i], true
		}
	}
	return models.Task{}, false
}

// Remove deletes a task. The bool reports whether it existed, so
// the controller can answer 404 when it did not.
func (s *TaskStore) Remove(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, t := range s.tasks {
		if t.ID == id {
			// "Slice trick" to cut element i out:
			//   tasks[:i]  (everything before i)
			//   tasks[i+1:]...  (everything after i)
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return true
		}
	}
	return false
}
