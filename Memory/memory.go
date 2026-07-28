// ============================================================
//  MEMORY IN GO — Stack vs Heap
// ============================================================
//
//  Go manages memory automatically with a garbage collector.
//  Values live on either the STACK or the HEAP.
//
//  ┌──────────────────────────────────────────────────────────┐
//  │  STACK (fast, automatic)                                 │
//  │  ┌────────────────────┐                                  │
//  │  │  func main()       │  ← stack frame for main         │
//  │  │  ├─ x: int = 42   │                                  │
//  │  │  └─ y: string     │                                  │
//  │  │                     │                                  │
//  │  │  func add(a, b)     │  ← stack frame for add          │
//  │  │  ├─ a: int = 10   │                                  │
//  │  │  └─ b: int = 20   │                                  │
//  │  └────────────────────┘                                  │
//  │   Cleaned up when function returns (LIFO)                │
//  └──────────────────────────────────────────────────────────┘
//
//  ┌──────────────────────────────────────────────────────────┐
//  │  HEAP (slower, persistent)                               │
//  │  ┌────────────────────┐                                  │
//  │  │  *Person{...}      │  ← survives function return      │
//  │  │  slice backing arr │  ← grows dynamically             │
//  │  │  S: "big string"   │  ← large values                  │
//  │  └────────────────────┘                                  │
//  │   Cleaned up by GARBAGE COLLECTOR when unreachable       │
//  └──────────────────────────────────────────────────────────┘
//
//  ┌────────────┬────────────────────────────────────────────────┐
//  │    Go      │                 Python                         │
//  ├────────────┼────────────────────────────────────────────────┤
//  │ Stack:     │ Everything is heap-allocated (object model).   │
//  │ int, bool, │ Small ints are cached (-5 to 257).              │
//  │ structs    │                                                │
//  │ (no &),    │                                                │
//  │ arrays     │                                                │
//  │            │ Python has NO stack allocation for user types. │
//  │            │                                                │
//  │ Heap:      │                                                │
//  │ pointers,  │                                                │
//  │ slices,    │                                                │
//  │ maps,      │                                                │
//  │ interfaces │                                                │
//  └────────────┴────────────────────────────────────────────────┘
// ============================================================

package main

import "fmt"

type LegoBrick struct {
	Color string
	Size  int
}

// buildOnStack creates a LegoBrick ON THE STACK
// The brick is local — nobody needs it after this function
func buildOnStack() {
	brick := LegoBrick{Color: "Red", Size: 4}
	fmt.Printf("  🧱 Stack brick: %+v (at %p)\n", brick, &brick)
}

// buildOnHeap creates a LegoBrick ON THE HEAP
// Because we return a pointer, the brick must outlive this function
// Go's compiler detects this — called ESCAPE ANALYSIS
func buildOnHeap() *LegoBrick {
	brick := LegoBrick{Color: "Blue", Size: 6}
	fmt.Printf("  🧱 Heap brick (created locally): %+v\n", brick)
	return &brick // This causes escape to heap!
}

// sliceOnHeap — slice's backing array is on the HEAP
func sliceOnHeap() []int {
	nums := make([]int, 3)
	nums[0] = 10
	nums[1] = 20
	nums[2] = 30
	fmt.Printf("  📦 Slice: header on stack, data on heap: %v\n", nums)
	return nums
}

type Building struct {
	Name string
	Cost int
}

func (b *Building) Describe() {
	fmt.Printf("  🏗️  Building '%s' cost $%d\n", b.Name, b.Cost)
}

