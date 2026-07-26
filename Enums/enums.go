// ============================================================
//  ENUMS (iota) IN GO — with Python Comparisons
// ============================================================
//  Go has NO `enum` keyword. Instead, it uses `iota` — an
//  auto-incrementing constant generator within `const` blocks.
//
//  ┌──────────────────────┬──────────────────────────────────┐
//  │         Go           │             Python               │
//  ├──────────────────────┼──────────────────────────────────┤
//  │ const (              │ from enum import Enum            │
//  │     A = iota         │ class Color(Enum):               │
//  │     B               │     RED = 1                      │
//  │     C               │     GREEN = 2                    │
//  │ )                    │     BLUE = 3                    │
//  ├──────────────────────┼──────────────────────────────────┤
//  │ iota starts at 0     │ enum auto() starts at 1          │
//  │ iota resets per      │ enum members are unique          │
//  │   const block        │   singleton instances            │
//  └──────────────────────┴──────────────────────────────────┘
// ============================================================

package main

import (
	"fmt"
	"strings"
)

// =============================================================================
// 1. BASIC iota — AUTO-INCREMENTING CONSTANTS
// =============================================================================
// `iota` is a special identifier that increments with each constant
// declaration within a `const` block.
//
// Python: from enum import Enum, auto
//         class Status(Enum):
//             PENDING = auto()  # 1
//             ACTIVE = auto()   # 2
//             INACTIVE = auto() # 3

const (
	Pending Status = iota // 0
	Active                // 1
	Inactive              // 2
)

func demoBasicIota() {
	fmt.Println("=== Basic iota ===")
	fmt.Printf("Pending  = %d\n", Pending)
	fmt.Printf("Active   = %d\n", Active)
	fmt.Printf("Inactive = %d\n", Inactive)
	fmt.Println("iota starts at 0 (Python's enum.auto() starts at 1)")
}

// =============================================================================
// 2. CUSTOM TYPE FOR ENUMS
// =============================================================================
// Best practice: define a custom type for your enum.
// This gives type safety — you can't accidentally pass a regular int.
//
// Python equivalent:
//   class Status(Enum):
//       PENDING = 0
//       ACTIVE = 1
//       INACTIVE = 2
//
//   def process(s: Status): ...

type Status int

// Using Status type — only Status values can be passed
func describe(s Status) string {
	switch s {
	case Pending:
		return "waiting to start"
	case Active:
		return "currently running"
	case Inactive:
		return "stopped"
	default:
		return "unknown status"
	}
}

func demoCustomType() {
	fmt.Println("\n=== Custom Type for Enums ===")

	fmt.Printf("describe(Pending)  = %q\n", describe(Pending))
	fmt.Printf("describe(Active)   = %q\n", describe(Active))
	fmt.Printf("describe(Inactive) = %q\n", describe(Inactive))

	// Type safety: this is a compile error:
	// describe(42)  // ❌ cannot use 42 (type int) as type Status

	// Passing Status via int is a compile error:
	// var n int = 1
	// describe(n)  // ❌ cannot use n (type int) as type Status

	fmt.Println("  Python enum has the same type safety")
	fmt.Println("  describe(Status.ACTIVE) ✅")
	fmt.Println("  describe(1)  ❌ TypeError in Python too")
}

// =============================================================================
// 3. iota WITH STRING REPRESENTATIONS
// =============================================================================
// Python has __str__ for custom string representation.
// Go: implement the String() method on your enum type.

func (s Status) String() string {
	switch s {
	case Pending:
		return "PENDING"
	case Active:
		return "ACTIVE"
	case Inactive:
		return "INACTIVE"
	default:
		return fmt.Sprintf("Status(%d)", s)
	}
}

