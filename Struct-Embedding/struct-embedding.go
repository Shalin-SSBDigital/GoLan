// ============================================================
//  STRUCT EMBEDDING IN GO — with Python Comparisons
// ============================================================
//  Go does NOT have inheritance. Instead, it uses COMPOSITION
//  through struct embedding. Fields and methods of the embedded
//  type are "promoted" to the outer type.
//
//  This is Go's answer to the principle:
//    "Favor composition over inheritance."
//
// ┌──────────────────────────────┬──────────────────────────────────────────┐
// │        Go                    │              Python                      │
// ├──────────────────────────────┼──────────────────────────────────────────┤
// │ type Dog struct {            │ class Animal:                           │
// │     Animal        ← embedded │     def breathe(self): ...              │
// │ }                            │                                         │
// │                              │ class Dog(Animal):  ← inheritance       │
// │ dog.Breathe()   ← promoted   │ dog.breathe()      ← inherited method   │
// │                              │                                         │
// │ NO "is-a" hierarchy          │ "is-a" relationship                     │
// │ NO method resolution order   │ Method Resolution Order (MRO)           │
// │ NO super()                   │ super() to call parent                  │
// │ NO diamond problem           │ Diamond problem (resolved via MRO)      │
// │                              │                                         │
// │ Embedding = HAS-A (by value) │ Inheritance = IS-A (by reference)       │
// │ The embedded type is a FIELD │ The parent class is an ancestor         │
// │ with an AUTOMATIC name      │ accessed through the MRO                │
// └──────────────────────────────┴──────────────────────────────────────────┘
// ============================================================

package main

import "fmt"

// =============================================================================
// 1. BASIC EMBEDDING — Composition over Inheritance
// =============================================================================
// Go's embedding is just a field WITHOUT a name.
// The type name BECOMES the field name automatically.
//
// Python inheritance:
//   class Animal:
//       def __init__(self, name):
//           self.name = name
//       def breathe(self):
//           return f"{self.name} is breathing"
//
//   class Dog(Animal):
//       def __init__(self, name, breed):
//           super().__init__(name)
//           self.breed = breed
//
// Go embedding (equivalent):
//   type Animal struct { Name string }
//   func (a Animal) Breathe() string { ... }
//
//   type Dog struct {
//       Animal          ← embedded (no field name)
//       Breed string
//   }
//
// KEY DIFFERENCE: Python Dog IS-A Animal (inheritance chain).
// Go Dog HAS-A Animal (composition). The relationship is
// fundamentally different.

// Base type — like a Python parent class
type Animal struct {
	Name string
}

func (a Animal) Breathe() string {
	return fmt.Sprintf("%s is breathing", a.Name)
}

func (a Animal) Eat() string {
	return fmt.Sprintf("%s is eating", a.Name)
}

// Dog embeds Animal — no "extends", no "inherits", no "super()"
type Dog struct {
	Animal        // ← embedded (no field name = automatic promotion)
	Breed string
}

// =============================================================================
// 2. FIELD AND METHOD PROMOTION
// =============================================================================
// When you embed a type, its fields and methods are PROMOTED to the
// outer type. You can access them as if they were defined directly
// on the outer type.
//
// Python:
//   dog = Dog(name="Rover", breed="Lab")
//   dog.breathe()     # inherited from Animal — no super() needed to CALL
//   dog.name          # inherited field
//
// Go:
//   dog := Dog{Animal: Animal{Name: "Rover"}, Breed: "Lab"}
//   dog.Breathe()     // promoted — calls Animal.Breathe()
//   dog.Name          // promoted field — dog.Animal.Name
//
// ⚠️ KEY DIFFERENCE:
//   Python: inherited attributes are RESOLVED at RUNTIME via MRO.
//   Go:     promoted fields/methods are RESOLVED at COMPILE TIME.
//   Go is FASTER (no runtime lookup) but LESS FLEXIBLE (no monkey-patching).

func demoPromotion() {
	fmt.Println("=== Field & Method Promotion ===")

	// Creating a Dog with embedded Animal
	dog := Dog{
		Animal: Animal{Name: "Rover"},
		Breed:  "Labrador",
	}

	// Promoted fields — access as if on Dog directly
	fmt.Println("Name:", dog.Name)         // ← promoted from Animal
	fmt.Println("Breed:", dog.Breed)       // ← own field

	// Explicit access (also works):
	fmt.Println("Name (explicit):", dog.Animal.Name)

	// Promoted methods
	fmt.Println(dog.Breathe())             // ← promoted from Animal
	fmt.Println(dog.Eat())                 // ← promoted from Animal
}

