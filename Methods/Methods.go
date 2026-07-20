package main

import (
	"fmt"
	"math"
)

// =============================================================================
// METHODS IN GO
// =============================================================================
//
// A method in Go is a function with a special "receiver" argument that sits
// between the `func` keyword and the function name.
//
// Syntax:    func (receiver ReceiverType) MethodName(params) ReturnType { ... }
//
// Unlike Python where methods are just functions defined inside a class body
// and always receive `self` as the first argument, Go lets you attach methods
// to ANY named type — not just structs. There is no `class` keyword.
//
// ┌──────────────────────┬──────────────────────────────────────────────┐
// │        Go            │                   Python                     │
// ├──────────────────────┼──────────────────────────────────────────────┤
// │ func (r Rect) Area() │ class Rect:                                  │
// │   ...                │     def area(self):                          │
// │                      │         ...                                  │
// ├──────────────────────┼──────────────────────────────────────────────┤
// │ Receiver is SEPARATE │ self is EXPLICIT but always the              │
// │ from the parameter   │ first parameter — you must write it.        │
// │ list — a special slot│                                              │
// ├──────────────────────┼──────────────────────────────────────────────┤
// │ No class keyword.    │ Everything lives inside `class`.             │
// │ Methods attach to    │                                              │
// │ any named type.      │                                              │
// └──────────────────────┴──────────────────────────────────────────────┘

// =============================================================================
// 1. BASIC METHOD — Value Receiver
// =============================================================================

// Define a simple struct (like a Python class with only data).
type Rect struct {
	Width  float64
	Height float64
}

// Area is a method on Rect with a VALUE receiver (r Rect).
// Python equivalent:
//   class Rect:
//       def area(self):
//           return self.width * self.height
//
// Key difference: in Go, `r` is a COPY of the original Rect. Modifying
// r.Width or r.Height inside Area() does NOT affect the caller's value.
// In Python, `self` is always a REFERENCE — mutations inside methods
// always affect the original object.
func (r Rect) Area() float64 {
	return r.Width * r.Height
}

// Scale — value receiver, CANNOT modify the original.
// This is like a Python method that returns a new object instead of
// mutating in place:
//     def scale(self, factor):
//         return Rect(self.width * factor, self.height * factor)
func (r Rect) Scale(factor float64) Rect {
	return Rect{Width: r.Width * factor, Height: r.Height * factor}
}

// =============================================================================
// 2. POINTER RECEIVER — MUTATING THE ORIGINAL
// =============================================================================

// To modify the receiver (like Python's self mutation), use a POINTER receiver.
// The receiver type is *Rect instead of Rect.
func (r *Rect) ScaleInPlace(factor float64) {
	r.Width *= factor  // This modifies the original struct
	r.Height *= factor
	// Python equivalent:
	//   def scale_in_place(self, factor):
	//       self.width *= factor
	//       self.height *= factor
	//
	// In Python this is the DEFAULT behavior — self is always a reference.
	// In Go, you MUST explicitly use a pointer receiver if you want mutations
	// to be visible outside the method. This is Go's philosophy of
	// explicitness over implicit behavior.
}

// =============================================================================
// 3. WHEN TO USE VALUE vs POINTER RECEIVER
// =============================================================================
//
// ┌─────────────────────────────┬──────────────────────────────────────────┐
│ // VALUE RECEIVER             │ POINTER RECEIVER                         │
│ //───────────────────────────│──────────────────────────────────────────│
│ // - Method does NOT mutate  │ - Method mutates the receiver            │
│ // - Receiver is small (int, │ - Receiver is large (big struct) —       │
│ //   float, small struct)    │   avoids copying                         │
│ // - Immutable semantics     │ - Consistency: if ANY method on a type   │
│ //                           │   uses a pointer receiver, usually ALL   │
│ //                           │   should (Go convention)                 │
│ //                           │ - When you need to share mutable state   │
│ └─────────────────────────────┴──────────────────────────────────────────┘
//
// Python note: Python completely hides this distinction — self is always a
// reference. This is simpler but can lead to accidental mutations. Go forces
// you to think about ownership and mutation at the call site.

// =============================================================================
// 4. METHODS ON ANY TYPE — Not Just Structs
// =============================================================================

// This is a BIG difference from Python. In Python, methods live inside class
// bodies. In Go, you can attach methods to ANY named type, including primitive
// types like int, float, string, etc.

// Define a named type based on float64
type Celsius float64
type Fahrenheit float64

