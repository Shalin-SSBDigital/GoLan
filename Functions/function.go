// ============================================================
//  FUNCTIONS IN GO — A Complete Guide with Python Comparisons
// ============================================================
//  Follows the 14-section teaching structure from Go-Teaching-Prompt.md
// ============================================================

package main

import "fmt"

// ============================================================
//  SECTION 1 — What is a Function?
// ============================================================
//  A function is a reusable block of code that:
//  - Takes INPUT (parameters)
//  - Performs some WORK
//  - Returns OUTPUT (return values)
//
//  Think of a function as a MINI-PROGRAM inside your program.
//  You give it ingredients (inputs), it does something, and
//  gives you back a result (output).
//
//  In Go, functions are the BUILDING BLOCKS of every program.
//  The simplest Go program is just:
//      func main() { }
//  A program with NO functions is like a house with NO rooms.

// ============================================================
//  SECTION 2 — Why Do We Need Functions?
// ============================================================
//  Without functions, every program would be one long, flat
//  list of instructions. Functions solve THREE big problems:
//
//  Problem 1: REPETITION
//     Without functions, you'd COPY-PASTE the same code
//     everywhere. Functions let you WRITE ONCE, USE MANY
//     TIMES. (DRY principle — Don't Repeat Yourself)
//
//     WRONG (no function):
//         a := 5 + 3
//         b := 10 + 3
//         c := 2 + 3   // repeated "+ 3" everywhere!
//
//     RIGHT (with function):
//         add(x, 3)    // one reusable function
//
//  Problem 2: COMPLEXITY
//     Functions let you BREAK a big problem into small,
//     digestible pieces. Each function does ONE thing well.
//     This is called "decomposition" or "modularity."
//
//     Example: A calculator program is NOT one giant block.
//     It's many small functions: add(), subtract(), etc.
//
//  Problem 3: TESTING
//     Small functions are EASY to test independently.
//     One big block of code is HARD to test.
//
//  Go specifically emphasizes functions because:
//  - Go is NOT object-oriented (no classes)
//  - Functions + types are the PRIMARY organizational tool
//  - Go programs are built from packages of functions

// ============================================================
//  SECTION 3 — Python Comparison
// ============================================================
//  Python also uses functions (def), but there are CRITICAL
//  differences in HOW Go and Python handle them:
//
//  ┌──────────────────────────┬──────────────────────────────-┬──────────────────────────────────┐
//  │ Concept                  │ Go                            │ Python                           │
//  ├──────────────────────────┼──────────────────────────────-┼──────────────────────────────────┤
//  │ Declaration keyword      │ func                          │ def                              │
//  │ Return type position     │ AFTER parameters              │ BEFORE colon (annotation)        │
//  │                          │ func(a int) int { ... }       │ def f(a: int) -> int: ...        │
//  │ Types required?          │ YES — compiler ENFORCES       │ NO — annotations are OPTIONAL    │
//  │ Type checking            │ AT COMPILE TIME               │ AT RUNTIME (or with mypy)        │
//  │ Return type              │ MANDATORY                     │ OPTIONAL (annotation)            │
//  │                          │ (must be declared)            │ (can omit completely)            │
//  │ Default parameter values │ NOT supported                 │ SUPPORTED (def f(a=5):)          │
//  │ Keyword arguments        │ NOT supported                 │ SUPPORTED (f(a=5))               │
//  │ *args / **kwargs         │ NOT supported                 │ SUPPORTED                        │
//  │ Multiple return values   │ TRUE multiple returns         │ Single tuple unpacked            │
//  │ Named return values      │ SUPPORTED (bare return)       │ NOT supported                    │
//  │ First-class citizens?    │ YES — functions are values    │ YES — functions are objects      │
//  │ Overloading              │ NOT supported                 │ NOT supported (natively)         │
//  │ (same name, diff params) │ (would be compile error)      │ (but *args + defaults simulate)  │
//  │ Can be passed as args?   │ YES (function types)          │ YES (functions are objects)      │
//  │ Can be returned?         │ YES (closures)                │ YES (closures)                   │
//  └──────────────────────────┴──────────────────────────────-┴──────────────────────────────────┘
//
//  The BIGGEST difference:
//      Go:  Everything is STATIC and EXPLICIT.
//           Types are REQUIRED and CHECKED at compile time.
//      Python: Everything is DYNAMIC and OPTIONAL.
//           Types are HINTS, not rules.
//
//  This means:
//  - Go catches TYPE ERRORS before you run the program
//  - Python catches TYPE ERRORS only when that line executes
//  - Go code is MORE VERBOSE but MORE PREDICTABLE
//  - Python code is SHORTER but you need MORE TESTS