// =============================================================================
// 3. METHOD OVERRIDING / SHADOWING
// =============================================================================
// If the outer type defines a method with the same name as an
// embedded type's method, the outer method TAKES PRECEDENCE (shadows).
// The embedded method is still accessible via explicit reference.
//
// Python:
//   class Dog(Animal):
//       def breathe(self):
//           base = super().breathe()
//           return f"{base} and wagging tail"
//
//   # super() lets you call the PARENT's version explicitly.
//
// Go:
//   func (d Dog) Breathe() string {
//       base := d.Animal.Breathe()  // ← explicit call to embedded
//       return base + " and wagging tail"
//   }
//
//   # No super() keyword — use the embedded type name explicitly.

// Override Breathe on Dog (shadows Animal.Breathe)
func (d Dog) Breathe() string {
	// Call the embedded Animal's Breathe EXPLICITLY
	base := d.Animal.Breathe()
	return base + " and wagging tail"
}

// =============================================================================
// 4. MULTIPLE EMBEDDING — Go's Answer to Multiple Inheritance
// =============================================================================
// Go allows embedding MULTIPLE types in one struct.
//
// Python:
//   class A:
//       def method(self): print("A")
//   class B:
//       def method(self): print("B")
//   class C(A, B):       # multiple inheritance with MRO
//       pass
//   c = C(); c.method()  # → A  (MRO resolves: C → A → B)
//
// Go:
//   type A struct{}
//   func (A) Method() { fmt.Println("A") }
//   type B struct{}
//   func (B) Method() { fmt.Println("B") }
//   type C struct {
//       A    // embedded
//       B    // embedded (compile error if both have Method!)
//   }
//
// ⚠️ CRITICAL DIFFERENCE — NO MRO IN GO:
//   Python resolves ambiguous method names via MRO (C3 linearization).
//   Go REJECTS ambiguity at COMPILE TIME — if A and B both have Method(),
//   and C doesn't provide its own, C.Method() is a COMPILE ERROR.
//   You MUST disambiguate explicitly: c.A.Method() or c.B.Method()
//
//   Go prioritizes SAFETY over flexibility here. If it's ambiguous,
//   the compiler refuses to guess. Python guesses for you (via MRO).

type Flying struct{}

func (Flying) Move() string {
	return "I'm flying!"
}

type Swimming struct{}

func (Swimming) Move() string {
	return "I'm swimming!"
}

// Duck embeds both Flying and Swimming.
// BOTH have Move() — this would be a compile error if we called
// duck.Move() without disambiguating.
type Duck struct {
	Flying
	Swimming
	Name string
}

// To resolve the ambiguity, Duck defines its OWN Move():
func (d Duck) Move() string {
	return "I'm doing both!"
}

// Or access explicitly:
func (d Duck) Fly() string {
	return d.Flying.Move()
}

func (d Duck) Swim() string {
	return d.Swimming.Move()
}

// =============================================================================
// 5. EMBEDDING vs NAMED FIELD — Two Forms of Composition
// =============================================================================
//
// EMBEDDED (no name):           NAMED FIELD:
//   type Car struct {               type Car struct {
//       Engine               ← name    eng Engine       ← explicit name
//   }                               }
//   car.Start()  ← promoted       car.eng.Start()  ← must qualify
//
// Use EMBEDDING when:            Use NAMED FIELD when:
//   - The relationship IS-A-like   - The relationship is HAS-A
//   - You want method promotion    - You want encapsulation
//   - You're modeling behavior     - You need multiple instances
//                                   - You want a clear namespace

type Engine struct {
	Horsepower int
}

func (e Engine) Start() string {
	return fmt.Sprintf("Engine with %d HP started", e.Horsepower)
}

// Embedded — method promoted
type Car struct {
	Engine             // ← embedded (promoted)
	Model    string
}

// Named — accessed explicitly
type Boat struct {
	eng Engine         // ← named field (NOT promoted)
	Name string
}

