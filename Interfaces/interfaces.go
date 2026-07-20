// ============================================================
//  INTERFACES IN GO — with Python Comparisons
// ============================================================
//  Interfaces define BEHAVIOR (method sets), not data.
//  They are Go's primary tool for polymorphism and abstraction.
//
// ┌─────────────────────────────┬────────────────────────────────────────────┐
// │        Go                   │                  Python                    │
// ├─────────────────────────────┼────────────────────────────────────────────┤
// │ type Animal interface {     │ from abc import ABC, abstractmethod        │
// │     Speak() string          │ class Animal(ABC):                        │
// │ }                           │     @abstractmethod                       │
// │                             │     def speak(self) -> str: ...            │
// │                             │                                            │
// │                             │ # OR (structural typing):                  │
// │                             │ from typing import Protocol                │
// │                             │ class Animal(Protocol):                   │
// │                             │     def speak(self) -> str: ...            │
// ├─────────────────────────────┼────────────────────────────────────────────┤
// │ IMPLICIT satisfaction       │ Protocol = structural (like Go)           │
// │ — NO "implements" keyword   │ ABC = nominal (explicit registration)     │
// │ — just HAS the methods      │                                            │
// │                             │                                            │
// │ type satisfies interface    │ isinstance(x, Animal) checks at           │
// │ at COMPILE TIME             │ RUNTIME. Go checks at COMPILE TIME.      │
// └─────────────────────────────┴────────────────────────────────────────────┘
//
// KEY DIFFERENCES:
//   1. IMPLICIT: Go types satisfy interfaces automatically (duck typing).
//      Python's ABC requires explicit registration; Protocol is structural.
//   2. STATIC: Go checks interface satisfaction at COMPILE TIME.
//      Python checks at RUNTIME (or with mypy separately).
//   3. NOMINAL vs STRUCTURAL: Go interfaces are STRUCTURAL (based on methods).
//      Python has both: ABC (nominal) and Protocol (structural).
//   4. SIZE: A Go interface is 2 words (pointer to type + pointer to value).
//      Python's object always carries its full class hierarchy.
// ============================================================

package main

import (
	"fmt"
	"math"
)

// =============================================================================
// 1. DEFINING AN INTERFACE
// =============================================================================
// An interface is a set of method signatures.
// Any type that implements ALL these methods satisfies the interface.
//
// Python (Protocol, structural typing):
//   from typing import Protocol, runtime_checkable
//
//   @runtime_checkable
//   class Speaker(Protocol):
//       def speak(self) -> str: ...
//
// Python (ABC, nominal typing):
//   from abc import ABC, abstractmethod
//
//   class Speaker(ABC):
//       @abstractmethod
//       def speak(self) -> str: ...
//
// KEY: Go has NO 'implements' keyword. Satisfaction is AUTOMATIC.
// Python's Protocol is closest. ABC requires explicit registration.

type Speaker interface {
	Speak() string
}

// =============================================================================
// 2. IMPLEMENTING AN INTERFACE (Implicitly)
// =============================================================================
// No "implements Speaker" needed — just having Speak() is enough.
//
// Python:
//   class Dog:
//       def speak(self) -> str:
//           return "Woof!"
//   # Dog automatically satisfies Speaker Protocol (structural typing)

type Dog struct {
	Name string
}

// Dog implements Speaker by defining Speak().
// There is NO explicit declaration of intent — it just fits.
func (d Dog) Speak() string {
	return fmt.Sprintf("%s says: Woof!", d.Name)
}

// Cat also implements Speaker — different struct, same interface.
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return fmt.Sprintf("%s says: Meow!", c.Name)
}

// =============================================================================
// 3. USING INTERFACES
// =============================================================================
// Functions can accept interfaces, making them polymorphic.
//
// Python:
//   def make_it_speak(s: Speaker) -> None:
//       print(s.speak())
//
// In Python, this is duck typing at RUNTIME.
// In Go, the compiler verifies 's' has Speak() at COMPILE TIME.
func makeItSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

// =============================================================================
// 4. THE EMPTY INTERFACE — interface{} / any
// =============================================================================
// An interface with ZERO methods is satisfied by EVERY type.
// This is Go's equivalent of Python's object or typing.Any.
//
// Python:
//   def print_value(v: object) -> None:   # or Any
//       print(v)
//
// In Go 1.18+, 'any' is an alias for 'interface{}'.
//
// ⚠️ Use SPARINGLY — it bypasses type safety. Prefer specific interfaces.

func printValue(v any) {
	fmt.Printf("Value: %v, Type: %T\n", v, v)
}

// =============================================================================
// 5. TYPE ASSERTION — Extracting the concrete type
// =============================================================================
// To get back to the concrete type, use a TYPE ASSERTION.
// This is like a checked downcast.
//
// Syntax:   value, ok := interfaceVar.(ConcreteType)
//
// If the assertion fails, ok is false (no panic).
// Without the second value, a failed assertion PANICS.
//
// Python:
//   if isinstance(v, Dog):
//       dog = v
//       dog.speak()
//   # Python's isinstance() is a runtime type check.
//   # Go's type assertion is also runtime, but with a DIFFERENT
//   # syntax and safety mechanism (comma-ok idiom vs try/except).