// ============================================================
//  SECTION 4 — Syntax
// ============================================================
//  The syntax for declaring a Go function:
//
//      func functionName(param1 type1, param2 type2) returnType {
//          // body — statements go here
//          return value
//      }
//
//  Let's break down EVERY piece:
//
//  ┌──────────┬───────────────────────────────────────────────────┐
//  │ Keyword  │ Meaning                                           │
//  ├──────────┼───────────────────────────────────────────────────┤
//  │ func     │ Keyword to DECLARE a function (like Python's def) │
//  │ name     │ The function's name — must start with a letter     │
//  │ ()       │ Encloses the PARAMETER LIST                       │
//  │ param    │ Variable name for the input value                  │
//  │ type     │ The DATA TYPE of the parameter/return value        │
//  │ {}       │ Encloses the FUNCTION BODY                         │
//  │ return   │ Sends a value BACK to the caller                  │
//  └──────────┴───────────────────────────────────────────────────┘
//
//  ⚠️  TYPE COMES AFTER the variable name in Go!
//     Go:  func add(a int, b int) int
//     Python: def add(a: int, b: int) -> int:
//
//  This is the OPPOSITE of many languages (C, Java, Python).
//  Get used to this — it's one of Go's most distinctive features.
//
//  SHORTHAND: consecutive params of the SAME type:
//      func add(a, b int) int     ← "a int, b int" shortened
//
//  Python has NO equivalent shorthand for same-type params.
//
//  RETURN TYPE is AFTER the parameter list, BEFORE the body:
//      func name(params) returnType { body }
//      func add(a, b int) int { return a + b }
//                    ↑ here — after params, before {

// ============================================================
//  SECTION 5 — Simple Example
// ============================================================
//  The SMALLEST possible working Go function example:
//
//      func add(a int, b int) int {
//          return a + b
//      }
//
//  Line-by-line explanation:
//
//  Line 1:  func add(a int, b int) int {
//           ─────────────────────────────
//           func     → "I'm declaring a function"
//           add      → "Its name is 'add'"
//           (a int, b int) → "It takes TWO parameters:
//                             'a' is an int, 'b' is an int"
//           int      → "It returns ONE value of type int"
//           {        → "The function body starts here"
//
//  Line 2:  return a + b
//           ────────────
//           return   → "Send this value back to the caller"
//           a + b    → "Add a and b together"
//
//  Line 3:  }
//           → "The function body ends here"

//  CALLING this function:
//      result := add(3, 4)
//      fmt.Println(result)   // 7
//
//  The call syntax is IDENTICAL to Python:  name(arg1, arg2)
//
//  := (short variable declaration) creates 'result' and infers
//  its type (int) from the return type of add().
//  This is like Python's:  result = add(3, 4)
//  But in Python, result's type is determined at RUNTIME.
//  In Go, result's type is determined at COMPILE TIME.

// ============================================================
//  SECTION 6 — Python Equivalent
// ============================================================
//  Go:
//      func add(a int, b int) int {
//          return a + b
//      }
//      result := add(3, 4)
//      fmt.Println(result)   // 7
//
//  Python:
//      def add(a: int, b: int) -> int:
//          return a + b
//
//      result = add(3, 4)
//      print(result)   # 7
//
//  LINE-BY-LINE COMPARISON:
//
//  ┌──────────────────────────────────┬──────────────────────────────────────┐
//  │ Go                               │ Python                              │
//  ├──────────────────────────────────┼──────────────────────────────────────┤
//  │ func add(a int, b int) int {     │ def add(a: int, b: int) -> int:     │
//  │   return a + b                   │   return a + b                      │
//  │ }                                │                                     │
//  │                                  │                                     │
//  │ result := add(3, 4)              │ result = add(3, 4)                  │
//  │ fmt.Println(result)              │ print(result)                       │
//  └──────────────────────────────────┴──────────────────────────────────────┘
//
//  What's DIFFERENT:
//  1. Keyword: func vs def
//  2. Type position: Go puts type AFTER name, Python AFTER colon
//  3. Return type: Go puts it before {, Python uses ->
//  4. Short declaration: := vs =
//  5. Print: fmt.Println vs print
//
//  What's THE SAME:
//  1. Function call syntax: add(3, 4)
//  2. return keyword
//  3. Parameters in parentheses
//  4. Body indented inside { } or :
//  5. The actual logic: a + b