func demoEmbeddedVsNamed() {
	fmt.Println("\n=== Embedded vs Named Field ===")

	// Embedded: methods promoted directly
	car := Car{Engine: Engine{Horsepower: 200}, Model: "Sedan"}
	fmt.Println("Car:", car.Start())   // ← promoted, no qualifier
	fmt.Println("Car HP:", car.Horsepower) // ← promoted field

	// Named: must qualify
	boat := Boat{eng: Engine{Horsepower: 150}, Name: "Yacht"}
	fmt.Println("Boat:", boat.eng.Start()) // ← must use boat.eng
	fmt.Println("Boat HP:", boat.eng.Horsepower)

	fmt.Println("\nKey difference:")
	fmt.Println("  car.Start()       ← promoted (shorter, but implicit)")
	fmt.Println("  boat.eng.Start()  ← qualified (longer, but explicit)")
	fmt.Println("  Go convention: embed for IS-A-LIKE behavior composition,")
	fmt.Println("  use named fields for strict HAS-A relationships.")
}

// =============================================================================
// 6. EMBEDDING INTERFACES IN STRUCTS
// =============================================================================
// You can embed INTERFACES in structs too. This is used for:
//   - Partial implementations (embed + override)
//   - Dependency injection
//   - Testing (mock by embedding real interface)
//
// Python:
//   from typing import Protocol
//
//   class Logger(Protocol):
//       def log(self, msg: str): ...
//
//   class DevNull:
//       def log(self, msg): pass  # silent logger
//
//   class App:
//       def __init__(self, logger: Logger):
//           self.logger = logger or DevNull()
//
// Go uses embedding of interfaces in structs for a similar pattern:

type Logger interface {
	Log(msg string)
}

// ConsoleLogger writes to stdout
type ConsoleLogger struct{}

func (ConsoleLogger) Log(msg string) {
	fmt.Println("[LOG]", msg)
}

// Server embeds a Logger — if no logger is set, calls PANIC (nil interface!)
// But we can provide a DefaultLogger to avoid that.
type Server struct {
	Logger // ← embedded interface (can be nil!)
	Port   int
}

// Start uses the embedded Logger (or panics if nil)
func (s Server) Start() {
	s.Log("Server starting on port " + fmt.Sprint(s.Port))
}

// =============================================================================
// 7. EMBEDDING with POINTERS
// =============================================================================
// You can embed a POINTER to a type. This allows SHARED state between
// the outer type and the embedded type — mutations are visible to both.
//
// Python:
//   class SharedState:
//       def __init__(self):
//           self.count = 0
//
//   class Counter:
//       def __init__(self):
//           self.state = SharedState()  # shared by reference
//
//   # All Python composition is by reference (like embedding a pointer in Go)

type SharedCounter struct {
	Count int
}

func (s *SharedCounter) Increment() {
	s.Count++
}

// CounterGroup embeds a *SharedCounter — all instances share the same counter
type CounterGroup struct {
	*SharedCounter // ← pointer embedding (shared state)
	Label  string
}

// =============================================================================
// 8. COMPLETE COMPARISON: Go Embedding vs Python Inheritance
// =============================================================================
//
// ┌──────────────────────────┬──────────────────────────┬──────────────────────────────┐
// │ Feature                  │ Go (Embedding)            │ Python (Inheritance)         │
// ├──────────────────────────┼──────────────────────────┼──────────────────────────────┤
// │ Relationship             │ HAS-A (composition)       │ IS-A (inheritance)           │
// │ Mechanism                │ Anonymous field in struct │ Class in MRO chain           │
// │ Method promotion         │ Automatic to outer type   │ Automatic via MRO            │
// │ Override parent          │ Define same method        │ Define same method           │
// │ Call parent version      │ embeddedType.Method()     │ super().method()             │
// │ Multiple "parents"       │ Multiple embedding        │ Multiple inheritance         │
// │ Ambiguity resolution     │ COMPILE ERROR (must fix)  │ MRO (C3 linearization)       │
// │ Constructor chain        │ No constructors           │ super().__init__() chain     │
// │ Private access to parent │ None (everything visible) │ None (Python has no private) │
// │ Runtime flexibility      │ None (compile-time)       │ Monkey-patching, super()     │
// │ Diamond problem          │ Impossible (compile err)  │ Handled by MRO               │
// │ Interface embedding      │ Supported                 │ No direct equivalent         │
// │ Pointer embedding        │ Supported (shared state)  │ All refs are "pointers"      │
// └──────────────────────────┴──────────────────────────┴──────────────────────────────┘

