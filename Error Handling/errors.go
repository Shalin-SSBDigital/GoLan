// ============================================================
//  ERROR HANDLING IN GO — with Python Comparisons
// ============================================================
//  Go does NOT have exceptions (try / except / finally).
//  Instead, Go uses explicit error VALUES returned from functions.
//
// ┌──────────────────────────┬─────────────────────────────────────┐
// │          Go              │              Python                 │
// ├──────────────────────────┼─────────────────────────────────────┤
// │ if err != nil { return } │ try: ... except Exception: ...     │
// ├──────────────────────────┼─────────────────────────────────────┤
// │ error (interface)        │ BaseException / Exception           │
// ├──────────────────────────┼─────────────────────────────────────┤
// │ errors.New("msg")        │ Exception("msg")                   │
// ├──────────────────────────┼─────────────────────────────────────┤
// │ fmt.Errorf("...%w", err) │ raise ... from err (chaining)      │
// ├──────────────────────────┼─────────────────────────────────────┤
// │ errors.Is(err, sentinel) │ isinstance(err, SomeException)     │
// ├──────────────────────────┼─────────────────────────────────────┤
// │ defer                    │ try: ... finally: ...              │
// ├──────────────────────────┼─────────────────────────────────────┤
// │ panic + recover          │ raise + except (but rarely used)   │
// └──────────────────────────┴─────────────────────────────────────┘
//
// KEY PHILOSOPHY:
//   Python: "Look before you leap" — use try/except anywhere
//   Go:     "Errors are values" — check and handle them explicitly
//   Go has NO try/except — every function that can fail returns an error
// ============================================================

package main

import (
	"errors"
	"fmt"
	"strings"
)

// =============================================================================
// 1. THE error INTERFACE
// =============================================================================
// In Go, `error` is a built-in INTERFACE (not a class, not an exception).
//
// Python:
//   All exceptions inherit from BaseException.
//   You RAISE them with `raise`.
//
// Go:
//   `error` is an interface with ONE method:
//
//     type error interface {
//         Error() string
//     }
//
//   Any type that has an `Error() string` method satisfies the error interface.
//   You RETURN the error instead of raising it.
//
// Python equivalent of the error interface:
//   class Error(ABC):
//       @abstractmethod
//       def __str__(self) -> str: ...
//
// KEY INSIGHT:
//   In Python, you say:   return {"error": "something went wrong"}
//   In Go, you say:       return 0, errors.New("something went wrong")
//   The error is a FIRST-CLASS VALUE — not a control flow mechanism.

// =============================================================================
// 2. CREATING ERRORS — errors.New
// =============================================================================
// Python:  raise Exception("file not found")
// Go:      return errors.New("file not found")
//
// errors.New returns an error value with the given message.
// It's the simplest way to create an error.

func divide(a, b int) (int, error) {
	if b == 0 {
		// Return zero value for int (0) and an error
		return 0, errors.New("division by zero")
	}
	// Return result and nil error (success)
	return a / b, nil
}

func demoBasicError() {
	fmt.Println("=== Basic Error Creation ===")

	// Successful case — err is nil
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	// Error case — err is non-nil
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
	// Output: Error: division by zero
	// The function returns immediately on error — no exception unwinding

	fmt.Println("\nPython equivalent:")
	fmt.Println("  def divide(a, b):")
	fmt.Println("      if b == 0:")
	fmt.Println("          raise Exception('division by zero')")
	fmt.Println("      return a / b")
	fmt.Println()
	fmt.Println("  try:")
	fmt.Println("      result = divide(10, 0)")
	fmt.Println("  except Exception as e:")
	fmt.Println("      print(f'Error: {e}')")
	fmt.Println()
	fmt.Println("KEY DIFFERENCE:")
	fmt.Println("  Python: raise  →  unwinds the stack automatically")
	fmt.Println("  Go:     return  →  caller MUST check err manually")
}

