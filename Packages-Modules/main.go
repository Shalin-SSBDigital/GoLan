// ============================================================
//  PACKAGES & MODULES — Main Entry Point
// ============================================================
//  This file is `package main` — the special package that
//  produces an executable.
//
//  A `package main` MUST have a `func main()` — this is where
//  the program starts executing.
//
//  ┌───────────────────────┬─────────────────────────────────┐
//  │         Go            │             Python              │
//  ├───────────────────────┼─────────────────────────────────┤
//  │ package main          │ if __name__ == '__main__':      │
//  │ func main() { ... }   │     main()                      │
//  ├───────────────────────┼─────────────────────────────────┤
//  │ import "fmt"          │ import sys                      │
//  │ import "os"           │ import os                       │
//  │ You import by PATH    │ You import by MODULE NAME       │
//  └───────────────────────┴─────────────────────────────────┘

package main

// =============================================================================
// IMPORTING PACKAGES
// =============================================================================
// Go imports are ALWAYS absolute paths (relative to the module root).
//
// Python:
//   from calculator import add, subtract
//   from models.user import User
//   import utils.strings as ut_str
//
// Go:
//   import "golan/packages-modules/calculator"
//   import "golan/packages-modules/models"
//
//   You call:  calculator.Add(1, 2)
//   NOT:       calculator.calculator.Add(1, 2)
//
// KEY RULE: Import path = directory path from module root.
//   The package NAME can differ from the directory NAME (but usually matches).
//
// =============================================================================
// STANDARD vs THIRD-PARTY vs LOCAL IMPORTS
// =============================================================================
// Go organizes imports into three groups (separated by blank lines):
//   1. Standard library  ("fmt", "os", "strings")
//   2. Third-party       ("github.com/gin-gonic/gin")
//   3. Local             ("golan/packages-modules/calculator")
//
// This is enforced by `go fmt` (and `goimports` does it automatically).
// Python's equivalent: standard lib → third-party → local in import blocks.

import (
	// Standard library
	"fmt"
	"strings"

	// Local packages (imported by MODULE PATH + SUBDIRECTORY)
	"golan/packages-modules/calculator"
	"golan/packages-modules/config"
	"golan/packages-modules/models"
	"golan/packages-modules/utils"
)

// =============================================================================
// DEMO 1: Basic Imports
// =============================================================================
// Each imported package has its own NAMESPACE.
// You access exported names with: packageName.ExportedName
//
// Python:
//   import calculator
//   result = calculator.add(1, 2)
//
// Go:
//   import "golan/packages-modules/calculator"
//   result := calculator.Add(1, 2)

func demoBasicImport() {
	fmt.Println("\n=== Demo 1: Basic Imports ===")
	fmt.Println("Calling calculator.Add(10, 5):", calculator.Add(10, 5))
	fmt.Println("Calling calculator.Subtract(10, 5):", calculator.Subtract(10, 5))

	result, err := calculator.Divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Calling calculator.Divide(10, 3):", result)
	}

	fmt.Println("Calculator version:", calculator.Version)
	fmt.Println("Calculator author:", calculator.GetAuthor())

	// Cannot access calculator.author — compile error!
	// calculator.author is unexported (lowercase)
	// fmt.Println(calculator.author)  // ❌ WOULD NOT COMPILE
}

// =============================================================================
// DEMO 2: Using Custom Types from Another Package
// =============================================================================
// Python:
//   from models.user import User
//   u = User("Alice", "a@b.com", 30)

func demoTypesFromPackage() {
	fmt.Println("\n=== Demo 2: Types from Packages ===")

	// Creating a User using the exported constructor
	u := models.NewUser("Alice", "alice@example.com", 30)
	fmt.Println(u.Greet())

	// Direct struct literal (all fields must be accessible)
	u2 := models.User{
		Name:  "Bob",
		Email: "bob@example.com",
		Age:   25,
	}
	fmt.Println(u2.Greet())

	// NOTE: Cannot set u.internalID — it's unexported
	// u.internalID = "abc"  // ❌ WOULD NOT COMPILE
}

// =============================================================================
// DEMO 3: Multi-File Package
// =============================================================================
// Both strings.go and math.go are in the SAME utils package.
// They share the namespace — you call utils.Reverse AND utils.IsEven.
//
// Python:
//   from utils.strings import reverse
//   from utils.numbers import is_even
//
// In Go, there's NO sub-package — everything is utils.

