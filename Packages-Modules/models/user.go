// ============================================================
//  PACKAGE MODELS — demonstrates package with custom types
// ============================================================
//  This package defines data types (structs) used by the app.
//
//  In Python, you might have:
//    models/
//    ├── __init__.py
//    └── user.py      # class User: ...
//
//  In Go:
//    models/
//    └── user.go      # package models
//
//  Note: Go folders are flat — no nested __init__.py needed.
//        All .go files in models/ get package models.

package models

import "fmt"

// =============================================================================
// EXPORTED TYPES (public structs)
// =============================================================================

// User is exported — uppercase U
// Python: class User: (always public)
type User struct {
	// Exported fields — accessible from other packages
	Name  string
	Email string
	Age   int

	// Unexported field — only accessible within models package
	// Python: self._internal_id = ...
	internalID string
}

// NewUser is a constructor (factory) — exported because uppercase
func NewUser(name, email string, age int) *User {
	return &User{
		Name:        name,
		Email:       email,
		Age:         age,
		internalID:  generateID(),  // calling unexported function
	}
}

// Greet is a method on *User — exported
func (u *User) Greet() string {
	return fmt.Sprintf("Hi, I'm %s (%s)", u.Name, u.Email)
}

// =============================================================================
// UNEXPORTED TYPES (private)
// =============================================================================

// internalConfig is unexported — cannot be used outside models package
type internalConfig struct {
	DatabaseURL string
	Timeout     int
}

// unexported function — generates internal IDs
// Only callable from within the models package
func generateID() string {
	return "usr_" + fmt.Sprintf("%d", counter)
}

// =============================================================================
// PACKAGE-LEVEL STATE
// =============================================================================

var counter int

func init() {
	counter = 1000
	fmt.Println("[models] package initialized, starting ID counter at", counter)
}