// Method on Celsius — note: Celsius is just a float64 underneath, but methods
// give it behavior like a Python class.
func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9.0/5.0 + 32.0)
}

func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - 32.0) * 5.0 / 9.0)
}

// Compare with Python:
//   This is hard to do cleanly in Python — you'd either:
//   (a) make a class with one field: class Celsius: def __init__(self, v): self.v = v
//   (b) use standalone functions: def to_fahrenheit(c): return c * 9/5 + 32
//
// Go lets you keep the primitive efficiency while adding behavior through
// methods. Best of both worlds.

// You can even attach methods to named slices:
type Point struct {
	X, Y float64
}

type Polygon []Point

// Area calculates the area of a polygon using the shoelace formula.
// In Python, you'd need a class wrapping a list:
//   class Polygon:
//       def __init__(self, points):
//           self.points = points
//       def area(self):
//           ...
func (p Polygon) Area() float64 {
	n := len(p)
	if n < 3 {
		return 0
	}
	var area float64
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += p[i].X*p[j].Y - p[j].X*p[i].Y
	}
	return math.Abs(area) / 2.0
}

// =============================================================================
// 5. METHOD EMBEDDING — Go's Alternative to Inheritance
// =============================================================================
//
// Go has NO inheritance and NO `class` hierarchy. Instead, it uses
// COMPOSITION through struct embedding. This is Go's answer to the
// "favor composition over inheritance" principle that Python supports
// but doesn't enforce.

type Shape struct {
	Color string
}

func (s Shape) Describe() string {
	return fmt.Sprintf("a %s shape", s.Color)
}

// Circle EMBEDS Shape — no "extends", no "inherits", no "derives".
// This is composition, not inheritance.
type Circle struct {
	Shape        // Embedded field (no name = embedding)
	Radius float64
}

// Area — Circle's own method. Python would use `def area(self)` inside
// a class that inherits from Shape.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Key insight: Circle "inherits" Shape's methods through embedding.
// circle.Describe() calls Shape.Describe() automatically.
// This is called PROMOTION — methods of the embedded type are promoted
// to the outer type.

// If Circle does NOT define Describe(), circle.Describe() will call
// Shape.Describe(). If Circle DOES define its own Describe(), it
// SHADOWS the embedded type's method.

// Override example — like Python's super():
func (c Circle) Describe() string {
	// Call the embedded type's method explicitly (like super() in Python)
	baseDesc := c.Shape.Describe()
	return fmt.Sprintf("%s circle with radius %.2f", baseDesc, c.Radius)
}

// ┌──────────────────────────────────────────────────────────────────────────┐
// │ Python comparison:                                                      │
// │                                                                         │
// │ class Shape:                       # Go: type Shape struct              │
// │     def __init__(self, color):     # (no constructor, just fields)      │
// │         self.color = color                                              │
// │     def describe(self):                                                 │
// │         return f"a {self.color} shape"                                  │
// │                                                                         │
// │ class Circle(Shape):               # Go: type Circle struct { Shape }   │
// │     def __init__(self, color, radius):                                  │
// │         super().__init__(color)     # Go: no super() needed; fields     │
// │         self.radius = radius        #      are promoted automatically   │
// │                                                                         │
// │     def describe(self):                                                 │
// │         base = super().describe()   # Go: c.Shape.Describe()            │
// │         return f"{base} circle with radius {self.radius}"               │
// │                                                                         │
// │ BIG DIFFERENCE: Python has method resolution order (MRO) and supports   │
// │ diamond inheritance. Go has NO MRO — only unambiguous promotion. If     │
// │ two embedded types have the same method, it's a COMPILE ERROR unless   │
// │ you disambiguate explicitly. Go chooses safety over flexibility here.   │
// └──────────────────────────────────────────────────────────────────────────┘

// =============================================================================
// 6. METHODS ON NIL RECEIVERS
// =============================================================================
//
// In Python, calling a method on None raises AttributeError.
// In Go, you CAN call a method on a nil receiver, and it's often used
// intentionally. This is Go's version of "null object pattern."

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// Sum returns the sum of all values in the linked list.
// This works even if ll is nil — summing an empty list returns 0.
// Python:
//   def sum(self):
//       if self is None: return 0    # ← you'd never write this in Python
//       return self.value + (self.next.sum() if self.next else 0)
func (ll *LinkedList) Sum() int {
	if ll == nil {
		return 0
	}
	return ll.Value + ll.Next.Sum()
}