// ============================================================
//  SECTION 7 — Step-by-Step Execution
// ============================================================
//  When you call:  result := add(3, 4)
//
//  STEP 1: The Go runtime encounters add(3, 4)
//          ┌─────────────┐
//          │ Call stack:  │  main() is running
//          │ main()       │  add(3, 4) is about to execute
//          └─────────────┘
//
//  STEP 2: Go pushes a new FRAME onto the call stack
//          ┌─────────────┐
//          │ add(a=3,    │  ↑ New frame created
//          │  b=4)       │  Parameters are COPIED into a and b
//          ├─────────────┤  a = 3, b = 4 (new memory locations)
//          │ main()      │  ↓ main() is PAUSED, waiting for add()
//          └─────────────┘
//
//  Memory layout (simplified):
//      ┌─────┬─────┐
//      │ a=3 │ b=4 │   ← add()'s local variables
//      └─────┴─────┘
//
//  STEP 3: Go executes: return a + b
//          a + b = 3 + 4 = 7
//          The value 7 is ready to be sent back
//
//  STEP 4: The return value 7 is copied back to main()
//          ┌─────────────┐
//          │ main()      │  ← add()'s frame is POPPED (destroyed)
//          │ result = 7  │  a and b no longer exist
//          └─────────────┘
//
//  STEP 5: Go executes: fmt.Println(result)
//          Prints: 7
//
//  KEY INSIGHT: Parameters are COPIED into the function.
//  If you modify 'a' inside add(), the ORIGINAL value in
//  the caller is NOT affected. This is called "pass by value."
//  (Python also passes by value — but the "value" of an object
//  is a reference, so it looks like pass-by-reference for
//  mutable objects. Go is CONSISTENTLY pass-by-value.)

// ============================================================
//  SECTION 8 — Visual Explanation (no recursion for clarity)
// ============================================================
//  Think of functions as FACTORY WORKERS on an assembly line:
//
//                    ┌─────────────────────────────┐
//       raw material │      FUNCTION (machine)      │ finished product
//     ──────────────▶│  ┌─────────────────────┐    │──────────────▶
//      inputs (3, 4) │  │  return a + b        │    │  output (7)
//                    │  └─────────────────────┘    │
//                    └─────────────────────────────┘
//
//  The CALL STACK is like a STACK OF PAPER on a desk:
//
//     ┌──────────────┐  ← top of stack (currently executing)
//     │ add(3, 4)    │
//     │  local vars  │
//     ├──────────────┤
//     │ main()       │  ← waiting for add() to finish
//     │  result = ?  │
//     └──────────────┘
//
//  When add() finishes, its paper is REMOVED:
//
//     ┌──────────────┐  ← main() resumes with result = 7
//     │ main()       │
//     │  result = 7  │
//     └──────────────┘
//
//  FUNCTION SIGNATURE VISUALIZATION:
//
//     func  square(n  int)    int     { return n * n }
//     ────  ───────  ───     ───     ─────────────────
//      │       │      │       │              │
//      │       │      │       │              └── the WORK
//      │       │      │       └── OUTPUT type
//      │       │      └── INPUT type
//      │       └── parameter name
//      └── keyword

// ============================================================
//  SECTION 9 — Real-World Analogy
// ============================================================
//  A function is like a VENDING MACHINE.
//
//  ┌────────────────────────────────────────────────────────────┐
//  │                    VENDING MACHINE                         │
//  │                                                            │
//  │    You insert:     Machine does:          You get:         │
//  │    ┌──────┐       ┌──────────────┐       ┌──────┐         │
//  │    │ $1.50│       │ Check money  │       │ Soda │         │
//  │    │ B1   │──────▶│ Drop can     │──────▶│      │         │
//  │    └──────┘       │ Calculate chg│       └──────┘         │
//  │                    └──────────────┘                        │
//  └────────────────────────────────────────────────────────────┘
//
//  Mapping to Go:
//
//  ┌──────────────┬──────────────────────┬──────────────────────┐
//  │ Vending      │ Go Function          │ Example              │
//  │ Machine      │ Concept              │                      │
//  ├──────────────┼──────────────────────┼──────────────────────┤
//  │ Money + code │ Parameters (inputs)  │ func buy(            │
//  │              │                      │   money float64,     │
//  │              │                      │   code string        │
//  │              │                      │ )                    │
//  │ The machine  │ Function body (work) │ { checkMoney(money)  │
//  │ checks,      │                      │   dropCan(code)      │
//  │ drops,       │                      │   calcChange(...)    │
//  │ calculates   │                      │ }                    │
//  │ Soda + change│ Return values        │ (can, change)        │
//  │              │                      │                      │
//  │ Different    │ Different ARGUMENTS  │ buy(2.00, "A1")      │
//  │ money/codes  │ = same machine,      │ buy(1.50, "B2")      │
//  │              │ different result     │                      │
//  │ Buy multiple │ Call function MULTI- │ soda1 := buy(1.50)   │
//  │ times        │ PLE TIMES            │ soda2 := buy(2.00)   │
//  └──────────────┴──────────────────────┴──────────────────────┘
//
//  Key insight: The function IS the machine. It contains the
//  logic. You call it with different inputs and get different
//  outputs. You don't need to know HOW it works inside —
//  just WHAT it needs and WHAT it returns.