func demoStringMethod() {
	fmt.Println("\n=== String Representations ===")

	// With String() method, %s and %v print nicely
	fmt.Printf("Status values as strings:\n")
	fmt.Printf("  %s  (%%s)\n", Pending)
	fmt.Printf("  %s  (%%s)\n", Active)
	fmt.Printf("  %s  (%%s)\n", Inactive)

	// Without String(), %s would print "0", "1", "2"
	fmt.Println("\nPython equivalent: def __str__(self): ...")
}

// =============================================================================
// 4. iota SKIPPING VALUES
// =============================================================================
// Use `_` to skip values, or use expressions.

const (
	_  = iota             // skip 0
	KB       = 1 << (10 * iota) // 1 << 10 = 1024
	MB                          // 1 << 20 = 1048576
	GB                          // 1 << 30 = 1073741824
	TB                          // 1 << 40
)

func demoSkipValues() {
	fmt.Println("\n=== Skipping Values with iota ===")
	fmt.Printf("KB = %d bytes\n", KB)
	fmt.Printf("MB = %d bytes\n", MB)
	fmt.Printf("GB = %d bytes\n", GB)
	fmt.Printf("TB = %d bytes\n", TB)
}

// =============================================================================
// 5. BITMASK ENUMS (FLAGS)
// =============================================================================
// Python: class Permission(Flag):
//             NONE = 0
//             READ = auto()
//             WRITE = auto()
//             EXECUTE = auto()
//
// Go: Use iota with bit shifts to create flag values.

type Permission uint8

const (
	None     Permission = 0
	Read     Permission = 1 << iota // 1 (0001)
	Write                            // 2 (0010)
	Execute                          // 4 (0100)
	Delete                           // 8 (1000)
)

// String for Permission flags
func (p Permission) String() string {
	var flags []string
	if p&Read != 0 {
		flags = append(flags, "READ")
	}
	if p&Write != 0 {
		flags = append(flags, "WRITE")
	}
	if p&Execute != 0 {
		flags = append(flags, "EXECUTE")
	}
	if p&Delete != 0 {
		flags = append(flags, "DELETE")
	}
	if len(flags) == 0 {
		return "NONE"
	}
	return strings.Join(flags, "|")
}

func demoBitmaskEnums() {
	fmt.Println("\n=== Bitmask Enums (Flags) ===")

	// Single permissions
	fmt.Printf("Read    = %04b = %s\n", Read, Read)
	fmt.Printf("Write   = %04b = %s\n", Write, Write)
	fmt.Printf("Execute = %04b = %s\n", Execute, Execute)
	fmt.Printf("Delete  = %04b = %s\n", Delete, Delete)

	// Combined permissions using |
	readWrite := Read | Write
	fmt.Printf("\nRead|Write = %04b = %s\n", readWrite, readWrite)

	fullAccess := Read | Write | Execute | Delete
	fmt.Printf("All flags = %04b = %s\n", fullAccess, fullAccess)

	// Check if a permission is set (AND)
	hasWrite := readWrite&Write != 0
	hasDelete := readWrite&Delete != 0
	fmt.Printf("\nreadWrite has Write:  %t\n", hasWrite)
	fmt.Printf("readWrite has Delete: %t\n", hasDelete)

	// Remove a permission (AND NOT)
	readOnly := readWrite &^ Write
	fmt.Printf("readWrite &^ Write = %04b = %s\n", readOnly, readOnly)
}

// =============================================================================
// 6. iota + EXPRESSIONS
// =============================================================================
// iota can be used in expressions — it increments per line, not per constant.
// This enables interesting patterns.

const (
	// Days of week (starting from Monday = 1)
	Monday    = iota + 1 // 1
	Tuesday              // 2
	Wednesday            // 3
	Thursday             // 4
	Friday               // 5
	Saturday             // 6
	Sunday               // 7
)

// Weekday type with String()
type Weekday int

const (
	Mon Weekday = iota + 1
	Tue
	Wed
	Thu
	Fri
	Sat
	Sun
)

func (d Weekday) String() string {
	days := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	if d >= 1 && int(d) <= len(days) {
		return days[d-1]
	}
	return fmt.Sprintf("Day(%d)", d)
}