// =============================================================================
// 3. THE if err != nil PATTERN
// =============================================================================
// This is Go's most common pattern. You'll see it everywhere.
//
// Python developers often find this repetitive at first, but it has advantages:
//   1. Error handling is ALWAYS visible — no hidden exceptions
//   2. Every error path is explicit
//   3. No surprise stack unwinding

func readConfigFile(path string) (string, error) {
	// Simulating a file read
	if path == "" {
		return "", errors.New("path cannot be empty")
	}
	if path == "missing.txt" {
		return "", errors.New("file not found: " + path)
	}
	return "config: port=8080", nil
}

func processConfig(path string) error {
	// Each call that can fail has its own err check
	contents, err := readConfigFile(path)
	if err != nil {
		// Early return — like an inverted try/except
		return fmt.Errorf("config error: %w", err)
	}

	fmt.Println("Config loaded:", contents)
	return nil
}

func demoErrCheckPattern() {
	fmt.Println("\n=== The if err != nil Pattern ===")

	// Python equivalent:
	//   try:
	//       contents = read_config_file(path)
	//   except Exception as e:
	//       raise ConfigError(f"config error: {e}")
	//   print(f"Config loaded: {contents}")

	err := processConfig("missing.txt")
	if err != nil {
		fmt.Println("Failed:", err)
	}

	err = processConfig("valid.txt")
	if err != nil {
		fmt.Println("Failed:", err)
	}

	fmt.Println("\nKey difference:")
	fmt.Println("  Python puts errors in ONE place (except block)")
	fmt.Println("  Go spreads error checks throughout the function")
	fmt.Println("  Go wins: you can see WHICH call failed and handle differently")
}

// =============================================================================
// 4. fmt.Errorf — FORMATTED ERRORS
// =============================================================================
// Python:  raise ValueError(f"invalid age: {age}")
// Go:      return fmt.Errorf("invalid age: %d", age)
//
// fmt.Errorf works like fmt.Sprintf but returns an error.
// It's the most common way to create descriptive errors.

func validateAge(age int) error {
	if age < 0 {
		return fmt.Errorf("age cannot be negative: got %d", age)
	}
	if age > 150 {
		return fmt.Errorf("age seems unrealistic: got %d", age)
	}
	return nil // nil means "no error"
}

func demoFormattedError() {
	fmt.Println("\n=== Formatted Errors with fmt.Errorf ===")

	// Python: raise ValueError(f"age cannot be negative: got {-5}")
	err := validateAge(-5)
	if err != nil {
		fmt.Println("Validation error:", err)
	}

	err = validateAge(200)
	if err != nil {
		fmt.Println("Validation error:", err)
	}

	err = validateAge(25)
	if err != nil {
		fmt.Println("Unexpected error:", err)
	} else {
		fmt.Println("Age 25 is valid (no error)")
	}
}

// =============================================================================
// 5. SENTINEL ERRORS
// =============================================================================
// A sentinel error is a PREDEFINED error value used for comparison.
//
// Python's equivalent:
//   class FileNotFoundError(Exception): ...
//
//   raise FileNotFoundError("config.json")
//   except FileNotFoundError:
//       # handle specifically
//
// In Go, you define a package-level error variable and compare with ==.

// Sentinel errors (defined at package level, by convention prefixed with Err)
var (
	ErrNotFound   = errors.New("resource not found")
	ErrPermission = errors.New("permission denied")
	ErrTimeout    = errors.New("operation timed out")
)

func fetchData(id string) (string, error) {
	if id == "secret" {
		return "", ErrPermission
	}
	if id == "" {
		return "", ErrNotFound
	}
	return "data for " + id, nil
}

