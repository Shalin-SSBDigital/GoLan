// ============================================================
//  STRUCTS IN GO — with Python Comparisons
// ============================================================
//  A struct is a collection of named fields. It's Go's version
//  of a class — data grouped together (without the behavior).
//
// ┌──────────────────────┬──────────────────────────────────────────────┐
// │        Go            │                  Python                      │
// ├──────────────────────┼──────────────────────────────────────────────┤
// │ type User struct {   │ class User:                                  │
// │     Name string      │     def __init__(self, name, age):           │
// │     Age  int         │         self.name = name                     │
// │ }                    │         self.age = age                       │
// ├──────────────────────┼──────────────────────────────────────────────┤
// │ DATA ONLY            │ class = data + methods                       │
// │ Methods go SEPARATELY│ methods live INSIDE the class body           │
// ├──────────────────────┼──────────────────────────────────────────────┤
// │ u := User{           │ u = User(name="Alice", age=30)              │
// │     Name: "Alice",   │     # positional also works                 │
// │     Age:  30,        │                                            │
// │ }                    │                                            │
// ├──────────────────────┼──────────────────────────────────────────────┤
// │ VALUE SEMANTICS      │ REFERENCE SEMANTICS                         │
// │ Assign = COPY        │ Assign = SAME object                        │
// │ Pass to func = COPY  │ Pass to func = SAME object (by ref)         │
// └──────────────────────┴──────────────────────────────────────────────┘
// ============================================================

package main

import (
	"fmt"
	"reflect"
	"strings"
)

// =============================================================================
// 1. DEFINING A STRUCT
// =============================================================================
// A struct is defined with the `type` keyword followed by the name and
// the `struct` keyword with fields inside curly braces.
//
// Python equivalent:
//   class User:
//       def __init__(self, name, email, age):
//           self.name = name
//           self.email = email
//           self.age = age
//
// KEY DIFFERENCE:
//   Go: No __init__ constructor. Fields are just declared.
//   Fields are PUBLIC if they start with uppercase, PRIVATE if lowercase.
//   Python: Everything is public by convention (self.name).

type User struct {
	Name  string
	Email string
	Age   int
}

// =============================================================================
// 2. CREATING STRUCT INSTANCES
// =============================================================================
// Go has THREE ways to create a struct (Python has one: __init__).

func demoCreation() {
	fmt.Println("=== Struct Creation ===")

	// Method 1: Field names (recommended) — like Python kwargs
	//   Python: u = User(name="Alice", email="a@b.com", age=30)
	u1 := User{
		Name:  "Alice",
		Email: "a@b.com",
		Age:   30,
	}
	fmt.Println("u1:", u1)

	// Method 2: Positional (fragile, avoid if many fields)
	//   Python: u = User("Bob", "b@b.com", 25)
	u2 := User{"Bob", "b@b.com", 25}
	fmt.Println("u2:", u2)

	// Method 3: Zero-value (all fields = Go zero values)
	//   Python: NO equivalent — Python requires __init__ to run
	//   In Go, variables always have a value (never nil for structs)
	var u3 User
	fmt.Println("u3 (zero value):", u3)
	// Fields are "" and 0 automatically — no __init__ needed

	// Method 4: Partial initialization (omitted fields stay zero)
	//   Python: would need default params in __init__
	u4 := User{Name: "Charlie"}
	fmt.Println("u4 (partial):", u4)
}

// =============================================================================
// 3. ACCESSING AND MODIFYING FIELDS
// =============================================================================
// Dot notation — same as Python.

func demoAccess() {
	fmt.Println("\n=== Accessing & Modifying Fields ===")

	u := User{Name: "Alice", Email: "alice@example.com", Age: 30}

	// Read fields (same as Python: u.name)
	fmt.Println("Name:", u.Name)
	fmt.Println("Email:", u.Email)
	fmt.Println("Age:", u.Age)

	// Modify fields (same as Python: u.name = "Bob")
	u.Name = "Bob"
	fmt.Println("After rename:", u)
}

