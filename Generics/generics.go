// ============================================================
//  GENERICS IN GO — with Python Comparisons
// ============================================================
//  Generics let you write functions, types, and data structures
//  that work with ANY type — without writing separate code for
//  each type.
//
// ┌──────────────────────┬──────────────────────────────────────────────┐
// │        Go            │                 Python                       │
// ├──────────────────────┼──────────────────────────────────────────────┤
// │ func Print[T any]    │ from typing import TypeVar                   │
// │     (v T) { ... }    │ T = TypeVar('T')                            │
// │                      │ def print_value(v: T) -> None:               │
// │                      │     print(v)                                 │
// ├──────────────────────┼──────────────────────────────────────────────┤
// │ Declared INLINE      │ Declared as SEPARATE variable                │
// │ [T any] after name   │ T = TypeVar('T') before function             │
// ├──────────────────────┼──────────────────────────────────────────────┤
// │ ENFORCED at compile  │ NOT enforced — just hints for mypy           │
// │ time. Wrong type =   │ Python still runs with wrong types.          │
// │ COMPILE ERROR.       │                                              │
// ├──────────────────────┼──────────────────────────────────────────────┤
// │ Zero-cost: compiler  │ No runtime checks — Python interpreter          │
// │ generates separate   │ checks types at runtime (or not at all)      │
// │ code per type.       │                                              │
// └──────────────────────┴──────────────────────────────────────────────┘
// ============================================================

package main

import "fmt"

// ============================================================
//  TYPE CONSTRAINT — Inline Definition
// ============================================================
// Since Go has no built-in "Ordered" constraint in the standard
// library (it lives in golang.org/x/exp/constraints), we define
// our own here for self-contained code.
//
// This is called a TYPE SET constraint (Go 1.18+).
// The | symbol means "any of these types."
// The ~ symbol means "any type whose UNDERLYING type matches."
//   ~int matches int, and also type MyInt int
//
// Python equivalent:
//   There is NO Python equivalent. Python is dynamically typed,
//   so it doesn't need — and can't express — type constraints.
//
//   In Python, you'd just use the < operator and hope:
//       def min_value(a, b):
//           return a if a < b else b
//   If you pass types that don't support <, you get TypeError
//   at RUNTIME. Go catches this at COMPILE TIME.
// ============================================================

// Ordered is a constraint that matches any ordered type.
// These types support the <, <=, >, >= operators.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// ============================================================
// 1. BASIC GENERIC FUNCTION
// ============================================================
// [T any] declares a type parameter named T.
//  - T can be ANY type (int, string, float64, custom struct, etc.)
//  - any is an alias for interface{} — it means "accept anything"
//
// Python equivalent:
//   T = TypeVar('T')
//   def print_value(v: T) -> None:
//       print(v)
//
// KEY DIFFERENCES:
//   - Go: T is declared INLINE: [T any] after the function name
//   - Python: T is a SEPARATE variable declared before the function
//   - Go: compiler ENFORCES type safety at compile time
//   - Python: TypeVar is just a hint for mypy; Python ignores it
// ============================================================
func Print[T any](value T) {
	fmt.Println(value)
}

// ============================================================
// 2. GENERIC FUNCTION WITH TYPE RELATIONSHIP
// ============================================================
// Identity takes a value of any type T and returns the SAME type T.
// This is the classic use case for generics: the input type and
// output type are RELATED (they're the same).
//
// Python:
//   def identity(v: T) -> T:
//       return v
//
// Without generics, you'd have to use any (interface{}):
//   func Identity(v any) any { return v }
// But then you lose type safety — the caller gets back any,
// not the original type. With generics, you get back EXACTLY T.
// ============================================================
func Identity[T any](value T) T {
	return value
}

// ============================================================
// 3. MULTIPLE TYPE PARAMETERS
// ============================================================
// [T, U any] declares TWO type parameters.
// T and U can be DIFFERENT types or the SAME type.
//
// This is like Go's equivalent of:
//   T = TypeVar('T')
//   U = TypeVar('U')
//   def pair(a: T, b: U) -> tuple[T, U]:
//       return (a, b)
//
// ⚠️ Python returns a TUPLE. Go returns TWO separate values
//    (like multiple return values, not a tuple).
// ============================================================
func Pair[T, U any](a T, b U) (T, U) {
	return a, b
}