// =============================================================================
// 7. METHODS vs FUNCTIONS — Explicit Choice
// =============================================================================
//
// In Go, you can always write a function instead of a method:
//
//   func Area(r Rect) float64  { return r.Width * r.Height }   // function
//   func (r Rect) Area()         float64 { return r.Width * r.Height }    // method
//
// Both work. Which do you choose?
//
//   Use a METHOD when:         Use a FUNCTION when:
//   - The operation "belongs   - The operation works equally on
//     to" the type (is-a)        multiple types (operates on)
//   - Multiple operations       - The first argument doesn't
//     on the same data            clearly suggest the type
//     (cohesion)                 - You need to pass behavior
//   - Interface satisfaction      (function values, callbacks)
//     (see section 8)
//
// In Python, this choice doesn't exist — a function inside a class body
// always becomes a method (or a @staticmethod/@classmethod with decorators).
// Go is more explicit about the distinction.

// =============================================================================
// 8. INTERFACE SATISFACTION — Go's Duck Typing
// =============================================================================
//
// This is where Go methods really shine. Go has INTERFACES, which are
// satisfied IMPLICITLY — a type satisfies an interface just by having
// the required methods. No "implements" keyword (unlike Java or Python's
// ABC/Protocol).

type Shape2D interface {
	Area() float64
	Perimeter() float64
}

// Give Rect a Perimeter method
func (r Rect) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Give Circle a Perimeter method
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Now both Rect and Circle satisfy Shape2D automatically.
// Python equivalent:
//   from typing import Protocol
//   class Shape2D(Protocol):
//       def area(self) -> float: ...
//       def perimeter(self) -> float: ...
//
// But Python's structural subtyping (Protocol) is optional and opt-in.
// Go's is the PRIMARY mechanism for polymorphism.

func PrintShapeInfo(s Shape2D) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
	// In Python, this would work the same via duck typing, but Go checks
	// at COMPILE TIME that s satisfies Shape2D. Python checks at RUNTIME.
	// Go catches missing methods before the program ever runs.
}

// =============================================================================
// 9. STRING() METHOD — Go's __str__ / __repr__
// =============================================================================

// Go's fmt package looks for the Stringer interface (like Python's __str__):
//
//   type Stringer interface {
//       String() string
//   }

func (r Rect) String() string {
	return fmt.Sprintf("Rect(%.1f x %.1f)", r.Width, r.Height)
	// Python: def __str__(self): return f"Rect({self.width} x {self.height})"
	//
	// Note: Go has ONLY String() — no separate __str__ vs __repr__.
	// fmt.Printf("%v") and fmt.Printf("%+v") can differ in VERBOSITY
	// but both call String() by default.
}

func (c Celsius) String() string {
	return fmt.Sprintf("%.1f°C", float64(c))
	// Python: def __str__(self): return f"{self.value}°C"
}

// =============================================================================
// 10. COMPLETE COMPARISON TABLE
// =============================================================================
//
// ┌──────────────────────────────┬────────────────────────────────────┬──────────────────────────────────────┐
// │          Feature             │            Go                     │              Python                   │
// ├──────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ How to define a method      │ func (r Rect) Area() float64       │ class Rect:                           │
// │                              │                                    │     def area(self):                   │
// ├──────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Receiver / self              │ Separate syntax slot — not a      │ Always the first parameter —          │
// │                              │ regular parameter                 │ named `self` by convention            │
// ├──────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Value vs reference          │ You CHOOSE: value (copy) or       │ Always reference — no value           │
// │                              │ pointer (reference) receiver      │ receiver concept exists               │
// ├──────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ What can have methods       │ ANY named type — structs,          │ Only class instances (and             │
// │                              │ ints, strings, slices, etc.      │ metaclass instances)                  │
// ├──────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Inheritance                  │ NONE — composition via embedding   │ Class-based inheritance, MRO,         │
// │                              │                                    │ super(), multiple inheritance         │
// ├──────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Polymorphism                 │ Interfaces (implicit satisfaction) │ Duck typing, ABC, Protocol            │
// ├──────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Constructor / __init__       │ No built-in constructors — use    │ __init__ is called automatically      │
// │                              │ factory functions or struct       │ when you do Rect(...)                 │
// │                              │ literals                          │                                      │
// ├──────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Operator overloading         │ NOT supported (by design)         │ Supported via __add__, __eq__, etc.   │
// ├──────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Static methods / @staticmethod│ No direct equivalent — use       │ @staticmethod decorator               │
// │                              │ package-level functions instead  │                                      │
// ├──────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Class methods / @classmethod │ No equivalent — no class-level    │ @classmethod decorator, receives cls  │
// │                              │ state outside global/package      │                                      │
// │                              │ variables                         │                                      │
// ├──────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ __str__ / __repr__           │ String() method                   │ __str__ and __repr__                  │
// ├──────────────────────────────┼────────────────────────────────────┼──────────────────────────────────────┤
// │ Nil / None receivers        │ Methods CAN be called on nil       │ Calling a method on None raises       │
// │                              │ receivers — common pattern        │ AttributeError                       │
// └──────────────────────────────┴────────────────────────────────────┴──────────────────────────────────────┘

