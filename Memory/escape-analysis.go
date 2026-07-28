// ============================================================
//  ESCAPE ANALYSIS IN GO
//  ============================================================
//  Go's compiler decides: STACK or HEAP?
//
//  Run this file with escape analysis output:
//    go build -gcflags="-m" escape-analysis.go
//
//  The compiler prints lines like:
//    ./escape-analysis.go:XX:YY: x escapes to heap
//    ./escape-analysis.go:XX:YY: y does not escape
// ============================================================
//
//  ┌────────────┬────────────────────────────────────────────────┐
//  │    Go      │                 Python                         │
//  ├────────────┼────────────────────────────────────────────────┤
//  │ Escape     │ Python has NO escape analysis.                 │
//  │ Analysis:  │ EVERYTHING is heap-allocated.                  │
//  │            │ Every int, every string, every object.         │
//  │ Go's       │                                                │
//  │ compiler   │ CPython uses reference counting + GC.          │
//  │ decides    │ There is no "stack vs heap" decision —         │
//  │ where      │ Python doesn't offer it.                       │
//  │ values go. │                                                │
//  │ It's      │                                                │
//  │ AUTOMATIC. │                                                │
//  └────────────┴────────────────────────────────────────────────┘
// ============================================================

package main

import "fmt"

// ============================================================
//  RULE 1: Taking an address causes escape
//  If &x is used after the function returns, x must be on heap.
//  This is the most common escape reason.
// ============================================================

// staysOnStack — value never escapes
// Compiler says: "does not escape"
func staysOnStack() int {
	x := 42        // stays on stack
	return x       // returned BY VALUE, copy is fine
}

// escapesToHeap — taking address forces heap
// Compiler says: "moves to heap: x"
func escapesToHeap() *int {
	x := 42        // compiler sees &x returned → moves to HEAP
	return &x      // address returned → x must outlive this function
}

// ============================================================
//  RULE 2: Passing to interface{} causes escape
//  When you pass a concrete type to an interface parameter,
//  Go must box it (wrap in interface value on heap).
//  Compiler says: "escapes to heap" for the argument.
// ============================================================

// printAny takes interface{} — the value gets boxed on heap
func printAny(v interface{}) {
	fmt.Println(v)
}

// ============================================================
//  RULE 3: Global variables cause escape
//  If a local value is assigned to a package-level variable,
//  it must escape because the global lives forever.
// ============================================================

var global *int // package-level pointer

func setGlobal() {
	x := 99
	global = &x  // x escapes! global holds the address forever
}

// ============================================================
//  RULE 4: Slices always allocate backing array on heap
//  The slice header (ptr, len, cap) may be on stack, but
//  the actual data array is always on heap.
//  Compiler says: "make([]int, n) escapes to heap"
// ============================================================

func sliceOnStack() {
	// Small fixed-size array — stays on stack
	var arr [3]int
	arr[0] = 1
	fmt.Println(arr)
}

func sliceOnHeap() {
	// Slice — backing array on heap
	sl := make([]int, 3)
	sl[0] = 1
	fmt.Println(sl)
}

// ============================================================
//  RULE 5: Returning a slice escapes the backing array
//  If the slice is returned, the backing array must live on.
// ============================================================

func returnSlice() []int {
	nums := make([]int, 1000) // backing array on heap
	return nums               // fine — it was already on heap
}

// ============================================================
//  RULE 6: String concatenation with + may escape
//  Strings are immutable. Concatenation allocates new string.
// ============================================================

func buildString() string {
	a := "Hello"
	b := "World"
	return a + " " + b // new string allocated
}

// ============================================================
//  RULE 7: fmt.Printf causes escapes
//  fmt functions take ...interface{}, so arguments escape.
//  Even simple printing causes escape!
//  Compiler says: "arg escapes to heap"
// ============================================================

func printExample() {
	x := 42
	fmt.Printf("x = %d\n", x) // x escapes!
}

// ============================================================
//  RULE 8: Closure captures escape
//  If a closure references a local variable, that variable
//  escapes to heap because the closure may outlive the function.
// ============================================================

func closureEscapes() func() int {
	x := 0 // x escapes — the closure references it
	return func() int {
		x++
		return x
	}
}

// ============================================================
//  RULE 9: Large values may be heap-allocated
//  Values > ~64 bytes may be moved to heap to avoid
//  stack overflow. This is compiler-dependent.
// ============================================================