func demoSentinelErrors() {
	fmt.Println("\n=== Sentinel Errors ===")

	// Python:
	//   try:
	//       result = fetch_data("secret")
	//   except PermissionError:
	//       print("Access denied!")
	//       return

	_, err := fetchData("secret")
	if errors.Is(err, ErrPermission) {
		fmt.Println("Handling permission error: Access denied!")
	}

	_, err = fetchData("")
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Handling not-found error: Resource missing!")
	}

	// Compare errors directly
	fmt.Println("\nerrors.Is vs ==")
	fmt.Println("  errors.Is(err, ErrNotFound) =", errors.Is(err, ErrNotFound))
	fmt.Println("  err == ErrNotFound           =", err == ErrNotFound)
	fmt.Println("  (both work for unwrapped errors, but errors.Is is preferred)")

	fmt.Println("\nPython equivalent:")
	fmt.Println("  class NotFoundError(Exception): pass")
	fmt.Println("  class PermissionError(Exception): pass")
	fmt.Println("  raise NotFoundError()")
	fmt.Println("  except NotFoundError: ...")
}

// =============================================================================
// 6. CUSTOM ERROR TYPES
// =============================================================================
// Sometimes you need MORE than just a message — you need structured data.
// In Python, you define a custom exception class with extra fields.
// In Go, you define a struct that implements the error interface.

// Python:
//   class ValidationError(Exception):
//       def __init__(self, field, value, reason):
//           self.field = field
//           self.value = value
//           self.reason = reason
//           super().__init__(f"{field}: {reason} (got {value})")
//
//   raise ValidationError("age", -5, "cannot be negative")
//   except ValidationError as e:
//       print(e.field)  # access structured data

type ValidationError struct {
	Field  string
	Value  any    // any type
	Reason string
}

// This makes ValidationError implement the error interface
func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s (got %v)", e.Field, e.Reason, e.Value)
}

func validateUsername(name string) error {
	if name == "" {
		return &ValidationError{
			Field:  "username",
			Value:  name,
			Reason: "cannot be empty",
		}
	}
	if len(name) < 3 {
		return &ValidationError{
			Field:  "username",
			Value:  name,
			Reason: "too short (min 3 characters)",
		}
	}
	return nil
}

func demoCustomErrors() {
	fmt.Println("\n=== Custom Error Types ===")

	err := validateUsername("ab")
	if err != nil {
		// Type assertion to access structured fields
		// Python: except ValidationError as e:
		var valErr *ValidationError
		if errors.As(err, &valErr) {
			fmt.Println("Structured error info:")
			fmt.Println("  Field:", valErr.Field)
			fmt.Println("  Value:", valErr.Value)
			fmt.Println("  Reason:", valErr.Reason)
			fmt.Println("  Full:", err)
		}
	}

	err = validateUsername("")
	if err != nil {
		var valErr *ValidationError
		if errors.As(err, &valErr) {
			fmt.Println("Structured error info:")
			fmt.Println("  Field:", valErr.Field)
			fmt.Println("  Value:", valErr.Value)
			fmt.Println("  Reason:", valErr.Reason)
		}
	}

	fmt.Println("\nPython vs Go:")
	fmt.Println("  Python: Custom exception class inherits Exception")
	fmt.Println("  Go:     Any type with Error() string is an error")
	fmt.Println("  Go:     Use errors.As to extract custom type")
}

// =============================================================================
// 7. ERROR WRAPPING — %w AND errors.Is / errors.As
// =============================================================================
// Go 1.13+ supports error wrapping — creating error CHAINS.
//
// Python equivalent:
//   try:
//       read_file("config.json")
//   except FileNotFoundError as e:
//       raise ConfigError("failed to load config") from e
//       # e.__cause__ contains the original error
//
// In Go:
//   - Use %w in fmt.Errorf to wrap an error
//   - Use errors.Is to check if ANY error in the chain matches a sentinel
//   - Use errors.As to extract ANY error of a specific TYPE from the chain

func loadConfig(path string) error {
	// Simulates: file not found
	if path == "" {
		return ErrNotFound
	}
	return nil
}