func demoIotaExpressions() {
	fmt.Println("\n=== iota with Expressions ===")
	fmt.Printf("Monday    = %d\n", Monday)
	fmt.Printf("Tuesday   = %d\n", Tuesday)
	fmt.Printf("Wednesday = %d\n", Wednesday)
	fmt.Printf("Thursday  = %d\n", Thursday)
	fmt.Printf("Friday    = %d\n", Friday)
	fmt.Printf("Saturday  = %d\n", Saturday)
	fmt.Printf("Sunday    = %d\n", Sunday)

	fmt.Println("\nWith String() method:")
	for d := Mon; d <= Sun; d++ {
		fmt.Printf("  %s\n", d)
	}
}

// =============================================================================
// 7. COMPLETE COMPARISON TABLE
// =============================================================================
//
// ┌──────────────────────────────┬──────────────────────────────┬────────────────────────────┐
// │          Feature             │          Go                  │          Python            │
// ├──────────────────────────────┼──────────────────────────────┼────────────────────────────┤
// │ Enum keyword                 │ ❌ No `enum` keyword         │ ✅ from enum import Enum   │
// │ Auto-increment               │ iota (starts at 0)           │ auto() (starts at 1)       │
// │ Custom type                  │ type Status int              │ class Status(Enum)         │
// │ String representation        │ func (s Status) String()     │ def __str__(self)          │
// │ Bitmask / flags              │ Permission = 1 << iota       │ class P(Flag): READ = auto │
// │ Reset per const block        │ ✅ Automatic                 │ ✅ Separate class          │
// │ Start at specific number     │ iota + N or iota - N         │ Enum member = N            │
// │ Skip value                   │ _ = iota                     │ Not needed                 │
// │ Switch on enum               │ switch s { case A: ...}     │ match s: case A: ...       │
// │ Iterate all values           │ Manual (list them)           │ list(Color)                │
// │ Type safety                  │ ✅ Compile-time              │ ✅ Runtime                 │
// └──────────────────────────────┴──────────────────────────────┴────────────────────────────┘

// =============================================================================
// MAIN
// =============================================================================

func main() {
	demoBasicIota()
	demoCustomType()
	demoStringMethod()
	demoSkipValues()
	demoBitmaskEnums()
	demoIotaExpressions()

	fmt.Println("\n" + strings.Repeat("=", 49))
	fmt.Println("ENUMS (IOTA) CHEAT SHEET")
	fmt.Println("=" + strings.Repeat("=", 49))
	fmt.Println(`
  // Basic enum with custom type
  type Status int
  const (
      Pending Status = iota  // 0
      Active                 // 1
      Inactive               // 2
  )

  // With String() method
  func (s Status) String() string { ... }

  // Bitmask (flags)
  type Permission uint8
  const (
      Read Permission = 1 << iota  // 1
      Write                        // 2
      Execute                      // 4
  )

  // Skip 0, start with expression
  const (
      _  = iota
      KB = 1 << (10 * iota)  // 1024
      MB = 1 << (20 * iota)  // 1048576
  )

  // Start at specific value
  const (
      Mon = iota + 1  // 1
      Tue             // 2
  )`)
}

// =============================================================================
// PRACTICE EXERCISES
// =============================================================================
//
// Easy: Define an enum type `Day` with iota starting from 1 for Mon-Sun.
//       Add a String() method that returns "Monday", "Tuesday", etc.
//       Print all days.
//
// Medium: Define a bitmask enum `Access` with Read=1, Write=2, Admin=4.
//         Write a function checkAccess(required, actual Access) bool.
//         Test: does actual=Read|Write satisfy required=Admin?
//
// Challenging: Create an enum `LogLevel` with Debug, Info, Warn, Error
//             using iota. Implement String(). Write a function
//             shouldLog(level, minLevel LogLevel) bool that returns
//             true if level >= minLevel. Make it work with >= by
//             ensuring iota order matches severity order.