// ============================================================
// 4. GENERIC STRUCT
// ============================================================
// Go lets you define generic STRUCTS (not just functions).
// A generic struct can hold values of any type.
//
// Python has no direct equivalent because Python has no
// compile-time generics. The closest is:
//   from dataclasses import dataclass
//   from typing import Generic, TypeVar
//
//   T = TypeVar('T')
//   @dataclass
//   class Box(Generic[T]):
//       value: T
//
// But Python's Generic[T] is just a TYPE HINT.
// Go's Box[T any] is ENFORCED — a Box[int] can ONLY hold ints.
// ============================================================

// Box is a generic struct that holds one value of any type.
type Box[T any] struct {
	Value T
}

// ============================================================
// 5. GENERIC METHOD
// ============================================================
// Methods on generic types work like you'd expect.
// The receiver uses the SAME type parameter.
//
// Python equivalent:
//   class Box(Generic[T]):
//       def __init__(self, value: T):
//           self.value = value
//       def get(self) -> T:
//           return self.value
//
// ⚠️ KEY RULE: A generic type's methods must use the SAME type
//    parameter declared on the type. We write *Box[T], not *Box.
// ============================================================

// Get returns the value inside the box.
func (b Box[T]) Get() T {
	return b.Value
}

// Set updates the value inside the box.
func (b *Box[T]) Set(value T) {
	b.Value = value
}

// ============================================================
// 6. CONSTRAINTS — Restricting What T Can Be
// ============================================================
// [T any] means "T can be ANYTHING." But sometimes you need T
// to support specific operations (like <, >, +, etc.).
//
// For that, you use a CONSTRAINT — an interface that T must satisfy.
//
// Python equivalent:
//   from typing import Protocol
//   class Comparable(Protocol):
//       def __lt__(self, other) -> bool: ...
//
//   def min_value[T: Comparable](a: T, b: T) -> T:
//       return a if a < b else b
//
// Go's constraints are INTERFACES. The type must have ALL methods
// the interface requires. For numeric types, Go provides the
// "golang.org/x/exp/constraints" package.
// ============================================================

// Min returns the smaller of two values of the same orderable type.
// constraints.Ordered covers types that support <, <=, >, >=:
//   - All integer types (int, int8, ..., uint, uint8, ...)
//   - All float types (float32, float64)
//   - string (lexicographic comparison)
//
// Python:
//   from typing import TypeVar
//   Comparable = TypeVar('Comparable', bound=...)
//   def min_value(a: Comparable, b: Comparable) -> Comparable:
//       return a if a < b else b
//
// KEY DIFFERENCE: Python's bound is a SUGGESTION. Go's constraint
// is ENFORCED at compile time.
func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// ============================================================
// 7. THE ZERO VALUE PROBLEM
// ============================================================
// Since T can be any type, what is the "zero" value of T?
// You CANNOT return nil because T might be int (nil doesn't work for ints).
//
// Solution: declare a variable of type T — it starts at its zero value.
//   var zero T  →  0 for int, "" for string, false for bool, nil for pointers
//
// Python:
//   def zero[T]() -> T:
//       # Python doesn't have a concept of "zero value"
//       # This is fundamentally a Go-specific pattern
//       ...
// ============================================================
func Zero[T any]() T {
	var zero T // ← T automatically gets its zero value
	return zero
}

// ============================================================
// 8. GENERIC STACK — A Real Data Structure
// ============================================================
// A stack is a LIFO (Last In, First Out) data structure.
// With generics, we write ONE stack that works with any type.
//
// Python:
//   from typing import Generic, TypeVar
//   T = TypeVar('T')
//
//   class Stack(Generic[T]):
//       def __init__(self):
//           self.items: list[T] = []
//
//       def push(self, item: T) -> None:
//           self.items.append(item)
//
//       def pop(self) -> T | None:
//           if not self.items:
//               return None
//           return self.items.pop()
//
// KEY DIFFERENCE: Python can return None for empty pop.
// Go returns the zero value of T + a bool (comma-ok pattern).
// ============================================================

