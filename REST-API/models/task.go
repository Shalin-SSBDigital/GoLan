// ============================================================
//
//	MODEL -- what a Task IS
//
// ============================================================
//
//	The "M" in MVC. A plain data structure with no HTTP logic.
//	Think of it as one row in a database table.
//
//	The `json:"..."` tags map Go field names to JSON keys:
//	  Go field       JSON on the wire
//	  ---------      ----------------
//	  ID             "id"
//	  Title          "title"
//	  Done           "done"
//	  CreatedAt      "created_at"   (snake_case is common JSON style)
//
//	WITHOUT a tag, Go would use the Go name as-is ("CreatedAt").
//
// ============================================================
package models

import "time"

// Task represents a single to-do item.
type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}