// This wraps the original error with additional context
func initializeApp(path string) error {
	err := loadConfig(path)
	if err != nil {
		// %w wraps the error — creates a chain: "app init failed" → ErrNotFound
		return fmt.Errorf("app init failed: %w", err)
	}
	fmt.Println("App initialized with config:", path)
	return nil
}

func demoErrorWrapping() {
	fmt.Println("\n=== Error Wrapping ===")

	err := initializeApp("")

	fmt.Println("Full error:", err)
	// Output: app init failed: resource not found

	// errors.Is checks THROUGH the chain
	// returns true if ANY wrapped error matches ErrNotFound
	if errors.Is(err, ErrNotFound) {
		fmt.Println("errors.Is(err, ErrNotFound) => true")
		fmt.Println("  (checked through the wrapping chain)")
	}

	// Compare with == (does NOT unwrap)
	fmt.Println("err == ErrNotFound =>", err == ErrNotFound)
	fmt.Println("  (false because err is wrapped with more context)")

	fmt.Println("\nPython equivalent:")
	fmt.Println("  try:")
	fmt.Println("      load_config('')")
	fmt.Println("  except NotFoundError as e:")
	fmt.Println("      raise ConfigInitError('app init failed') from e")
	fmt.Println("  # e.__cause__ contains the original NotFoundError")

	fmt.Println("\nWhy wrapping matters:")
	fmt.Println("  - Gives CONTEXT to low-level errors")
	fmt.Println("  - Preserves the ORIGINAL error for matching")
	fmt.Println("  - Creates a searchable chain for debugging")
}

// =============================================================================
// 8. DEFER — Go's "finally"
// =============================================================================
// Python:  try: ... finally: cleanup()
// Go:      defer cleanup()
//
// defer schedules a function call to run JUST BEFORE the enclosing
// function returns — whether it succeeds OR panics.
//
// KEY PROPERTIES:
//   1. Executed when the function exits (normal return OR panic)
//   2. Arguments are EVALUATED IMMEDIATELY (not when deferred runs)
//   3. Multiple defers run in LIFO order (stack: last-in-first-out)
//   4. Deferred functions CAN modify named return values

func demoDeferBasic() {
	fmt.Println("\n=== Defer (Go's finally) ===")

	// defer + anonymous function — runs when this function ends
	defer fmt.Println("  3. defer: cleanup complete")
	defer fmt.Println("  2. defer: closing resources...")

	fmt.Println("  1. Function body executing")

	// Output order:
	//   1. Function body executing
	//   2. defer: closing resources...  (LIFO - runs second)
	//   3. defer: cleanup complete       (LIFO - runs first)

	fmt.Println("\nDefer stack (LIFO order):")
	fmt.Println("  defer statements are PUSHED onto a stack")
	fmt.Println("  They POP in reverse order on function exit")
	fmt.Println("  Like Python's contextlib.ExitStack but automatic")
}

// Practical example: simulating file operations
func readFileWithDefer(path string) error {
	fmt.Println("\n=== Defer in Practice ===")

	// Simulate opening a file
	fmt.Println("  1. Opening file:", path)
	defer fmt.Println("  3. defer: Closing file", path)

	// Simulate error during processing
	if path == "" {
		return errors.New("empty path")
	}

	// Simulate processing
	fmt.Println("  2. Processing file:", path)
	return nil
	// defer runs HERE — "Closing file" is printed after "Processing"
}

func demoDeferClosings() {
	readFileWithDefer("data.txt")
	readFileWithDefer("")  // Even on ERROR, defer runs
	fmt.Println("  (defer runs even when the function returns an error!)")

	fmt.Println("\nPython equivalent:")
	fmt.Println("  def read_file(path):")
	fmt.Println("      f = open(path)")
	fmt.Println("      try:")
	fmt.Println("          return process(f)")
	fmt.Println("      finally:")
	fmt.Println("          f.close()  # ← always runs")
	fmt.Println()
	fmt.Println("  In Go: defer f.Close()  # ← always runs")
	fmt.Println("  Go's defer is cleaner — the close is RIGHT NEXT to the open")
}