// Stack is a generic LIFO stack that works with any type.
type Stack[T any] struct {
	items []T
}

// Push adds an item to the top of the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item.
// Returns the zero value + false if the stack is empty.
// This is the COMMA-OK pattern, same as map access.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T         // ← zero value for T
		return zero, false // false = stack is empty
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true // true = success
}

// Peek returns the top item WITHOUT removing it.
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty returns true if the stack has no items.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack.
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// ============================================================
// 9. GENERIC MAP FUNCTION (like Python's map())
// ============================================================
// [T, U any] declares TWO type parameters.
// T is the input type, U is the output type.
//
// This is like Python's built-in map():
//   map(lambda x: x * 2, [1, 2, 3])  # → [2, 4, 6]
//
// But Python's map() returns a lazy iterator. Go's version
// is EAGER — it creates the new slice immediately.
//
// KEY INSIGHT: Multiple type params let you TRANSFORM types:
//   Map[int, string]([1, 2, 3], strconv.Itoa)  → ["1", "2", "3"]
// ============================================================
func Map[T, U any](items []T, fn func(T) U) []U {
	result := make([]U, len(items))
	for i, item := range items {
		result[i] = fn(item)
	}
	return result
}

// ============================================================
// 10. GENERIC API RESPONSE (Real-World Pattern)
// ============================================================
// This is an EXTREMELY common pattern in Go backends.
// A single generic type represents ANY API response,
// keeping the response structure consistent while allowing
// different data types in the "data" field.
//
// Python equivalent (Pydantic / FastAPI):
//   from pydantic import BaseModel
//   from typing import Generic, TypeVar, Optional
//
//   T = TypeVar('T')
//   class APIResponse(BaseModel, Generic[T]):
//       success: bool
//       data: Optional[T] = None
//       error: Optional[str] = None
//
//   class User(BaseModel):
//       id: int
//       name: str
//
//   def get_user() -> APIResponse[User]: ...
// ============================================================

// APIResponse wraps any data type in a standard API response.
type APIResponse[T any] struct {
	Success bool   `json:"success"`
	Data    T      `json:"data"`
	Error   string `json:"error,omitempty"` // omitempty = skip if empty
}

// NewSuccess creates a success response with data.
func NewSuccess[T any](data T) APIResponse[T] {
	return APIResponse[T]{
		Success: true,
		Data:    data,
	}
}

// NewError creates an error response.
// Note: T is still required even though Data is zero-valued.
// The caller specifies what type the response "would have" held.
func NewError[T any](errMsg string) APIResponse[T] {
	var zero T
	return APIResponse[T]{
		Success: false,
		Data:    zero,
		Error:   errMsg,
	}
}

// ============================================================
// GENERICS WITH INTERFACES — Types and constraint
// ============================================================
// These are defined at PACKAGE level because Go doesn't allow
// defining functions or types inside other functions.
//
// Python equivalent:
//   from typing import Protocol
//
//   class Named(Protocol):
//       def name(self) -> str: ...
//
//   class Person:
//       def __init__(self, first_name: str):
//           self.first_name = first_name
//       def name(self) -> str:
//           return self.first_name
//
//   def greet(entity: Named) -> str:
//       return f"Hello, {entity.name()}!"
// ============================================================

// Named is an interface constraint: any type with a Name() method.
type Named interface {
	Name() string
}

// Greet is a generic function constrained to Named types.
// T must satisfy Named (i.e., have a Name() string method).
func Greet[T Named](entity T) string {
	return "Hello, " + entity.Name() + "!"
}

// Person implements Named with a value receiver.
type Person struct {
	firstName string
}

func (p Person) Name() string {
	return p.firstName
}

// Pet also implements Named — different struct, same interface.
type Pet struct {
	nickname string
}

func (p Pet) Name() string {
	return p.nickname
}

// ============================================================
// DEMO — Putting it all together
// ============================================================