// ============================================================
//  SECTION 10 — Real-World Programming Use Cases
// ============================================================
//
//  1. HTTP HANDLERS (Web servers)
//     Every web request is handled by a function:
//         func handleRequest(w http.ResponseWriter, r *http.Request)
//     This is the CORE of Go web frameworks (net/http, Gin, Echo).
//     In Python:  def handle_request(request): ...
//
//  2. DATABASE OPERATIONS
//     Each database operation is a function:
//         func getUserByID(db *sql.DB, id int) (*User, error)
//         func createOrder(db *sql.DB, order Order) error
//     In Python:  def get_user_by_id(db, user_id): ...
//
//  3. VALIDATION
//     Input validation is typically one function per rule:
//         func validateEmail(email string) bool
//         func checkPasswordStrength(pw string) error
//
//  4. MIDDLEWARE (HTTP)
//     Functions that wrap other functions:
//         func loggingMiddleware(next http.Handler) http.Handler
//     This is a "higher-order function" — a function that takes
//     or returns another function. Python does this too with
//     decorators: @logging_middleware
//
//  5. CONCURRENT WORKERS (Go's specialty)
//     Functions run as GOROUTINES (lightweight threads):
//         go processJob(job)   // runs in the BACKGROUND
//     Python equivalent:  threading.Thread(target=process_job)
//     But Goroutines are MUCH lighter than threads.
//
//  6. COMMAND-LINE TOOLS
//     Each CLI subcommand maps to a function:
//         func runServe(cmd *cobra.Command, args []string)
//         func runMigrate(cmd *cobra.Command, args []string)
//
//  In ALL these cases, functions provide:
//  - ISOLATION (each function has its own scope)
//  - REUSABILITY (call the same function anywhere)
//  - TESTABILITY (test each function independently)
//  - READABILITY (function names document intent)

// ============================================================
//  SECTION 11 — Common Beginner Mistakes
// ============================================================
//
//  MISTAKE 1: Forgetting return type
//  ─────────────────────────────────
//  WRONG:
//      func add(a, b int) {       // ← NO return type!
//          return a + b
//      }
//  ERROR: "too many arguments to return"
//
//  RIGHT:
//      func add(a, b int) int {   // ← return type declared
//          return a + b
//      }
//
//  Why this happens: In Python, return types are never declared.
//  Python:  def add(a, b): return a + b   ← works fine
//  Go:      func add(a, b int) int { ... }  ← type REQUIRED
//
//  Memory tip: If you SEE a return keyword, ADD a return type!

//  MISTAKE 2: Returning the wrong type
//  ───────────────────────────────────
//  WRONG:
//      func add(a, b int) int {
//          return 3.14      // ← float64, but declared int!
//      }
//  ERROR: "cannot use 3.14 (type float64) as type int"
//
//  Python:  def add(a, b) -> int: return 3.14   ← NO error!
//  Python only checks at runtime (if at all).
//  Go catches this at COMPILE TIME.
//
//  The Go compiler is your FRIEND here. It PREVENTS bugs
//  before they happen. Python would silently return 3.14
//  even though the annotation says int.

//  MISTAKE 3: Ignoring returned values
//  ───────────────────────────────────
//  WRONG:
//      add(3, 4)       // ← return value is DISCARDED
//
//  Go allows this (no error), but it's usually a bug.
//  Python:  add(3, 4)    # same — Python allows this too
//
//  If a function returns a value, you should USE it:
//      result := add(3, 4)
//      fmt.Println(result)
//
//  Exception: Some functions return errors that MUST be checked.
//  The Go compiler actually WARNS you if you don't use a variable,
//  so the pattern is:
//      result, _ := someFunc()  // _ discards the error
//  But this is usually BAD practice — check errors!