// =============================================================================
// 4. VALUE SEMANTICS — Structs are COPIES
// =============================================================================
// This is the BIGGEST difference from Python.
//
// Python:
//   u1 = User("Alice", 30)
//   u2 = u1      # u2 REFERENCES the SAME object
//   u2.name = "Bob"    # u1.name is ALSO "Bob" now!
//
// Go:
//   u1 := User{Name: "Alice", Age: 30}
//   u2 := u1    # u2 is a FULL COPY of u1
//   u2.Name = "Bob"    # u1 is UNCHANGED!

func demoValueSemantics() {
	fmt.Println("\n=== Value Semantics (Structs are COPIED) ===")

	u1 := User{Name: "Alice", Age: 30}
	u2 := u1 // ← COPY, not reference!

	u2.Name = "Bob"

	fmt.Println("u1.Name:", u1.Name) // "Alice" (unchanged)
	fmt.Println("u2.Name:", u2.Name) // "Bob" (changed)

	// Python comparison:
	fmt.Println("\nPython would say:")
	fmt.Println("  u2 = u1  → u2 REFERENCES u1 (same object)")
	fmt.Println("  u2.name = 'Bob' → u1.name is ALSO 'Bob'")
	fmt.Println("Go says:")
	fmt.Println("  u2 := u1 → u2 COPIES u1 (independent)")
	fmt.Println("  u2.Name = 'Bob' → u1.Name is STILL 'Alice'")
}

// =============================================================================
// 5. POINTER TO STRUCT — Go's Reference Semantics
// =============================================================================
// To get Python-like reference behavior, use a POINTER (*User).
//
// Python:
//   u = User("Alice", 30)    # u is always a reference
//   def change_name(u):       # u is passed by reference
//       u.name = "Bob"        # modifies the ORIGINAL
//
// Go — two ways to create a pointer:
//   Method A: &User{...} — address of a literal
//   Method B: new(User)  — returns *User (zero values)

func demoPointers() {
	fmt.Println("\n=== Pointer to Struct (Reference Semantics) ===")

	// Creating a pointer to a struct
	u1 := &User{Name: "Alice", Age: 30}
	// u1 is *User (pointer to User)

	// u2 points to the SAME struct (like Python assignment)
	u2 := u1

	u2.Name = "Bob"

	fmt.Println("u1.Name:", u1.Name) // "Bob" — changed through pointer!
	fmt.Println("u2.Name:", u2.Name) // "Bob"

	// Go auto-dereferences: u1.Name works even though u1 is *User
	// In Go, (*u1).Name and u1.Name are the same on pointers

	fmt.Println("\nThis is how Python works ALL the time.")
	fmt.Println("In Go, you CHOOSE: copy (value) or share (pointer).")
}

// =============================================================================
// 6. NESTED STRUCTS — Structs Inside Structs
// =============================================================================
// A struct field can itself be a struct.
//
// Python:
//   class Address:
//       def __init__(self, city, zip):
//           self.city = city
//           self.zip = zip
//
//   class Person:
//       def __init__(self, name, address):
//           self.name = name
//           self.address = address   # ← nested object

type Address struct {
	City    string
	ZipCode string
}

type Person struct {
	Name    string
	Address Address // ← nested struct (like Python composition)
}

func demoNested() {
	fmt.Println("\n=== Nested Structs ===")

	p := Person{
		Name: "Alice",
		Address: Address{
			City:    "New York",
			ZipCode: "10001",
		},
	}

	// Access nested fields with dot chain
	fmt.Println("Name:", p.Name)
	fmt.Println("City:", p.Address.City)
	fmt.Println("Zip:", p.Address.ZipCode)

	// Modify nested fields
	p.Address.City = "Boston"
	fmt.Println("New city:", p.Address.City)
}