func main() {
	// ============================================================
	// 1. Basic Generic Function
	// ============================================================
	fmt.Println("=== 1. Basic Generic Function ===")
	Print[int](42)       // Explicit: T = int
	Print[string]("hi")  // Explicit: T = string
	Print[float64](3.14) // Explicit: T = float64

	// Type INFERENCE — Go infers T from the argument:
	Print(42)      // Go infers T = int automatically
	Print("hello") // Go infers T = string automatically
	fmt.Println()

	// ============================================================
	// 2. Identity (Type Relationship)
	// ============================================================
	fmt.Println("=== 2. Identity (Type Relationship) ===")
	// x is of type int (not 'any'!)
	x := Identity[int](42)
	fmt.Printf("Identity returned: %d (type: %T)\n", x, x)

	// With inference:
	y := Identity("still a string")
	fmt.Printf("Identity returned: %s (type: %T)\n", y, y)
	fmt.Println()

	// ============================================================
	// 3. Multiple Type Params
	// ============================================================
	fmt.Println("=== 3. Multiple Type Params ===")
	a, b := Pair[int, string](1, "hello")
	fmt.Printf("Pair: a=%d (int), b=%s (string)\n", a, b)

	// T and U can be the SAME type:
	c, d := Pair(3.14, 2.71) // T = float64, U = float64
	fmt.Printf("Pair: c=%.2f, d=%.2f (both float64)\n", c, d)
	fmt.Println()

	// ============================================================
	// 4. Generic Struct
	// ============================================================
	fmt.Println("=== 4. Generic Struct ===")
	intBox := Box[int]{Value: 42}
	strBox := Box[string]{Value: "hello, generics!"}

	fmt.Printf("intBox: %+v\n", intBox)
	fmt.Printf("strBox: %+v\n", strBox)

	// Type safety: this would NOT compile:
	// intBox.Value = "not an int"  // COMPILE ERROR!

	// Calling methods:
	fmt.Printf("intBox.Get() = %d\n", intBox.Get())
	intBox.Set(100)
	fmt.Printf("intBox after Set(100): %+v\n", intBox)
	fmt.Println()

	// ============================================================
	// 5. Constraints (Min)
	// ============================================================
	fmt.Println("=== 5. Constraints (Min) ===")
	fmt.Printf("Min(10, 20) = %d\n", Min(10, 20))
	fmt.Printf("Min(3.14, 2.71) = %.2f\n", Min(3.14, 2.71))
	fmt.Printf(`Min("apple", "banana") = %s`+"\n", Min("apple", "banana"))
	fmt.Println()

	// ============================================================
	// 6. Zero Value
	// ============================================================
	fmt.Println("=== 6. Zero Value ===")
	fmt.Printf("Zero[int]() = %d\n", Zero[int]())
	fmt.Printf("Zero[string]() = %q\n", Zero[string]())
	fmt.Printf("Zero[bool]() = %v\n", Zero[bool]())
	fmt.Println()

	// ============================================================
	// 7. Generic Stack
	// ============================================================
	fmt.Println("=== 7. Generic Stack ===")

	// Stack of strings
	nameStack := Stack[string]{}
	nameStack.Push("Alice")
	nameStack.Push("Bob")
	nameStack.Push("Charlie")

	fmt.Printf("Stack size: %d\n", nameStack.Size())

	top, ok := nameStack.Peek()
	fmt.Printf("Peek: %s (exists: %v)\n", top, ok)

	for !nameStack.IsEmpty() {
		name, exists := nameStack.Pop()
		fmt.Printf("Popped: %s (exists: %v)\n", name, exists)
	}

	// Popping from an empty stack:
	emptyName, exists := nameStack.Pop()
	fmt.Printf("Pop from empty: %q (exists: %v)\n", emptyName, exists)

	// Stack of ints — same code, different type
	fmt.Println()
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)

	for !intStack.IsEmpty() {
		val, _ := intStack.Pop()
		fmt.Printf("Popped: %d\n", val)
	}
	fmt.Println()

	// ============================================================
	// 8. Generic Map (Transform)
	// ============================================================
	fmt.Println("=== 8. Generic Map (Transform) ===")

	// Transform int -> int (double)
	nums := []int{1, 2, 3, 4, 5}
	doubled := Map(nums, func(n int) int { return n * 2 })
	fmt.Printf("Original: %v\n", nums)
	fmt.Printf("Doubled:  %v\n", doubled)

	// Transform int -> string
	words := Map(nums, func(n int) string {
		if n%2 == 0 {
			return "even"
		}
		return "odd"
	})
	fmt.Printf("As words: %v\n", words)
	fmt.Println()

	// ============================================================
	// 9. Generic API Response (Real-World)
	// ============================================================
	fmt.Println("=== 9. Generic API Response ===")

	// Success with User data
	type User struct {
		ID   int
		Name string
	}

	userResp := NewSuccess(User{ID: 1, Name: "Alice"})
	fmt.Printf("Success response: %+v\n", userResp)

	// Error response (still need to specify the type)
	errResp := NewError[User]("user not found")
	fmt.Printf("Error response:   %+v\n", errResp)

	// Different type, same wrapper
	productResp := NewSuccess(map[string]float64{"price": 29.99})
	fmt.Printf("Product response: %+v\n", productResp)
	fmt.Println()

	// ============================================================
	// 10. GENERICS WITH INTERFACES (Advanced)
	// ============================================================
	fmt.Println("=== 10. Generics with Interfaces ===")
	// We can use the Named constraint with Greet because
	// Person and Pet (defined at package level) both have Name().
	person := Person{firstName: "Shalin"}
	fmt.Println(Greet(person))

	pet := Pet{nickname: "Buddy"}