func describeSpeaker(s Speaker) {
	// Type assertion to Dog
	if dog, ok := s.(Dog); ok {
		fmt.Printf("  It's a dog named %s!\n", dog.Name)
	} else if cat, ok := s.(Cat); ok {
		fmt.Printf("  It's a cat named %s!\n", cat.Name)
	} else {
		fmt.Println("  Unknown animal")
	}
}

// =============================================================================
// 6. TYPE SWITCH — Clean multi-type dispatch
// =============================================================================
// A type switch is a switch statement that matches on TYPES, not values.
// This is Go's clean version of Python's match/isinstance chain.
//
// Python:
//   match v:
//       case Dog():
//           print(f"Dog: {v.speak()}")
//       case Cat():
//           print(f"Cat: {v.speak()}")
//       case _:
//           print("Unknown")
//
// Go's type switch is MORE CONCISE than an if/else chain of type assertions.

func classify(s any) {
	switch v := s.(type) { // ← special syntax: .(type) in a switch
	case Dog:
		fmt.Printf("Dog: %s (happy to see you!)\n", v.Name)
	case Cat:
		fmt.Printf("Cat: %s (judging you silently)\n", v.Name)
	case string:
		fmt.Printf("Just a string: %s\n", v)
	case int:
		fmt.Printf("Just a number: %d\n", v)
	case nil:
		fmt.Println("Nothing at all!")
	default:
		fmt.Printf("Something else: %T\n", v)
	}
}

// =============================================================================
// 7. INTERFACE EMBEDDING — Composition for interfaces
// =============================================================================
// Just like structs, interfaces can EMBED other interfaces.
// The resulting interface requires ALL methods of the embedded interfaces.
//
// This is Go's alternative to interface inheritance.
// Python has no direct equivalent — you'd use multiple inheritance or
// compose ABCs.

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

// ReadWriter embeds both Reader and Writer.
// A type must have BOTH Read() and Write() to satisfy it.
type ReadWriter interface {
	Reader
	Writer
}

// A concrete type that satisfies ReadWriter
type File struct {
	content string
}

func (f *File) Read() string {
	return f.content
}

func (f *File) Write(data string) {
	f.content = data
}

// =============================================================================
// 8. STANDARD INTERFACES — Go's Built-in Contracts
// =============================================================================
// Go's standard library is built around small, focused interfaces.
// This is the "interface tax" philosophy — prefer many small interfaces
// over one large one.
//
// ┌────────────┬────────────────────┬──────────────────────────────────────┐
// │ Interface  │ Method(s)          │ Python equivalent                    │
// ├────────────┼────────────────────┼──────────────────────────────────────┤
// │ fmt.Stringer│ String() string   │ __str__() / __repr__()              │
// │ error      │ Error() string    │ BaseException (but very different!)  │
// │ io.Reader  │ Read(p []byte)    │ file.read(size)                      │
// │ io.Writer  │ Write(p []byte)   │ file.write(data)                     │
// │ io.Closer  │ Close() error     │ file.close()                         │
// │ comparable │ ==, !=            │ __eq__ / __ne__                      │
// └────────────┴────────────────────┴──────────────────────────────────────┘

// Implementing fmt.Stringer (like Python's __str__)
func (d Dog) String() string {
	return fmt.Sprintf("Dog(%s)", d.Name)
}

func (c Cat) String() string {
	return fmt.Sprintf("Cat(%s)", c.Name)
}

// Implementing error (like a Python exception type)
type ValidationError struct {
	Field string
	Value any
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed: %s = %v", e.Field, e.Value)
}

// Python equivalent:
//   class ValidationError(Exception):
//       def __init__(self, field, value):
//           self.field = field
//           self.value = value
//       def __str__(self):
//           return f"validation failed: {self.field} = {self.value}"
//
// KEY DIFFERENCE: Go's error is an INTERFACE, not a class hierarchy.
// ANY type with Error() string is an error, including primitives.
// Python exceptions are a CLASS HIERARCHY rooted in BaseException.

// =============================================================================
// 9. VALUE vs POINTER RECEIVERS with Interfaces
// =============================================================================
// ⚠️ IMPORTANT: If a method is defined on a POINTER receiver (*T),
//    then only *T satisfies the interface, not T.
//
// If a method is defined on a VALUE receiver (T),
// then BOTH T and *T satisfy the interface.
//
// Python has no equivalent concept — self is always a reference.

type Counter interface {
	Increment()
	Value() int
}

type IntCounter struct {
	val int
}

// Value receiver — both IntCounter and *IntCounter satisfy Counter
func (c IntCounter) Value() int {
	return c.val
}

// Pointer receiver — ONLY *IntCounter satisfies Counter
func (c *IntCounter) Increment() {
	c.val++
}