// Deferred functions can MODIFY named return values
func demoDeferNamedReturn() {
	fmt.Println("\n=== Defer Modifying Named Returns ===")

	count := func() (total int) {
		// This defer runs AFTER the return statement but BEFORE the caller gets the value
		defer func() {
			total += 10  // modifies the return value!
			fmt.Println("  defer: adjusted total to", total)
		}()

		total = 5
		fmt.Println("  function: setting total to", total)
		return
		// The return sets total=5, then defer adds 10 → total=15
	}()

	fmt.Println("  caller got:", count)
	// Output: caller got: 15

	fmt.Println("\nPython has NO equivalent for this.")
	fmt.Println("A Python finally block cannot modify the return value.")
}

// =============================================================================
// 9. PANIC — Go's "Exception"
// =============================================================================
// Python:  raise RuntimeError("something terrible happened")
// Go:      panic("something terrible happened")
//
// panic is Go's version of raising an exception:
//   - It IMMEDIATELY stops normal execution
//   - All deferred functions still run
//   - The program crashes unless recovered
//
// WHEN TO USE PANIC:
//   - Programming errors (index out of bounds, nil dereference)
//   - Impossible states (things that should NEVER happen)
//   - NOT for normal error handling (use error values for that!)
//
// In Go, you should PANIC for things like:
//   - Python: assert x > 0, "x must be positive"
//   - Go:     if x <= 0 { panic("x must be positive") }
//
// But the Go way is to return errors for EXPECTED failures
// and use panic only for UNEXPECTED failures.

func demoPanic() {
	fmt.Println("\n=== Panic ===")

	defer fmt.Println("  defer: this ALWAYS runs")

	fmt.Println("  About to panic...")

	// Uncomment to see panic in action:
	// dangerousOperation := func() {
	// 	panic("something went terribly wrong!")
	// }
	// dangerousOperation()

	fmt.Println("  (panic is commented out — would crash the program)")

	// Real-world panics happen automatically:
	fmt.Println("\nGo panics automatically for:")
	fmt.Println("  - Index out of bounds:  s[10] on a 3-element slice")
	fmt.Println("  - Nil pointer dereference: var p *int; *p = 5")
	fmt.Println("  - Type assertion failure (without comma-ok): x.(string)")
	fmt.Println("  - Map concurrent read/write")

	fmt.Println("\nPython equivalent:")
	fmt.Println("  raise RuntimeError('something went terribly wrong')")
	fmt.Println()
	fmt.Println("But in Go, panic is RARE — most 'errors' are returned values.")
	fmt.Println("Python raises exceptions for EVERYTHING (even FileNotFound).")
}

// =============================================================================
// 10. RECOVER — CATCHING PANICS
// =============================================================================
// Python:  try: ... except Exception: ...
// Go:      defer func() { recover() }()
//
// recover is a built-in function that STOPS a panic and returns the
// panic value. It ONLY works inside a deferred function.
//
// IMPORTANT:
//   - recover() is useless outside a deferred function
//   - It returns nil if there's no active panic
//   - It's NOT for regular error handling — only for truly unexpected panics

func safeOperation(shouldPanic bool) (result string, err error) {
	// Recover must be in a defer — this catches ANY panic in this function
	defer func() {
		if r := recover(); r != nil {
			// Convert the panic to an error (like catching an exception)
			err = fmt.Errorf("panic recovered: %v", r)
			result = ""  // reset the result
			fmt.Println("  recover: caught panic, converted to error")
		}
	}()

	if shouldPanic {
		panic("internal state corrupted!")
	}

	return "operation successful", nil
}