fmt.Println(Greet(pet))


}

// ============================================================
// SUMMARY: Go Generics vs Python TypeVar
// ============================================================
//
// ┌──────────────────────────┬──────────────────────────┬──────────────────────────────┐
// │ Feature                  │ Go                       │ Python                       │
// ├──────────────────────────┼──────────────────────────┼──────────────────────────────┤
// │ Syntax                   │ [T any] AFTER func name  │ T = TypeVar('T') BEFORE func │
// │ Enforced at              │ COMPILE TIME (strict)    │ RUNTIME (hints only)         │
// │ Zero value of T          │ var zero T               │ No concept (return None?)    │
// │ Multiple type params     │ [T, U any]               │ T, U = TypeVar('T'), ...     │
// │ Constraint               │ Interface (structual)    │ bound=SomeClass              │
// │ Data structures          │ Stack[T any] struct      │ class Stack(Generic[T])      │
// │ Methods on generic type  │ func (s *Stack[T]) Push  │ def push(self, item: T)      │
// │ Performance              │ Zero-cost (compile time) │ Dynamic dispatch only        │
// │ Operators (+, <, etc.)   │ Via constraints package  │ Via __dunder__ methods       │
// │ Comptime type safety     │ ✅ YES                   │ ❌ No (mypy separate)        │
// └──────────────────────────┴──────────────────────────┴──────────────────────────────┘
//
// ─── Key Takeaways ───
//
// 1. Generics let you write ONE function/type for ANY type
//    — the compiler generates specialized code per type.
//
// 2. [T any] is the syntax: square brackets, not angle brackets.
//    T = type parameter, any = constraint (accept anything).
//
// 3. Monomorphization = zero runtime cost.
//    Go generates separate machine code for each concrete type.
//    No boxing, no interface dispatch overhead.
//
// 4. Use constraints for type safety:
//    [T constraints.Ordered] lets you use <, >, <=, >=.
//    [T constraints.Integer] lets you use +, -, *, /.
//
// 5. Use generics when TYPE RELATIONSHIPS matter:
//    - Input type = output type (Identity[T any] T)
//    - Multiple params share the same type (Min[T Ordered] a, b T)
//    - Data structure needs to hold any type (Stack[T any])
//
// 6. DON'T use generics when any works:
//    - Print[T any](v T) → use Print(v any) instead
//    - Only reach for generics when type relationships matter.
//
// 7. Returning zero: var zero T; return zero.
//    You CANNOT return nil for a generic type.
//
// 8. Generic types need instantiation: Stack[int]{}, not Stack{}.
//
// 9. Combined with interfaces: [T Named] where Named is an interface
//    with methods. T must have all those methods.
//
// 10. Go 1.18+ only. Before 2022, Go didn't have generics at all!