//  MISTAKE 4: Confusing := and = for function results
//  ──────────────────────────────────────────────────
//  := DECLARES new variables. = ASSIGNS to existing ones.
//
//  WRONG:
//      var result int
//      result := add(3, 4)   // ← ERROR: no new variables on left
//  ERROR: "no new variables on left side of :="
//
//  RIGHT:
//      var result int
//      result = add(3, 4)    // ← assign to existing variable
//
//  RIGHT (shorter):
//      result := add(3, 4)   // ← declares AND assigns
//
//  Python has only = for both declaration and assignment:
//      result = add(3, 4)    # both declare and assign

//  MISTAKE 5: Parameter order confusion
//  ────────────────────────────────────
//  Go is POSITIONAL only (like C, Java):
//      func divide(a, b int) int { return a / b }
//      divide(10, 2)   // → 5
//      divide(2, 10)   // → 0 (integer division!)
//
//  In Python, you could use KEYWORD arguments:
//      divide(b=2, a=10)  # → 5  (order doesn't matter)
//
//  Go has NO keyword arguments. ORDER MATTERS. Always.

//  MISTAKE 6: Thinking functions are "special" objects
//  ────────────────────────────────────────────────────
//  In Python, functions are regular objects:
//      def add(a, b): return a + b
//      my_func = add      # assigning the function itself
//      my_func(3, 4)      # → 7
//
//  Go functions are ALSO first-class values:
//      var myFunc func(int, int) int
//      myFunc = add
//      myFunc(3, 4)       // → 7
//
//  But the SYNTAX for declaring function types is different:
//      Go:       func(int, int) int
//      Python:   Callable[[int, int], int]
//
//  Go's syntax reads like a function signature without the name:
//      func(int, int) int
//      ──┬── ──┬── ─┬─
//        │     │    └── Return type
//        │     └── Parameter types
//        └── It's a function

// ============================================================
//  SECTION 12 — Best Practices
// ============================================================
//  1. FUNCTION NAMES
//     - Use camelCase (not snake_case like Python)
//       Go:  getUserByID     Python:  get_user_by_id
//     - Exported functions start with CAPITAL letter:
//       GetUserByID (available outside the package)
//     - Unexported functions start with LOWERCASE:
//       getUserByID (private to the package)
//     - Name should DESCRIBE what the function DOES
//     - Verbs are best: saveUser, loadConfig, sendEmail
//
//  2. KEEP FUNCTIONS SMALL
//     - One function = one job
//     - If a function does TWO things, SPLIT it
//     - Ideal size: 5-20 lines (like a good Python function)
//     - If you need a comment to explain a block inside
//       the function, EXTRACT that block into its own function
//
//  3. EARLY RETURNS
//     - Check errors/edge cases FIRST and return early
//     - This reduces nesting and makes the "happy path" clear
//     - Go idiom: "if err != nil { return ... }"
//
//     Good:
//         func divide(a, b int) (int, error) {
//             if b == 0 {
//                 return 0, fmt.Errorf("cannot divide by zero")
//             }
//             return a / b, nil  // happy path — no nesting!
//         }
//
//  4. NO DEFAULT VALUES
//     Unlike Python, Go functions CANNOT have default parameter
//     values. If you need defaults, use a CONFIG STRUCT:
//
//         type Config struct {
//             Timeout  int  // default applied by caller
//             MaxRetry int  // default applied by caller
//         }
//         func connect(cfg Config) { ... }
//
//     Python:  def connect(timeout=30, max_retry=3): ...
//
//  5. RETURN EARLY, RETURN ERRORS
//     - Go's philosophy: ERRORS ARE VALUES, not exceptions
//     - Return errors explicitly, don't panic/throw
//     - Every function that CAN fail SHOULD return an error
//
//  6. FUNCTION COMPOSITION
//     - Small functions COMBINE to do complex work
//     - Like LEGO blocks: each is simple, together they build
//       anything
//
//     Example:
//         err := validate(input)
//         if err == nil {
//             user := createUser(input)
//             sendWelcomeEmail(user)
//         }
//
//  7. ZERO VALUE IS YOUR FRIEND
//     In Go, every type has a ZERO VALUE (default):
//         int     → 0
//         string  → ""
//         bool    → false
//         pointer → nil
//     If your function returns a zero value on error, that's OK:
//         func getCount() int { return 0 }  // 0 on empty is fine