// =============================================================================
// 7. ANONYMOUS STRUCTS — One-off Structs Without a Name
// =============================================================================
// Python has no direct equivalent — you'd use a dict or SimpleNamespace.

func demoAnonymous() {
	fmt.Println("\n=== Anonymous Structs ===")

	// Define and create in one shot — no `type` keyword needed
	book := struct {
		Title  string
		Author string
		Pages  int
	}{
		Title:  "The Go Programming Language",
		Author: "Donovan & Kernighan",
		Pages:  400,
	}

	fmt.Println("Book:", book.Title)
	fmt.Println("Author:", book.Author)

	// Useful for:
	//   - Test fixtures
	//   - JSON response shaping (one-off data)
	//   - Grouping related variables temporarily
	fmt.Println("\nPython equivalent would be a dict or SimpleNamespace")
}

// =============================================================================
// 8. STRUCT TAGS — Metadata on Fields
// =============================================================================
// Tags are NOT accessible at runtime through normal code. They are
// STRING literals attached to fields, read via reflection.
//
// Python equivalent: There is NO direct equivalent. Closest would be
//   - dataclass field metadata
//   - Pydantic's Field(alias=..., description=...)
//   - Attrs field attributes
//
//   @dataclass
//   class Config:
//       host: str = field(metadata={"env": "HOST"})
//       port: int = field(metadata={"env": "PORT"})

type Config struct {
	Host string `json:"host" env:"HOST"`
	Port int    `json:"port" env:"PORT"`
	Debug bool  `json:"debug" env:"DEBUG"`
}

func demoStructTags() {
	fmt.Println("\n=== Struct Tags ===")

	c := Config{Host: "localhost", Port: 8080, Debug: true}

	// Read tags via reflection (not something you'd do in production often)
	t := reflect.TypeOf(c)
	for i := range 3 {
		field := t.Field(i)
		fmt.Printf("Field: %-5s  json: %-10s  env: %s\n",
			field.Name, field.Tag.Get("json"), field.Tag.Get("env"))
	}

	fmt.Println("\nTags are used by:")
	fmt.Println("  - encoding/json (marshaling)")
	fmt.Println("  - encoding/xml")
	fmt.Println("  - ORMs (GORM, etc.)")
	fmt.Println("  - Validators")
	fmt.Printf("  - Config loaders\n\n")
	fmt.Println("They are just strings — Go doesn't enforce tag structure.")
}

// =============================================================================
// 9. COMPARING STRUCTS
// =============================================================================
// Structs are COMPARABLE if all their fields are comparable.
// You can use == and != directly (no __eq__ needed).
//
// Python:
//   class User:
//       def __init__(self, name, age):
//           self.name = name
//           self.age = age
//       def __eq__(self, other):         # ← must implement
//           return self.name == other.name and self.age == other.age
//
//   u1 = User("Alice", 30)
//   u2 = User("Alice", 30)
//   u1 == u2  # False WITHOUT __eq__, True WITH __eq__
//
// Go:
//   u1 := User{Name: "Alice", Age: 30}
//   u2 := User{Name: "Alice", Age: 30}
//   u1 == u2  // TRUE — Go compares field-by-field automatically!
//              // No __eq__ equivalent needed.
//
// ⚠️ LIMITATION: If ANY field is NOT comparable (slice, map, function),
//    the struct becomes INCOMPARABLE — compile error on ==.

type Comparable struct {
	A int
	B string
	C bool
}

type NotComparable struct {
	D []int // slices are NOT comparable
}

func demoComparison() {
	fmt.Println("\n=== Struct Comparison ===")

	a := Comparable{A: 1, B: "hello", C: true}
	b := Comparable{A: 1, B: "hello", C: true}
	c := Comparable{A: 2, B: "hello", C: true}

	fmt.Println("a == b:", a == b) // true
	fmt.Println("a == c:", a == c) // false

	// NotComparable would cause: struct containing []int cannot be compared
	// nc1 := NotComparable{D: []int{1, 2, 3}}
	// nc2 := NotComparable{D: []int{1, 2, 3}}
	// fmt.Println(nc1 == nc2)  // ← COMPILE ERROR!

	fmt.Println("\nPython requires __eq__ for custom equality.")
	fmt.Println("Go compares structs FIELD BY FIELD automatically.")
}