type SmallStruct struct {
	A, B, C, D int   // 32 bytes — usually stack
}

type LargeStruct struct {
	Data [100]int     // 800 bytes — usually heap
}

func newLargeStruct() LargeStruct {
	return LargeStruct{} // large = may escape to heap
}

// ============================================================
//  RULE 10: Indirect escape — through function parameters
//  If you call a function that stores a pointer in a global,
//  your value escapes even if you don't know it.
// ============================================================

func storeInGlobal(p *int) {
	global = p // stores pointer — p's value escapes
}

func caller() {
	x := 100
	storeInGlobal(&x) // x escapes because storeInGlobal saves it
}

func main() {
	fmt.Println("========================================")
	fmt.Println("  🔍  ESCAPE ANALYSIS DEEP DIVE 🔍")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("  To see escape analysis output, run:")
	fmt.Println("    go build -gcflags=\"-m\" escape-analysis.go")
	fmt.Println()
	fmt.Println("  The -m flag prints compiler decisions.")
	fmt.Println("  More -m flags = more detail:")
	fmt.Println("    -gcflags=\"-m -m -m\"")
	fmt.Println()

	// ── Demonstrate each scenario ──
	fmt.Println("─── Scenario 1: Return by value (stack) ───")
	stackVal := staysOnStack()
	fmt.Printf("  Value: %d — returned by copy, stack safe\n", stackVal)
	fmt.Println()

	fmt.Println("─── Scenario 2: Return pointer (heap escape) ───")
	heapPtr := escapesToHeap()
	fmt.Printf("  Value: %d — pointer returned, heap allocated\n", *heapPtr)
	fmt.Println()

	fmt.Println("─── Scenario 3: interface{} parameter ───")
	printAny(42) // 42 escapes to heap
	fmt.Println("  (42 was boxed into interface{} on heap)")
	fmt.Println()

	fmt.Println("─── Scenario 4: Global variable ───")
	setGlobal()
	fmt.Printf("  Global holds: %d — escaped to heap\n", *global)
	fmt.Println()

	fmt.Println("─── Scenario 5: Array vs Slice ───")
	sliceOnStack()
	sliceOnHeap()
	fmt.Println("  Array: [3]int stays on stack")
	fmt.Println("  Slice: make([]int, 3) data on heap")
	fmt.Println()

	fmt.Println("─── Scenario 6: Closure capture ───")
	counter := closureEscapes()
	fmt.Println("  Closure captures x — x escapes to heap")
	fmt.Printf("  Count: %d\n", counter())
	fmt.Printf("  Count: %d\n", counter())
	fmt.Printf("  Count: %d\n", counter())
	fmt.Println()

	fmt.Println("─── Scenario 7: fmt.Printf escapes ───")
	printExample()
	fmt.Println("  Every fmt call causes escape — unavoidable")
	fmt.Println()

	fmt.Println("========================================")
	fmt.Println("  RUN ESCAPE ANALYSIS:")
	fmt.Println("  go build -gcflags=\"-m\" escape-analysis.go")
	fmt.Println("========================================")
}

// ============================================================
//  EXPECTED ESCAPE ANALYSIS OUTPUT (go build -gcflags="-m")
//
//  $ go build -gcflags="-m" escape-analysis.go
//
//  # command-line-arguments
//  ./escape-analysis.go:XX:YY: staysOnStack x does not escape
//  ./escape-analysis.go:XX:YY: escapesToHeap x escapes to heap
//  ./escape-analysis.go:XX:YY: printAny v does not escape
//  ./escape-analysis.go:XX:YY: setGlobal x escapes to heap
//  ./escape-analysis.go:XX:YY: make([]int, 3) escapes to heap
//  ./escape-analysis.go:XX:YY: closureEscapes x escapes to heap
//  ./escape-analysis.go:XX:YY: storeInGlobal p does not escape
//  ./escape-analysis.go:XX:YY: caller x escapes to heap
//  ./escape-analysis.go:XX:YY: ... argument does not escape
// ============================================================
//
//  ⚠️  Don't try to outsmart escape analysis!
//  Write clean, readable code. The compiler is very smart.
//  Only check with -gcflags="-m" when profiling shows
//  that GC pressure is a real bottleneck.
// ============================================================