// =============================================================================
// 10. COMPLETE COMPARISON: Go Interfaces vs Python
// =============================================================================
//
// ┌──────────────────────────┬──────────────────────────┬──────────────────────────────┐
// │ Feature                  │ Go                       │ Python                       │
// ├──────────────────────────┼──────────────────────────┼──────────────────────────────┤
// │ Declaration              │ type X interface { ... } │ class X(Protocol): / ABC     │
// │ Satisfaction             │ IMPLICIT (just have the  │ explicit extends ABC or      │
// │                          │ methods)                 │ structural via Protocol      │
// │ Check timing             │ COMPILE TIME             │ RUNTIME (or mypy separately) │
// │ "implements" keyword     │ NONE                     │ None (ABC.register())        │
// │ Empty interface          │ any / interface{}        │ object / Any                │
// │ Type assertion           │ v.(Type) with comma-ok  │ isinstance() / match-case   │
// │ Type switch              │ switch v.(type) { }     │ match v: case X:            │
// │ Embedding                │ Embedded interfaces      │ Multiple inheritance / ABCs │
// │ Size in memory           │ 2 words (ptr + type)     │ Full object header          │
// │ Nil interface            │ nil (no value, no type)  │ None                        │
// │ Methods on interface     │ Method set only          │ Can have default impls (ABC)│
// │ Generics                 │ Go 1.18+ (type params)   │ Fully supported (always)    │
// │ Typical interface size   │ 1-3 methods (small!)     │ Can be any size             │
// │ Idiomatic usage          │ Accept interfaces,       │ Duck typing everywhere      │
// │                          │ return structs           │                             │
// └──────────────────────────┴──────────────────────────┴──────────────────────────────┘

// =============================================================================
// DEMO — Putting it all together
// =============================================================================

func main() {
	// --- Basic interface usage ---
	fmt.Println("=== Basic Interface Usage ===")
	dog := Dog{Name: "Rover"}
	cat := Cat{Name: "Whiskers"}

	makeItSpeak(dog) // Dog implements Speaker
	makeItSpeak(cat) // Cat implements Speaker

	// Speaker slice — polymorphic!
	fmt.Println("\n=== Polymorphic Slice ===")
	animals := []Speaker{dog, cat}
	for _, a := range animals {
		fmt.Println(" ", a.Speak())
	}

	// --- Empty interface ---
	fmt.Println("\n=== Empty Interface (any) ===")
	printValue(42)
	printValue("hello")
	printValue(dog)
	printValue(3.14)

	// --- Type assertion ---
	fmt.Println("\n=== Type Assertion ===")
	describeSpeaker(dog)
	describeSpeaker(cat)

	// --- Type switch ---
	fmt.Println("\n=== Type Switch ===")
	classify(dog)
	classify(cat)
	classify("hello")
	classify(42)
	classify(nil)

	// --- Interface embedding ---
	fmt.Println("\n=== Interface Embedding ===")
	file := &File{}
	file.Write("Hello, World!")
	fmt.Println("Read from file:", file.Read())

	// Using ReadWriter
	var rw ReadWriter = file
	fmt.Println("Read via ReadWriter:", rw.Read())

	// --- Stringer ---
	fmt.Println("\n=== Stringer (like __str__) ===")
	fmt.Println("Dog stringer:", dog) // calls Dog.String()
	fmt.Println("Cat stringer:", cat) // calls Cat.String()

	// --- Error interface ---
	fmt.Println("\n=== Error Interface ===")
	err := ValidationError{Field: "age", Value: -5}
	fmt.Println("Error:", err)
	// Errors can be passed around and checked by type:
	if valErr, ok := err.(ValidationError); ok {
		fmt.Printf("  Field=%s, Value=%v\n", valErr.Field, valErr.Value)
	}

	// --- Pointer vs Value receiver with interfaces ---
	fmt.Println("\n=== Pointer vs Value Receiver ===")
	var ctr Counter = &IntCounter{val: 0}
	// Note: must use &IntCounter because Increment() is on *IntCounter
	// IntCounter{} alone would NOT compile here!
	ctr.Increment()
	ctr.Increment()
	fmt.Println("Counter value:", ctr.Value())

	// ============================================================
	//  SUMMARY: When to Use Interfaces
	// ============================================================
	//  ┌─────────────────────────────────────────────────────────────────────┐
	//  │ 1. Keep interfaces SMALL — 1-3 methods is ideal                    │
	//  │ 2. Name interfaces after their method: Stringer, Reader, Writer    │
	//  │ 3. "Accept interfaces, return structs" — decouple callers from     │
	//  │    implementation details                                          │
	//  │ 4. Use type assertions/type switches SPARINGLY — prefer            │
	//  │    interface methods instead                                       │
	//  │ 5. Prefer many small interfaces over one large one (Interface Seg- │
	//  │    regation Principle)                                             │
	//  │ 6. The empty interface (any) is a LAST RESORT, not a default       │
	//  └─────────────────────────────────────────────────────────────────────┘
	//
	// Bottom line: Go interfaces are SMALL, STATIC, and IMPLICIT.
	// Python's equivalent is Protocol (structural typing) or ABC (explicit).
	// Go catches interface errors at COMPILE TIME; Python catches them at
	// RUNTIME. This is Go's biggest safety advantage over Python for
	// large codebases.
	// ============================================================
}