// =============================================================================
// 10. EMPTY STRUCT — Zero Bytes of Memory
// =============================================================================
// The empty struct `struct{}{}` takes ZERO bytes of memory.
// It's used as a signal, not data — like Python's None or NoneType.
//
// Python equivalents:
//   - None as a sentinel value
//   - object() as a key in sets
//   - threading.Event() as a signal

type empty struct{} // zero-size type

var sentinel = struct{}{} // reusable sentinel value

func demoEmptyStruct() {
	fmt.Println("\n=== Empty Struct (0 bytes) ===")

	var e empty
	fmt.Printf("Size of empty struct: %d bytes\n", reflect.TypeOf(e).Size())

	// Common uses:
	// 1. Set implementation (map[T]struct{})
	//    Python: set = {1, 2, 3}
	//    Go:     set := map[int]struct{}{1: {}, 2: {}, 3: {}}
	set := map[string]struct{}{
		"apple":  {},
		"banana": {},
	}
	fmt.Println("Set contains 'apple':", contains(set, "apple"))
	fmt.Println("Set contains 'grape':", contains(set, "grape"))

	// 2. Channel signals (chan struct{})
	//    Python: threading.Event()
	fmt.Println("\nEmpty struct — zero allocation, used for signaling.")
}

func contains(set map[string]struct{}, key string) bool {
	_, ok := set[key]
	return ok
}

// =============================================================================
// 11. CONSTRUCTOR / FACTORY PATTERN
// =============================================================================
// Go has NO __init__, NO __new__, NO constructors.
// Convention: use a function named NewTypeName().
//
// Python:          __init__ called automatically
// Go:              NewUser() is a CONVENTION, not a language feature
//
// Why use a constructor?
//   - Validation (ensure fields are valid on creation)
//   - Default values (Go initializes to zero, but sometimes you want custom)
//   - Pointer return (consistency: always return *T)

func NewUser(name, email string, age int) (*User, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	if email == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}
	if age < 0 {
		return nil, fmt.Errorf("age cannot be negative")
	}

	return &User{
		Name:  name,
		Email: email,
		Age:   age,
	}, nil
}

func demoConstructor() {
	fmt.Println("\n=== Constructor / Factory Pattern ===")

	// Using the constructor (returns *User so nil is possible for error)
	u, err := NewUser("Alice", "alice@example.com", 30)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Created:", u)
	}

	// Invalid — error case
	_, err = NewUser("", "bob@example.com", 25)
	if err != nil {
		fmt.Println("Expected error:", err)
	}
}

// =============================================================================
// 12. EMBEDDED vs NAMED FIELDS
// =============================================================================
// A field can be:
//   1. NAMED:    fieldName Type    — explicit access via fieldName
//   2. EMBEDDED: Type (no name)    — promoted fields/methods
//
// Python has no embedding — all fields are "named" attributes.
// See `Struct-Embedding/` for the full treatment.

type Engine struct {
	Horsepower int
}

// Embedded — methods promoted to Car
type Car struct {
	Engine
	Model string
}

// Named — must access explicitly
type Boat struct {
	Eng   Engine // named, not embedded
	Model string
}

func demoEmbeddedVsNamed() {
	fmt.Println("\n=== Embedded vs Named Fields ===")

	car := Car{Engine: Engine{Horsepower: 200}, Model: "Sedan"}
	// Promoted fields — access as if on Car directly
	fmt.Printf("Car %s: %d HP\n", car.Model, car.Horsepower)

	boat := Boat{Eng: Engine{Horsepower: 150}, Model: "Yacht"}
	// Must use the field name
	fmt.Printf("Boat %s: %d HP\n", boat.Model, boat.Eng.Horsepower)
}