// =============================================================================
// DEMO — Putting it all together
// =============================================================================

func main() {
	// --- Value receiver demo ---
	r := Rect{Width: 10, Height: 5}
	fmt.Println("Original:", r)
	fmt.Printf("Area: %.2f\n", r.Area())

	bigger := r.Scale(2.0)
	fmt.Println("Scaled copy:", bigger)
	fmt.Println("Original unchanged:", r)

	// --- Pointer receiver demo ---
	fmt.Println("\n--- Pointer Receiver ---")
	r.ScaleInPlace(2.0)
	fmt.Println("After ScaleInPlace:", r)

	// --- Methods on non-struct types ---
	fmt.Println("\n--- Methods on Custom Types ---")
	boiling := Celsius(100.0)
	fmt.Println("Boiling:", boiling)
	fmt.Printf("Boiling in F: %.1f\n", boiling.ToFahrenheit())

	// --- Embedded methods ---
	fmt.Println("\n--- Embedding (Composition) ---")
	c := Circle{
		Shape:  Shape{Color: "red"},
		Radius: 5.0,
	}
	fmt.Println(c.Describe()) // Shadowed method — calls Circle's Describe
	fmt.Printf("Area: %.2f\n", c.Area())
	// Promoted method:
	fmt.Println("Shape color:", c.Color) // promoted field, like c.Shape.Color

	// --- Nil receiver ---
	fmt.Println("\n--- Nil Receiver ---")
	var emptyList *LinkedList
	fmt.Printf("Sum of nil list: %d (no panic!)\n", emptyList.Sum())
	list := &LinkedList{Value: 1, Next: &LinkedList{Value: 2, Next: &LinkedList{Value: 3}}}
	fmt.Printf("Sum of [1,2,3]: %d\n", list.Sum())

	// --- Interface satisfaction ---
	fmt.Println("\n--- Interface Satisfaction (Compile-time Duck Typing) ---")
	PrintShapeInfo(r)
	PrintShapeInfo(c)

	// --- Polygon ---
	fmt.Println("\n--- Methods on Slice Types ---")
	triangle := Polygon{{X: 0, Y: 0}, {X: 4, Y: 0}, {X: 0, Y: 3}}
	fmt.Printf("Triangle area: %.2f (should be 6.00)\n", triangle.Area())

	// --- Stringer ---
	fmt.Println("\n--- String() like __str__ ---")
	fmt.Println("Using %%v:", r)
	fmt.Println("Using %%v:", Celsius(25))
}

// =============================================================================
// SUMMARY: Python → Go Method Translation Cheat Sheet
// =============================================================================
//
// Python:                              Go:
// ───────                              ──
// class Dog:                           type Dog struct {
//     def __init__(self, name):                 Name string
//         self.name = name             }
//
//     def bark(self, times):           func (d Dog) Bark(times int) string {
//         return self.name * times           return strings.Repeat(d.Name+"! ", times)
//                                       }
//
//     @staticmethod                   package-level function:
//     def species():                  func Species() string {
//         return "Canis"                     return "Canis"
//                                       }
//
//     def __str__(self):              func (d Dog) String() string {
//         return self.name                   return d.Name
//                                       }
//
// ─── Key takeaways ───
//
// 1. Go methods are MORE EXPLICIT — you choose value vs pointer receiver.
//    Python always uses references (like Go's pointer receiver).
//
// 2. Go methods use a SEPARATE receiver slot, not a regular parameter.
//    Python puts self as the first parameter.
//
// 3. Go has NO inheritance, but embedding often achieves the same goals
//    more safely. Python has full class inheritance with MRO.
//
// 4. Go lets you attach methods to ANY type, not just structs/classes.
//    This is more flexible than Python.
//
// 5. Go interfaces are satisfied IMPLICITLY — no "implements" keyword.
//    Python's Protocol (PEP 544) is the closest equivalent.
//
// 6. Go methods CAN be called on nil receivers — Python methods on None
//    will always panic/error. This is a deliberate design choice in Go.
//