func main() {
	fmt.Println("========================================")
	fmt.Println("  🧠  STACK vs HEAP DEMO 🧠")
	fmt.Println("========================================")
	fmt.Println()

	// ── Scenario 1: Stack allocation ──
	fmt.Println("─── Scenario 1: Stack allocation ───")
	fmt.Println()
	buildOnStack()
	fmt.Println("  👤 Back in main — stack space is clean!")
	fmt.Println()

	// ── Scenario 2: Heap allocation ──
	fmt.Println("─── Scenario 2: Heap allocation (escape) ───") 
	fmt.Println()
	heapBrick := buildOnHeap()
	fmt.Printf("  👤 Got brick from heap: %+v (addr: %p)\n", *heapBrick, heapBrick)
	fmt.Println("  👤 The brick survived buildOnHeap() returning!")
	fmt.Println()

	// ── Scenario 3: Slice internals ──
	fmt.Println("─── Scenario 3: Slice internals ───")
	fmt.Println()
	mySlice := sliceOnHeap()
	fmt.Printf("  👤 Slice value: %v (len=%d, cap=%d)\n", mySlice, len(mySlice), cap(mySlice))
	fmt.Println("  👤 Slice header (ptr, len, cap) on stack, data on heap")
	fmt.Println()

	// ── Scenario 4: Pointer receiver ──
	fmt.Println("─── Scenario 4: Pointer receiver method ───")
	fmt.Println()
	house := Building{Name: "Lego House", Cost: 100}
	house.Describe()
	fmt.Println("  👤 Go auto-takes &house because Describe uses *Building")
	fmt.Println()

	// ── Scenario 5: Memory visualization ──
	fmt.Println("─── Scenario 5: Memory Map Visualization ───")
	fmt.Println()
	fmt.Println("  ┌────────────────── STACK ──────────────────┐")
	fmt.Println("  │  buildOnStack() frame (reclaimed)          │")
	fmt.Println("  │    └─ brick {Red, 4} ✅ GONE               │")
	fmt.Println("  │                                            │")
	fmt.Println("  │  main() frame                              │")
	fmt.Println("  │    ├─ heapBrick = *ptr ──────→ (points to  │")
	fmt.Println("  │    ├─ mySlice  = {ptr,len,cap}   heap)     │")
	fmt.Println("  │    └─ house    = {Lego House, 100}         │")
	fmt.Println("  └────────────────────────────────────────────┘")
	fmt.Println()
	fmt.Println("  ┌─────────────────── HEAP ──────────────────┐")
	fmt.Println("  │  {Blue, 6}  ← buildOnHeap returned &      │")
	fmt.Println("  │  [10, 20, 30]  ← sliceOnHeap backing arr  │")
	fmt.Println("  └────────────────────────────────────────────┘")
	fmt.Println()

	fmt.Println("========================================")
	fmt.Println("  🧠  STACK vs HEAP COMPLETE 🧠")
	fmt.Println("========================================")
}

// ============================================================
//  ESCAPE ANALYSIS CHEAT SHEET
// ============================================================
//
//  ┌──────────────────────┬───────────┬──────────────────────┐
//  │ What you write       │ Stack?    │ Heap?                │
//  ├──────────────────────┼───────────┼──────────────────────┤
//  │ x := 42              │ ✅ Stack  │                      │
//  │ s := struct{...}     │ ✅ Stack  │                      │
//  │ arr := [3]int{...}   │ ✅ Stack  │                      │
//  │ fmt.Println(&x)      │           │ ✅ Heap (addr taken) │
//  │ return &x            │           │ ✅ Heap (escaped)    │
//  │ make([]int, n)       │           │ ✅ Heap (slice data) │
//  │ make(map[string]int) │           │ ✅ Heap (map intern) │
//  │ box := &Struct{...}  │           │ ✅ Heap (explicit)   │
//  │ interface{}          │           │ ✅ Heap (interface)  │
//  └──────────────────────┴───────────┴──────────────────────┘
//
//  ⚠️ You're NOT supposed to optimize for this!
//  Write clean code. Go's escape analysis is smart.
//  Check with: go build -gcflags="-m"
//
//  ⚡ Performance: Stack < Heap < GC (garbage collection)
// ============================================================