func demoMultiFilePackage() {
	fmt.Println("\n=== Demo 3: Multi-File Package ===")
	fmt.Println("utils.Reverse('hello'):", utils.Reverse("hello"))
	fmt.Println("utils.CountWords('hello world go'):", utils.CountWords("hello world go"))
	fmt.Println("utils.IsEven(7):", utils.IsEven(7))
	fmt.Println("utils.Max(10, 20):", utils.Max(10, 20))
	fmt.Println("utils.Clamp(50, 0, 100):", utils.Clamp(50, 0, 100))
	fmt.Println("utils.Clamp(150, 0, 100):", utils.Clamp(150, 0, 100))
	fmt.Println("utils.Clamp(-5, 0, 100):", utils.Clamp(-5, 0, 100))

	// Cannot call utils.logUsage or utils.min — unexported!
	// utils.logUsage("test")  // ❌ WOULD NOT COMPILE
}

// =============================================================================
// DEMO 4: init() Execution Order
// =============================================================================
// init() functions run in DEPENDENCY ORDER before main().
//
// For this program, the order is:
//   1. calculator.init()    (no dependencies)
//   2. models.init()        (no dependencies)
//   3. config.init()        (no dependencies)
//   4. utils/strings.init() (no dependencies) — alphabetical order
//   5. utils/math.init()    (no dependencies) — alphabetical order
//   6. main.init()          (if any)
//   7. main()
//
// This is like Python import-time execution:
//   import calculator  # calculator/__init__.py runs
//   import models      # models/__init__.py runs
//   ...
//   if __name__ == '__main__':
//       main()

// This init() runs AFTER all dependency package init() functions
func init() {
	fmt.Println("[main] package initialization")
	fmt.Println("[main] ready to start")

	// Demonstrate that config was loaded by its init()
	cfg := config.GetConfig()
	fmt.Printf("[main] app config: %s running on port %d\n",
		cfg.AppName, cfg.Port)
}

// =============================================================================
// ADVANCED IMPORT PATTERNS (demonstrated but not used directly)
// =============================================================================
// These import patterns exist in Go but are NOT used in this file.
// See the comments for explanation.

// ─────────────────────────────────────────────────────────────
// Pattern 1: Import ALIASING
// ─────────────────────────────────────────────────────────────
// Python:  import very_long_package_name as vlpn
// Go:      import vlpn "module/pkg/very/long/package/name"
//
// Use when:
//   - Package name conflicts with your code
//   - Package name is too long
//   - You want to avoid name collision
//
// Example (commented — would conflict with our current imports):
//   import calc "golan/packages-modules/calculator"
//   calc.Add(1, 2)

// ─────────────────────────────────────────────────────────────
// Pattern 2: DOT IMPORT (import into current namespace)
// ─────────────────────────────────────────────────────────────
// Python:  from calculator import add, subtract
// Go:      import . "golan/packages-modules/calculator"
//
// WARNING: Dot imports are RARE and controversial in Go.
// They're mainly used in TESTING (test helpers) and not
// recommended for production code.
//
// Example (commented):
//   import . "fmt"
//   Println("hello")  // no fmt. prefix needed

// ─────────────────────────────────────────────────────────────
// Pattern 3: BLANK IDENTIFIER IMPORT (_)
// ─────────────────────────────────────────────────────────────
// Python:  import config  # even if unused, no error
// Go:      import _ "module/pkg/config"
//
// Go gives a COMPILE ERROR for unused imports.
// But sometimes you need to IMPORT a package SOLELY for its
// side effects (like init() or driver registration).
//
// Use _ for the package alias — the package is imported but
// nothing is available in your code.
//
// This is needed in Go (NOT Python) because:
//   - Python doesn't complain about unused imports
//   - Go treats unused imports as compile errors
//   - Driver registration (database/sql, image formats) relies on blank imports
//
// Example:
//   import _ "github.com/lib/pq"  // registers PostgreSQL driver
//
// We DON'T need it here because config is used (GetConfig).
// But if you import a package ONLY for init(), use _ .

// =============================================================================
// DEMO 5: What Happens with a BLANK IMPORT
// =============================================================================
// To demonstrate blank imports, imagine we have a package that
// only runs init() with no public API. We import it with _.
//
// For the purpose of this demo, we'll just explain the concept.