// ============================================================
//  SECTION 13 — Summary Table
// ============================================================
//  ┌─────────────────────┬──────────────────────────┬─────────────────────────────┐
//  │ Feature             │ Go                       │ Python                      │
//  ├─────────────────────┼──────────────────────────┼─────────────────────────────┤
//  │ Keyword             │ func                     │ def                         │
//  │ Type position       │ func name(p type) ret    │ def name(p: type) -> ret:   │
//  │                     │ (type AFTER name)        │ (type AFTER colon)          │
//  │ Types required?     │ YES — compile error if    │ NO — optional annotations   │
//  │                     │ missing or wrong          │                             │
//  │ Return type         │ MANDATORY                 │ OPTIONAL                    │
//  │ Multiple returns    │ True multiple values      │ Single tuple (unpacked)     │
//  │ Named returns       │ Yes (with bare return)    │ No                          │
//  │ Default params      │ Not supported             │ Supported (param=value)     │
//  │ Keyword args        │ Not supported             │ Supported                   │
//  │ *args               │ Not supported             │ Supported (variadic)        │
//  │ Variadic syntax     │ ...type (produces slice)  │ *param (produces tuple)     │
//  │ Pass by value?      │ Yes (always)              │ Yes (for immutables)        │
//  │                     │                           │ No (for mutables — ref)    │
//  │ First-class?        │ Yes (function type)       │ Yes (Callable)              │
//  │ Naked return        │ Yes (with named returns)  │ No (syntax error)           │
//  │ Overloading         │ Not supported             │ Not supported (natively)    │
//  │ Body delimiters     │ { } (braces)              │ : + indentation             │
//  └─────────────────────┴──────────────────────────┴─────────────────────────────┘

// ============================================================
//  SECTION 14 — Key Takeaways
// ============================================================
//  ─────────────────────────────────────────────────────────────
//  1. func is Go's keyword for declaring functions (like def)
//  2. Types come AFTER the name in Go: func name(param type) ret
//     This is OPPOSITE to Python's type annotation position.
//  3. Types are MANDATORY and ENFORCED at compile time in Go.
//     Python types are OPTIONAL hints — NOT enforced.
//  4. Return type must be DECLARED in Go. Python has no such
//     requirement (annotations are optional).
//  5. Go has NO default parameters, no keyword arguments,
//     and no *args/**kwargs. Everything is positional.
//  6. Parameters are PASSED BY VALUE in Go. The function gets
//     a COPY. (Python also passes by value, but objects are
//     references, so mutability complicates things.)
//  7. The return type goes AFTER the parameter list, BEFORE
//     the opening brace: func name(params) ret { body }
//  8. Consecutive same-type params can be SHORTENED:
//     func add(a, b int) int  (instead of "a int, b int")
//  9. Functions are FIRST-CLASS CITIZENS in both languages.
//     They can be assigned to variables, passed as arguments,
//     and returned from other functions.
// 10. Capitalized function names are EXPORTED (public).
//     Lowercase names are UNEXPORTED (private to the package).
//  ─────────────────────────────────────────────────────────────

// ============================================================
//  CODE EXAMPLES
// ============================================================

// ┌────────────────────────────────────────────────────────────┐
// │  EXAMPLE 1: Basic function with single return value        │
// │  Shows the SIMPLEST possible function structure.           │
// ├────────────────────────────────────────────────────────────┤
// │  Python:  def square(x: int) -> int: return x * x         │
// │  Call:    result := square(5)     # result = 25           │
// └────────────────────────────────────────────────────────────┘
func square(x int) int {
	return x * x
}

// ┌────────────────────────────────────────────────────────────┐
// │  EXAMPLE 2: Consecutive same-type parameter shorthand     │
// │  Instead of:  func multiply(a int, b int) int             │
// │  You can write: a, b int  (one type for both)             │
// ├────────────────────────────────────────────────────────────┤
// │  Python has NO equivalent shorthand — you always           │
// │  write:  def multiply(a: int, b: int) -> int:             │
// └────────────────────────────────────────────────────────────┘
func multiply(a, b int) int {
	return a * b
}

// ┌────────────────────────────────────────────────────────────┐
// │  EXAMPLE 3: Function with no parameters and no return      │
// │  Some functions just DO something without needing input    │
// │  or producing output. These are common for:                │
// │  - Printing status messages                                │
// │  - Initializing resources                                  │
// │  - Sending notifications                                   │
// ├────────────────────────────────────────────────────────────┤
// │  Python:  def greet(): print("Hello!")                    │
// │  Python equivalent with type hints:                        │
// │    def greet() -> None: print("Hello!")                   │
// │  (In Python, -> None is optional. In Go, you just OMIT     │
// │  the return type when there's nothing to return.)          │
// └────────────────────────────────────────────────────────────┘
func greet() {
	fmt.Println("Hello from Go!")
	// No return statement needed — function returns when done
}