func demoRecover() {
	fmt.Println("\n=== Recover (Catching Panics) ===")

	// Safe call — no panic
	result, err := safeOperation(false)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Call that panics — caught by recover, returned as error
	result, err = safeOperation(true)
	if err != nil {
		fmt.Println("Caught as error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	fmt.Println("Program continues normally after recovered panic!")

	fmt.Println("\nPython equivalent:")
	fmt.Println("  def safe_operation(should_panic):")
	fmt.Println("      try:")
	fmt.Println("          if should_panic:")
	fmt.Println("              raise RuntimeError('internal state corrupted!')")
	fmt.Println("          return 'operation successful'")
	fmt.Println("      except RuntimeError as e:")
	fmt.Println("          return None, str(e)")

	fmt.Println("\nWhen to use recover:")
	fmt.Println("  1. HTTP server middleware — catch panics, return 500")
	fmt.Println("  2. Long-running goroutines — prevent one crash from taking down the app")
	fmt.Println("  3. Cleanup on unexpected failure")
	fmt.Println("  DON'T use it as a substitute for checking errors!")
}

// =============================================================================
// 11. REAL-WORLD ERROR HANDLING PATTERN
// =============================================================================
// This shows how errors are handled in production Go code.

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidInput    = errors.New("invalid input")
	ErrDatabaseTimeout = errors.New("database timeout")
)

type DatabaseError struct {
	Query string
	Err   error
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("database error on query %q: %v", e.Query, e.Err)
}

func (e *DatabaseError) Unwrap() error {
	return e.Err
}

// Simulated user lookup
func findUserByID(id int) (string, error) {
	if id <= 0 {
		return "", fmt.Errorf("%w: id must be positive (got %d)", ErrInvalidInput, id)
	}
	if id == 404 {
		// Wrap the sentinel with a custom error
		return "", &DatabaseError{
			Query: fmt.Sprintf("SELECT * FROM users WHERE id=%d", id),
			Err:   ErrUserNotFound,
		}
	}
	if id == 500 {
		return "", &DatabaseError{
			Query: fmt.Sprintf("SELECT * FROM users WHERE id=%d", id),
			Err:   ErrDatabaseTimeout,
		}
	}
	return fmt.Sprintf("User_%d", id), nil
}

// High-level handler — demonstrates layered error handling
func getUserHandler(userID int) {
	fmt.Println("\n=== Real-World Error Handling ===")

	name, err := findUserByID(userID)
	if err != nil {
		// Check specific sentinel errors (works through wrapping!)
		switch {
		case errors.Is(err, ErrInvalidInput):
			fmt.Printf("[400 Bad Request] %v\n", err)
		case errors.Is(err, ErrUserNotFound):
			fmt.Printf("[404 Not Found] %v\n", err)
		case errors.Is(err, ErrDatabaseTimeout):
			fmt.Printf("[503 Service Unavailable] %v\n", err)
		default:
			fmt.Printf("[500 Internal Error] %v\n", err)
		}

		// Extract structured error info
		var dbErr *DatabaseError
		if errors.As(err, &dbErr) {
			fmt.Printf("  (database query was: %s)\n", dbErr.Query)
		}
		return
	}
	fmt.Printf("[200 OK] Found user: %s\n", name)
}

func demoRealWorld() {
	getUserHandler(1)    // success
	getUserHandler(0)    // invalid input
	getUserHandler(404)  // not found
	getUserHandler(500)  // database timeout
}

// =============================================================================
// 12. COMPARISON TABLE
// =============================================================================
//
// ┌──────────────────────────────┬──────────────────────────────────┬────────────────────────────────────┐
// │          Concept             │            Go                    │              Python                 │
// ├──────────────────────────────┼──────────────────────────────────┼────────────────────────────────────┤
// │ Error type                   │ error interface                  │ BaseException / Exception           │
// ├──────────────────────────────┼──────────────────────────────────┼────────────────────────────────────┤
// │ Create error                 │ errors.New("msg")               │ Exception("msg")                   │
// ├──────────────────────────────┼──────────────────────────────────┼────────────────────────────────────┤
// │ Formatted error              │ fmt.Errorf("x=%d", x)           │ ValueError(f"x={x}")               │
// ├──────────────────────────────┼──────────────────────────────────┼────────────────────────────────────┤
// │ Custom error with data       │ struct implementing error       │ Custom class inheriting Exception   │
// ├──────────────────────────────┼──────────────────────────────────┼────────────────────────────────────┤
// │ Sentinel error               │ var ErrX = errors.New(...)      │ class XError(Exception): pass      │
// ├──────────────────────────────┼──────────────────────────────────┼────────────────────────────────────┤
// │ Check error type             │ errors.Is(err, ErrX)            │ isinstance(err, XError)            │
// ├──────────────────────────────┼──────────────────────────────────┼────────────────────────────────────┤
// │ Extract structured error     │ errors.As(err, &target)         │ except XError as e:                │
// ├──────────────────────────────┼──────────────────────────────────┼────────────────────────────────────┤
// │ Error chaining               │ fmt.Errorf("ctx: %w", err)      │ raise ... from err                 │
// ├──────────────────────────────┼──────────────────────────────────┼────────────────────────────────────┤
// │ Always-run cleanup           │ defer f.Close()                 │ try: ... finally: f.close()        │
// ├──────────────────────────────┼──────────────────────────────────┼────────────────────────────────────┤
// │ Crash program                │ panic("msg")                    │ raise RuntimeError("msg")          │
// ├──────────────────────────────┼──────────────────────────────────┼────────────────────────────────────┤
// │ Catch crash                  │ recover() in defer              │ except:                            │
// ├──────────────────────────────┼──────────────────────────────────┼────────────────────────────────────┤
// │ Default behavior             │ Return error VALUE              │ Raise EXCEPTION                    │
// ├──────────────────────────────┼──────────────────────────────────┼────────────────────────────────────┤
// │ Stack trace                  │ Manual (fmt.Printf("%+v", err)) │ Automatic (traceback)              │
// └──────────────────────────────┴──────────────────────────────────┴────────────────────────────────────┘

// =============================================================================
// MAIN — Run All Demos
// =============================================================================

func main() {
	demoBasicError()
	demoErrCheckPattern()
	demoFormattedError()
	demoSentinelErrors()
	demoCustomErrors()
	demoErrorWrapping()
	demoDeferBasic()
	demoDeferClosings()
	demoDeferNamedReturn()
	demoPanic()
	demoRecover()
	demoRealWorld()

	fmt.Println("\n" + "=" + strings.Repeat("=", 49))
	fmt.Println("ERROR HANDLING CHEAT SHEET")
	fmt.Println("=" + strings.Repeat("=", 49))
}

// =============================================================================
// CHEAT SHEET
// =============================================================================
//
// // 1. Basic error
// if err != nil { return err }
//
// // 2. Create error
// errors.New("message")
// fmt.Errorf("formatted: %d", n)
//
// // 3. Sentinel error
// var ErrNotFound = errors.New("not found")
// if errors.Is(err, ErrNotFound) { ... }
//
// // 4. Custom error type
// type MyError struct { Field string }
// func (e *MyError) Error() string { return e.Field }
// var target *MyError
// if errors.As(err, &target) { ... }
//
// // 5. Wrapping
// fmt.Errorf("context: %w", err)  // wrap with %w
// errors.Is(err, ErrX)            // checks through chain
// errors.As(err, &target)         // extracts through chain
//
// // 6. Defer
// defer file.Close()              // runs on function exit
// defer func() { ... }()          // anonymous defer
//
// // 7. Panic + Recover
// panic("unexpected")             // crash (rare)
// defer func() {                  // catch (rare)
//     if r := recover(); r != nil { ... }
// }()