func demoBlankImport() {
	fmt.Println("\n=== Demo 5: Blank Identifier Import ===")
	fmt.Println("Import a package for its init() side effects only:")
	fmt.Println("  import _ \"golan/packages-modules/config\"")
	fmt.Println()
	fmt.Println("The package's init() runs, but no names are available in code.")
	fmt.Println("Common use: database driver registration:")
	fmt.Println("  import _ \"github.com/lib/pq\"")
	fmt.Println("  import _ \"github.com/go-sql-driver/mysql\"")
	fmt.Println()
	fmt.Println("Without _ — unused import = COMPILE ERROR.")
	fmt.Println("With _ — Go allows it because you're saying 'I know this is unused'.")
}

// =============================================================================
// DEMO 6: Working with go.mod
// =============================================================================
// The go.mod file defines:
//   1. Module name (used as import path root)
//   2. Go version
//   3. Dependencies
//
// Python equivalent:
//   pyproject.toml  or  setup.py  or  requirements.txt
//
// Commands you need to know:
//
//   go mod init <name>     — Creates go.mod (like poetry init)
//   go mod tidy             — Adds missing, removes unused deps (like poetry lock)
//   go get <pkg>@<version> — Downloads a dependency (like pip install)
//   go mod vendor           — Copies deps into vendor/ folder
//   go list -m all          — Lists all dependencies

func demoGoMod() {
	fmt.Println("\n=== Demo 6: go.mod ===")

	fmt.Println("Current go.mod content:")
	fmt.Println("  module golan/packages-modules")
	fmt.Println("  go 1.22.0")
	fmt.Println("  (no dependencies — this is a pure stdlib project)")
	fmt.Println()
	fmt.Println("To add a dependency:")
	fmt.Println("  go get github.com/gin-gonic/gin")
	fmt.Println("  # This updates go.mod and creates go.sum")
	fmt.Println()
	fmt.Println("Python vs Go dependency files:")
	fmt.Printf("%-30s | %s\n", "Go", "Python")
	fmt.Printf("%-30s | %s\n", "------", "------")
	fmt.Printf("%-30s | %s\n", "go.mod", "pyproject.toml / setup.py")
	fmt.Printf("%-30s | %s\n", "go.sum (content hash)", "poetry.lock / pip freeze")
	fmt.Printf("%-30s | %s\n", "go get", "pip install / poetry add")
	fmt.Printf("%-30s | %s\n", "go mod tidy", "poetry lock --no-update")
	fmt.Printf("%-30s | %s\n", "go mod vendor", "pip install -t vendor/")

	// Demonstrate Go module behavior
	fmt.Println()
	fmt.Println("Module path determines import paths:")
	fmt.Println("  module: golan/packages-modules")
	fmt.Println("  import \"golan/packages-modules/calculator\"")
	fmt.Println("  import \"golan/packages-modules/models\"")
	fmt.Println("  import \"golan/packages-modules/utils\"")
}

// =============================================================================
// DEMO 7: Package Visibility Summary
// =============================================================================
// Quick reference for what's accessible from outside the package.

func demoVisibilitySummary() {
	fmt.Println("\n=== Demo 7: Package Visibility Rules ===")
	fmt.Println(`
  ┌───────────────────┬──────────────┬──────────────────────┐
  │ Go Declaration    │ Visible in   │ Visible from         │
  │                   │ Own Package  │ Other Packages       │
  ├───────────────────┼──────────────┼──────────────────────┤
  │ func Add()        │     ✅       │     ✅ (exported)     │
  │ func add()        │     ✅       │     ❌ (unexported)   │
  │ var Version       │     ✅       │     ✅ (exported)     │
  │ var author        │     ✅       │     ❌ (unexported)   │
  │ type User struct  │     ✅       │     ✅ (exported)     │
  │ type config struct│     ✅       │     ❌ (unexported)   │
  └───────────────────┴──────────────┴──────────────────────┘

  RULE: UPPERCASE = exported (public)
        lowercase = unexported (private)

  Applies to: functions, types, variables, constants, methods, fields`)
}

// =============================================================================
// MAIN — Entry Point
// =============================================================================

func main() {
	demoBasicImport()
	demoTypesFromPackage()
	demoMultiFilePackage()
	demoBlankImport()
	demoGoMod()
	demoVisibilitySummary()

	fmt.Println("\n" + "=" + strings.Repeat("=", 49))
	fmt.Println("PACKAGES & MODULES CHEAT SHEET")
	fmt.Println("=" + strings.Repeat("=", 49))
}

// NOTE: main.go uses strings.Repeat so we need "strings" imported
// oh wait — we DID import "fmt" and "os" but NOT strings!