// =============================================================================
// DEMO — Putting it all together
// =============================================================================

func main() {
	// --- Promotion demo ---
	demoPromotion()

	// --- Method overriding ---
	fmt.Println("\n=== Method Overriding ===")
	dog := Dog{
		Animal: Animal{Name: "Rover"},
		Breed:  "Labrador",
	}
	fmt.Println(dog.Eat())     // ← promoted from Animal (no override)
	fmt.Println(dog.Breathe()) // ← Dog's override (not Animal's)
	// Can still access the original:
	fmt.Println(dog.Animal.Breathe()) // ← explicit access to Animal's version

	// --- Multiple embedding ---
	fmt.Println("\n=== Multiple Embedding ===")
	duck := Duck{
		Flying:   Flying{},
		Swimming: Swimming{},
		Name:     "Donald",
	}
	fmt.Println("Duck:", duck.Move())       // Duck's own method
	fmt.Println("Fly:", duck.Fly())         // explicit: calls Flying.Move()
	fmt.Println("Swim:", duck.Swim())       // explicit: calls Swimming.Move()

	// --- Embedded vs Named ---
	demoEmbeddedVsNamed()

	// --- Shared state via pointer embedding ---
	fmt.Println("\n=== Pointer Embedding (Shared State) ===")
	counter := &SharedCounter{Count: 0}
	group1 := CounterGroup{SharedCounter: counter, Label: "Group A"}
	group2 := CounterGroup{SharedCounter: counter, Label: "Group B"}

	group1.Increment()
	group2.Increment()
	group1.Increment()

	fmt.Printf("Counter value: %d (shared across groups)\n", counter.Count)
	// Both groups share the SAME counter because they both point to it.

	// --- When to use embedding vs inheritance ---
	fmt.Println("\n=== Philosophy: Embedding vs Inheritance ==")
	fmt.Println("Go says: What does my type HAVE (composition)?")
	fmt.Println("Python says: What IS my type (inheritance hierarchy)?")
	fmt.Println()
	fmt.Println("  Go's approach leads to flatter, more flexible hierarchies.")
	fmt.Println("  Python's approach leads to deeper, more rigid taxonomies.")
	fmt.Println()
	fmt.Println("  If you find yourself writing 'is-a' in Go, you're probably")
	fmt.Println("  overusing embedding. Ask: does this type HAVE the behavior?")
	fmt.Println("  If yes, use composition (embedding or named field).")
	fmt.Println("  If you need polymorphism, use INTERFACES, not embedding.")

	// ============================================================
	//  PRACTICAL GUIDELINES
	// ============================================================
	//  ┌─────────────────────────────────────────────────────────────────────┐
	//  │ 1. Embed for BEHAVIOR, not just fields. If you're only embedding  │
	//  │    to get field promotion, use a named field instead.              │
	//  │                                                                    │
	//  │ 2. Avoid embedding types that you don't want to expose publicly.    │
	//  │    All promoted methods become part of the outer type's API.        │
	//  │                                                                    │
	//  │ 3. Multiple embedding is powerful but risky — disambiguate early.   │
	//  │    The compiler forces you to, which is a FEATURE not a bug.        │
	//  │                                                                    │
	//  │ 4. Embedding + Interfaces = Go's version of polymorphism.           │
	//  │    A struct can embed an interface (for testing, DI, stubbing).     │
	//  │                                                                    │
	//  │ 5. Pointer embedding creates shared state — use with caution.       │
	//  │    If embedded by value, each outer type has its OWN copy.          │
	//  │                                                                    │
	//  │ 6. There is NO super() in Go. Type name + dot = explicit parent.    │
	//  │    e.g., d.Animal.Breathe() instead of super().breathe()           │
	//  └─────────────────────────────────────────────────────────────────────┘
	//
	// Bottom line: Go embedding is COMPOSITION that LOOKS like inheritance
	// but is fundamentally different underneath. It's safer (no MRO ambiguity,
	// no diamond problem) but less flexible (no runtime monkey-patching).
	// Python inheritance is runtime-resolved; Go embedding is compile-time
	// resolved. Go replaces the "is-a" hierarchy with a "has-a" composition.
	// ============================================================
}