// ┌────────────────────────────────────────────────────────────┐
// │  EXAMPLE 4: Function with ONE parameter, no return         │
// │  Prints a message using the parameter value.               │
// ├────────────────────────────────────────────────────────────┤
// │  Python:  def say_hello(name: str) -> None:               │
// │               print(f"Hello, {name}!")                    │
// └────────────────────────────────────────────────────────────┘
func sayHello(name string) {
	fmt.Println("Hello,", name)
}

// ┌────────────────────────────────────────────────────────────┐
// │  EXAMPLE 5: Function returning a STRING                    │
// │  Returns are not limited to numbers — any type works.      │
// ├────────────────────────────────────────────────────────────┤
// │  Python:  def greet_user(name: str) -> str:               │
// │               return f"Welcome, {name}!"                  │
// └────────────────────────────────────────────────────────────┘
func greetUser(name string) string {
	return "Welcome, " + name + "!"
}

// ┌────────────────────────────────────────────────────────────┐
// │  EXAMPLE 6: Function with MULTIPLE PARAMETERS              │
// │  Different types in the same parameter list.               │
// ├────────────────────────────────────────────────────────────┤
// │  Python:  def describe_person(name: str, age: int) -> str:│
// │               return f"{name} is {age} years old"          │
// └────────────────────────────────────────────────────────────┘
func describePerson(name string, age int) string {
	return name + " is " + fmt.Sprint(age) + " years old"
	// fmt.Sprint converts any value to its string representation
}

// ┌────────────────────────────────────────────────────────────┐
// │  EXAMPLE 7: Named Return Values (Go-specific feature)      │
// │  Go lets you NAME the return values in the signature.      │
// │  This:                                                      │
// │    1. Documents what each returned value MEANS              │
// │    2. Creates ZERO-VALUED local variables automatically     │
// │    3. Allows NAKED (bare) return — just "return"            │
// │       returns the current values of named returns           │
// ├────────────────────────────────────────────────────────────┤
// │  Python has NO equivalent. You'd need to:                   │
// │    def divide(a, b):                                        │
// │        quotient = a // b                                    │
// │        remainder = a % b                                    │
// │        return quotient, remainder  # no "naked return"     │
// │                                                             │
// │  ⚠️  USE WITH CAUTION: Many Go style guides say to use      │
// │  naked returns only in SHORT functions (≤10 lines).         │
// │  In longer functions, explicit returns are clearer.         │
// └────────────────────────────────────────────────────────────┘
func divide(a, b int) (quotient int, remainder int) {
	// quotient and remainder are ALREADY declared with zero values
	// quotient = 0, remainder = 0

	// Assign to the named return variables:
	quotient = a / b
	remainder = a % b

	return // "naked return" — returns quotient and remainder
	// Equivalent to:  return quotient, remainder
}

// ┌────────────────────────────────────────────────────────────┐
// │  EXAMPLE 8: Function with bool return (predicate function) │
// │  A "predicate" is a function that returns true/false.      │
// │  Common for validation, filtering, and conditional checks. │
// ├────────────────────────────────────────────────────────────┤
// │  Python:  def is_even(n: int) -> bool: return n % 2 == 0  │
// └────────────────────────────────────────────────────────────┘
func isEven(n int) bool {
	return n%2 == 0
}

// ┌────────────────────────────────────────────────────────────┐
// │  EXAMPLE 9: Functions are first-class values               │
// │  You can ASSIGN a function to a variable, PASS it to       │
// │  another function, or RETURN it from a function.           │
// │                                                             │
// │  This is the foundation of closures, callbacks, and         │
// │  higher-order functions.                                    │
// ├────────────────────────────────────────────────────────────┤
// │  Python:  def double(x): return x * 2                      │
// │           f = double        # assign function to variable   │
// │           f(5)              # → 10                         │
// │           map(double, [1,2,3])  # pass as callback          │
// │                                                             │
// │  🆕 NEW to Go: the function TYPE syntax                     │
// │     func(int) int  = "a function that takes one int         │
// │                      and returns one int"                   │
// │     This reads like a function sig without the name:        │
// │     {func} {name} (params) {ret}                            │
// │     Omitted →   func          (int)  int                    │
// └────────────────────────────────────────────────────────────┘

// double is a regular function:
func double(x int) int {
	return x * 2
}

// apply takes a FUNCTION as a parameter:
//   fn has type "func(int) int" — any function that takes
//   one int and returns one int can be passed here.
//
// Python equivalent:
//   def apply(fn: Callable[[int], int], value: int) -> int:
//       return fn(value)
func apply(fn func(int) int, value int) int {
	return fn(value)
}

