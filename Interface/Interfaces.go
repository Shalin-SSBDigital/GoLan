package main

import "fmt"

// =============================================
// 1. Basic Interface Definition & Implementation
// =============================================

// Speaker defines a contract: any type that has a Speak() method
// that returns a string is a Speaker.
type Speaker interface {
	Speak() string
}

// Dog implements the Speaker interface.
type Dog struct {
	Name string
}

// Speak satisfies the Speaker interface.
func (d Dog) Speak() string {
	return d.Name + " says Woof!"
}

// Cat implements the Speaker interface.
type Cat struct {
	Name string
}

// Speak satisfies the Speaker interface.
func (c Cat) Speak() string {
	return c.Name + " says Meow!"
}

// =============================================
// 2. Interface as Function Parameter
// =============================================

// MakeNoise accepts any Speaker and calls its Speak method.
// This is the power of interfaces — we can pass in any type
// that satisfies the Speaker contract.
func MakeNoise(s Speaker) {
	fmt.Println(s.Speak())
}

// =============================================
// 3. Multiple Interfaces
// =============================================

// Mover defines a contract for moving.
type Mover interface {
	Move() string
}

// Robot implements both Speaker and Mover.
type Robot struct {
	Model string
}

func (r Robot) Speak() string {
	return r.Model + " says: I am a robot."
}

func (r Robot) Move() string {
	return r.Model + " is moving."
}

// =============================================
// 4. Interface Composition (Embedding)
// =============================================

// MovingSpeaker composes the Speaker and Mover interfaces.
// Any type that satisfies both Speaker and Mover
// automatically satisfies MovingSpeaker.
type MovingSpeaker interface {
	Speaker
	Mover
}

// Activate takes a MovingSpeaker and uses both capabilities.
func Activate(ms MovingSpeaker) {
	fmt.Println(ms.Speak())
	fmt.Println(ms.Move())
}

// =============================================
// 5. Empty Interface
// =============================================

// DescribeValue takes an empty interface (any type) and
// uses a type switch to handle different types.
func DescribeValue(v any) {
	switch val := v.(type) {
	case int:
		fmt.Println("Integer:", val)
	case string:
		fmt.Println("String:", val)
	case bool:
		fmt.Println("Boolean:", val)
	default:
		fmt.Println("Unknown type:", val)
	}
}

// =============================================
// 6. Type Assertion
// =============================================

// TrySpeak attempts to treat an interface value as a Speaker.
func TrySpeak(v any) {
	// Type assertion: try to extract the Speaker from v.
	s, ok := v.(Speaker)
	if ok {
		fmt.Println(s.Speak())
	} else {
		fmt.Println("Value does not implement Speaker")
	}
}

// =============================================
// 7. Interface Slice (Polymorphism)
// =============================================

// AnimalShow puts multiple speakers in a slice and iterates.
func AnimalShow(speakers []Speaker) {
	fmt.Println("--- Animal Show ---")
	for _, s := range speakers {
		fmt.Println(s.Speak())
	}
}

// =============================================
// 8. Interface with Pointer Receiver
// =============================================

// Counter shows the difference between value and pointer receivers
// when implementing interfaces.
type Counter struct {
	Value int
}

// Increment uses a pointer receiver to mutate the Counter.
func (c *Counter) Increment() {
	c.Value++
}

// Incrementer defines an interface with a pointer-receiver method.
type Incrementer interface {
	Increment()
}

// =============================================
// main
// =============================================

func main() {
	fmt.Println("=== Go Interfaces ===")
	fmt.Println()

	// --- Basic interface usage ---
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	MakeNoise(dog)
	MakeNoise(cat)

	fmt.Println()

	// --- Multiple interfaces ---
	robot := Robot{Model: "R2-D2"}
	MakeNoise(robot)        // robot satisfies Speaker
	fmt.Println(robot.Move()) // robot also satisfies Mover

	fmt.Println()

	// --- Interface composition ---
	Activate(robot)

	fmt.Println()

	// --- Empty interface & type switch ---
	DescribeValue(42)
	DescribeValue("hello")
	DescribeValue(true)
	DescribeValue(3.14)

	fmt.Println()

	// --- Type assertion ---
	TrySpeak(dog)
	TrySpeak(42) // not a Speaker

	fmt.Println()

	// --- Interface slice (polymorphism) ---
	speakers := []Speaker{dog, cat, robot}
	AnimalShow(speakers)

	fmt.Println()

	// --- Pointer receiver & interface ---
	counter := Counter{Value: 0}
	var inc Incrementer = &counter // *Counter satisfies Incrementer
	// NOTE: Counter (value) does NOT satisfy Incrementer because
	// Increment has a pointer receiver. Only *Counter works.

	inc.Increment()
	inc.Increment()
	inc.Increment()
	fmt.Println("Counter value after 3 increments:", counter.Value) // 3
}