// =============================================================================
// 13. STRUCT WITH METHODS (Quick Demo)
// =============================================================================
// Although structs are "data only", you attach METHODS separately.
// This is covered in depth by `Methods/` — just showing the syntax here.

func (u User) Greet() string {
	return fmt.Sprintf("Hi, I'm %s and I'm %d years old.", u.Name, u.Age)
}

func (u *User) Birthday() {
	u.Age++ // modifies original — pointer receiver
}

func demoStructMethods() {
	fmt.Println("\n=== Structs + Methods ===")

	u := User{Name: "Alice", Age: 30}
	fmt.Println(u.Greet())

	u.Birthday()
	fmt.Println("After birthday:", u.Age)
	fmt.Println(u.Greet())
}

// =============================================================================
// 14. COMPLETE COMPARISON TABLE
// =============================================================================
//
// ┌────────────────────────────┬────────────────────────────────────┬──────────────────────────────────────┐
// │        Feature             │            Go                     │              Python                   │
// ├────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Definition                 │ type T struct { Fields }           │ class T:                             │
// │                            │                                    │     def __init__(self, ...)          │
// ├────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Construction               │ T{Field: val}  (literal)           │ T(field=val)  (calls __init__)       │
// ├────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Constructor                │ NewT() function (convention)       │ __init__ (language feature)          │
// ├────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Default values             │ Zero values (0, "", false, nil)    │ __init__ parameters with defaults    │
// ├────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Methods                    │ func (t T) Method() {}  (separate) │ def method(self):  (inside class)    │
// ├────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Assignment                 │ COPY (value semantics)             │ REFERENCE (object identity)          │
// ├────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Pass to function           │ COPY by default                    │ REFERENCE by default                 │
// ├────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Reference semantics        │ Use *T (pointer) explicitly        │ Always (everything is an object ref) │
// ├────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Equality (==)              │ Field-by-field (if comparable)     │ Uses __eq__ (must be defined)        │
// ├────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Nested types               │ struct field of struct type        │ object attribute of class type       │
// ├────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Inheritance                │ Embedding (composition)            │ Class inheritance (MRO, super())     │
// ├────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Access control             │ Uppercase = public, lowercase = pvt│ Everything public (convention: _pvt) │
// ├────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Tags / metadata            │ Struct tags (string tags)          │ dataclass metadata, Pydantic Field() │
// ├────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Anonymous structs          │ struct{...}{...} inline            │ No direct equivalent (use dict)      │
// ├────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Empty struct               │ struct{}{} (0 bytes)               │ No direct equivalent                 │
// └────────────────────────────┴────────────────────────────────────┴──────────────────────────────────────┘

// =============================================================================
// DEMO — Putting it all together
// =============================================================================

func main() {
	demoCreation()
	demoAccess()
	demoValueSemantics()
	demoPointers()
	demoNested()
	demoAnonymous()
	demoStructTags()
	demoComparison()
	demoEmptyStruct()
	demoConstructor()
	demoEmbeddedVsNamed()
	demoStructMethods()

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("STRUCTS CHEAT SHEET")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println(`
  // Define struct
  type T struct { Field Type }

  // Create (recommended — named fields)
  t := T{Field: value}

  // Create (positional — fragile)
  t := T{val1, val2}

  // Create (zero value — no __init__ needed)
  var t T

  // Pointer
  t := &T{Field: value}

  // Constructor (convention)
  func NewT(args) (*T, error) {
      if err { return nil, err }
      return &T{...}, nil
  }

  // Access / modify
  t.Field = newValue

  // Anonymous struct (one-off)
  x := struct{ A int }{A: 1}

  // Empty struct (0 bytes, for signaling / sets)
  set := map[int]struct{}{1: {}}
`)

}