// ┌────────────────────────────────────────────────────────────┐
// │  EXAMPLE 10: Function returning a function (factory)       │
// │  Creates and returns a NEW function. Each call to           │
// │  makeMultiplier creates a FRESH, independent function.     │
// │                                                             │
// │  Python:  def make_multiplier(factor: int):                │
// │               def multiply(x: int) -> int:                  │
// │                   return x * factor                        │
// │               return multiply                              │
// └────────────────────────────────────────────────────────────┘
func makeMultiplier(factor int) func(int) int {
	// Returns an anonymous function that "remembers" factor
	return func(x int) int {
		return x * factor
	}
}

// ============================================================
//  EXECUTABLE DEMO
// ============================================================
func main4() {
	fmt.Println("══════════════════════════════════════")
	fmt.Println("  FUNCTIONS IN GO — Demo")
	fmt.Println("══════════════════════════════════════")

	// --- Example 1: Basic function ---
	fmt.Println("\n▶ Example 1: Basic function (square)")
	result := square(5)
	fmt.Println("  square(5) =", result)
	// Output: square(5) = 25

	// --- Example 2: Shorthand syntax ---
	fmt.Println("\n▶ Example 2: Shorthand parameters (multiply)")
	result2 := multiply(6, 7)
	fmt.Println("  multiply(6, 7) =", result2)
	// Output: multiply(6, 7) = 42

	// --- Example 3: No-return function ---
	fmt.Println("\n▶ Example 3: Void function (greet)")
	greet()
	// Output: Hello from Go!

	// --- Example 4: Function with parameter ---
	fmt.Println("\n▶ Example 4: Parameter only (sayHello)")
	sayHello("Alice")
	// Output: Hello, Alice

	// --- Example 5: String return ---
	fmt.Println("\n▶ Example 5: String return (greetUser)")
	msg := greetUser("Bob")
	fmt.Println(" ", msg)
	// Output: Welcome, Bob!

	// --- Example 6: Multiple params of different types ---
	fmt.Println("\n▶ Example 6: Multiple types (describePerson)")
	desc := describePerson("Charlie", 25)
	fmt.Println(" ", desc)
	// Output: Charlie is 25 years old

	// --- Example 7: Named return values ---
	fmt.Println("\n▶ Example 7: Named returns (divide)")
	q, r := divide(17, 5)
	fmt.Printf("  17 / 5 = quotient=%d, remainder=%d\n", q, r)
	// Output: 17 / 5 = quotient=3, remainder=2

	// --- Example 8: Predicate function ---
	fmt.Println("\n▶ Example 8: Predicate (isEven)")
	fmt.Println("  isEven(4) =", isEven(4))  // true
	fmt.Println("  isEven(7) =", isEven(7))  // false

	// --- Example 9: Function as value ---
	fmt.Println("\n▶ Example 9: Function as value")
	// Assign the function itself to a variable:
	var myFunc func(int) int
	myFunc = double
	fmt.Println("  myFunc(10) =", myFunc(10))
	// Output: myFunc(10) = 20

	// Pass a function to another function:
	result3 := apply(double, 7)
	fmt.Println("  apply(double, 7) =", result3)
	// Output: apply(double, 7) = 14

	// --- Example 10: Function factory ---
	fmt.Println("\n▶ Example 10: Function factory")
	doubleBy3 := makeMultiplier(3)
	doubleBy5 := makeMultiplier(5)

	fmt.Println("  doubleBy3(4) =", doubleBy3(4))   // 12
	fmt.Println("  doubleBy5(4) =", doubleBy5(4))   // 20
	fmt.Println("  doubleBy3(10) =", doubleBy3(10)) // 30
	// Each is an INDEPENDENT function with its own factor.

	// ============================================================
	//  SUMMARY VISUAL
	// ============================================================
	fmt.Println("\n══════════════════════════════════════")
	fmt.Println("  QUICK REFERENCE")
	fmt.Println("══════════════════════════════════════")
	fmt.Println(`
  DECLARATION:

  func name(param type, param type) returnType {
      return value
  }

  CALLING:

  result := name(arg1, arg2)

  SHORTHAND (same-type params):

  func add(a, b int) int { return a + b }

  NAMED RETURNS:

  func div(a, b int) (q, r int) {
      q = a / b
      r = a % b
      return  // naked return
  }

  FUNCTION TYPE (for variables/params):

  var fn func(int) int         = "fn is a function that
                                  takes int, returns int"
  func apply(fn func(int) int) = "apply takes a function
                                  as a parameter"`)
}
