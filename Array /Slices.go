// ============================================================
//  SLICES IN GO — with Python Comparisons
// ============================================================
//  Slices are dynamically-sized, flexible views into arrays.
//  They are the most common data structure in Go (used more
//  than arrays in practice).
//
// ┌─────────────────────┬────────────────────────────────────────────┐
// │        Go           │                 Python                     │
// ├─────────────────────┼────────────────────────────────────────────┤
// │ []int               │ list — dynamic, resizable                   │
// │                     │                                            │
// │ Slice is a VIEW     │ list[:] creates a SHALLOW COPY (but a      │
// │ into an underlying  │ new list object). Slice syntax SAME but    │
// │ array — NO COPY     │ Go never copies data on slice.             │
// │ on slice            │                                            │
// │                     │                                            │
// │ len and cap         │ Only len — Python has no "capacity"        │
// │ (capacity)          │ concept. Python lists auto-grow.           │
// │                     │                                            │
// │ append() returns    │ list.append() modifies IN-PLACE,           │
// │ a NEW slice header  │ returns None.                              │
// │ (reassign or lose)  │                                            │
// └─────────────────────┴────────────────────────────────────────────┘
// ============================================================
//
//  SLICE INTERNALS (runtime representation):
//    ┌──────────┐
//    │  ptr     │──> underlying array (backing array)
//    │  len     │──> number of accessible elements
//    │  cap     │──> max elements before reallocation
//    └──────────┘
//
//  Python list internals:
//    ┌──────────┐
//    │  ptr     │──> contiguous array of PyObject* pointers
//    │  len     │──> number of elements
//    │  cap     │──> allocated capacity (hidden, auto-managed)
//    └──────────┘
//
//  Key difference from arrays:
//    - Arrays:  [5]int  — fixed size, VALUE type
//    - Slices:  []int   — dynamic size, REFERENCE type
// ============================================================

package main

import (
	"fmt"
	"slices"
)

func slicesDemo() {

	// ============================================================
	//  DECLARATION — nil slice (zero value)
	//  Syntax: var s []type
	// ============================================================
	// A nil slice has no backing array. It has len=0, cap=0, and
	// s == nil is true. You CAN append to a nil slice — Go will
	// allocate the backing array on the first append.
	//
	// Python:
	//   s = None        # similar, but...
	//   s.append("a")   # AttributeError! Can't append to None
	//   # In Go, append(nilSlice, "a") works perfectly.
	//   # Go's nil slice is USABLE — Python's None is not.
	// ============================================================
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)
	// Output: uninit: [] true true

	// ============================================================
	//  make() — create a slice with a backing array
	//  Syntax: s = make([]type, length, capacity)
	//          capacity is optional (defaults to length)
	// ============================================================
	// Creates a slice backed by a new ["" "" ""] array.
	// len = 3 (accessible), cap = 3 (realloc threshold).
	// All elements are zero-valued ("" for strings).
	//
	// Memory:
	//   ptr ──> ["", "", ""]
	//   len = 3
	//   cap = 3
	//
	// Python:
	//   s = [""] * 3      # ["", "", ""]
	//   # Python has no make() or pre-allocated capacity.
	//   # Under the hood, Python lists DO have capacity (over-allocation),
	//   # but you cannot control or observe it directly.
	// ============================================================
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
	// Output: emp: [  ] len: 3 cap: 3

	// ============================================================
	//  SET & GET BY INDEX (same as arrays)
	// ============================================================
	// Index bounds are 0..len(s)-1. Accessing outside panics.
	//
	// Python:
	//   s[0] = "a"          # same syntax
	//   s[-1]               # "c" (negative index) — NOT supported in Go!
	//   s[10]               # IndexError
	// ============================================================
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	// Output: set: [a b c]
	fmt.Println("get:", s[2])
	// Output: get: c

	// ============================================================
	//  len() — number of accessible elements
	// ============================================================
	fmt.Println("len:", len(s))
	// Output: len: 3

	// ============================================================
	//  append() — add elements (may reallocate)
	//  Syntax: s = append(s, elem1, elem2, ...)
	// ============================================================
	// append() is the ONLY way to grow a slice. It always returns
	// a new slice header (ptr/len/cap). If the backing array has
	// enough capacity, it reuses it. If not, it allocates a NEW
	// array (doubling capacity), copies old elements, and returns
	// a slice pointing to the new array.
	//
	// IMPORTANT: ALWAYS assign the result back: s = append(s, v)
	//
	// Python:
	//   s.append("d")          # modifies in-place, returns None
	//   s.extend(["e", "f"])   # adds multiple elements
	//   # Python lists ALWAYS auto-grow. There is no "capacity"
	//   # concept at the language level — the allocator handles it.
	//   # Under the hood, CPython uses the same exponential growth
	//   # strategy (~1.125x), but you never see it.
	//
	// CRITICAL DIFFERENCE:
	//   Go:  s = append(s, x)  # MUST reassign — append MAY return a
	//                            # different backing array
	//   Python: s.append(x)    # mutates in-place, no reassignment needed
	// ============================================================
	s = append(s, "d")        // appends 1 element
	s = append(s, "e", "f")   // appends 2 elements
	fmt.Println("apd:", s)
	// Output: apd: [a b c d e f]

	// ============================================================
	//  copy() — duplicate elements into destination
	//  Syntax: copy(dst, src)  -> returns number of elements copied
	// ============================================================
	// copy() copies min(len(dst), len(src)) elements. Here c is
	// created with len(s), so all 6 elements are copied.
	// After copy, c and s have SEPARATE backing arrays.
	//
	// Python:
	//   c = s.copy()            # shallow copy
	//   c = s[:]                # also a shallow copy
	//   # Python always COPIES — there's no view/generation
	//   # Both create new list objects with copied references.
	//
	// Go's copy is unique because:
	//   1. You control the destination size (partial copy)
	//   2. It copies DATA, not references (no shared mutation risk)
	// ============================================================
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	// Output: cpy: [a b c d e f]

	// ============================================================
	//  SLICING — extract a sub-slice
	//  Syntax: slice[low:high]
	// ============================================================
	// Creates a NEW slice header pointing into the SAME backing
	// array. No data is copied — it's a view, not a clone.
	//
	// Rules:
	//   s[low:high] -> elements from low up to (but not including) high
	//   s[2:5]      -> indices 2, 3, 4
	//   s[:5]       -> from start to index 4  (low defaults to 0)
	//   s[2:]       -> from index 2 to end    (high defaults to len)
	//
	// Python:
	//   l = s[2:5]              # same syntax!
	//   # BUT Python creates a NEW list (shallow copy).
	//   # Go creates a VIEW — no allocation, no copy.
	//
	//   # This matters for performance:
	//   #   Go:  O(1) — just a pointer + length
	//   #   Py:  O(n) — allocates and copies n elements
	//
	//   # ⚠️ Shared backing array GOTCHA:
	//   #   Go: modifying l[0] ALSO modifies s[2] (same memory!)
	//   #   Py: modifying l[0] does NOT affect s[2] (separate list)
	// ============================================================

	// Slice from index 2 to 4 (exclusive of 5): [c d e]
	l := s[2:5]
	fmt.Println("sl1:", l)
	// Output: sl1: [c d e]

	// Slice from start to index 4: [a b c d e]
	l = s[:5]
	fmt.Println("sl2:", l)
	// Output: sl2: [a b c d e]

	// Slice from index 2 to end: [c d e f]
	l = s[2:]
	fmt.Println("sl3:", l)
	// Output: sl3: [c d e f]

	// ============================================================
	//  SLICE LITERAL — declare + initialize
	//  Syntax: t := []type{val1, val2, ...}
	// ============================================================
	// This is the most common way to create a small slice.
	// Go creates a backing array of size 3 and returns a slice
	// pointing to it with len=3, cap=3.
	//
	// Python:
	//   t = ["g", "h", "i"]    # same syntax!
	//   # Same result, very different internals:
	//   # Go: backing array [g h i] with slice header pointing to it
	//   # Py: PyListObject with array of 3 PyObject* pointers
	// ============================================================
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	// Output: dcl: [g h i]

	// ============================================================
	//  slices.Equal() — compare slices (Go 1.21+)
	//  Package: "slices"
	// ============================================================
	// Unlike arrays, slices CANNOT be compared with == (it causes
	// a compile error). Use slices.Equal() instead, which compares
	// element-by-element.
	//
	// Python:
	//   t == t2          # returns True (element-by-element comparison)
	//   # Python lists DO support ==, and it compares by value.
	//   # This is a notable Python convenience — Go does not allow
	//   # == on slices to avoid ambiguity (reference vs value equality).
	// ============================================================
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}
	// Output: t == t2

	// ============================================================
	//  MULTI-DIMENSIONAL SLICES (jagged arrays)
	//  Each inner slice can have a DIFFERENT length
	// ============================================================
	// make([][]int, 3) creates a slice of 3 slices (all nil).
	// Each inner slice is then created separately with make(),
	// allowing each row to have a different length.
	//
	// Result structure:
	//   twoD[0] -> [0]       (len=1)
	//   twoD[1] -> [0 1]     (len=2)
	//   twoD[2] -> [0 1 2]   (len=3)
	//
	// This is DIFFERENT from a 2D array [3][3]int which forces
	// all rows to be the same length.
	//
	// Python:
	//   twoD = [[i + j for j in range(i + 1)] for i in range(3)]
	//   # Python nested lists are ALWAYS jagged by default.
	//   # There is no "rectangular" list type.
	//   # Go: you CHOOSE between rectangular ([N][M]T) and jagged ([][]T).
	//   # Python: all lists are jagged (no rectangular option).
	// ============================================================
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	// Output: 2d:  [[0] [0 1] [0 1 2]]

	// ============================================================
	//  SUMMARY: Slices vs Arrays (with Python comparison)
	// ============================================================
	//  ┌──────────────────┬──────────────────┬──────────────────┬──────────────────────┐
	//  │ Feature          │ Go Array [N]T     │ Go Slice []T     │ Python list           │
	//  ├──────────────────┼──────────────────┼──────────────────┼──────────────────────┤
	//  │ Size             │ Fixed (compile)   │ Dynamic (runtime)│ Dynamic (runtime)     │
	//  │ Type identity    │ Size IS part of   │ Size NOT part of │ Size NEVER part of    │
	//  │                  │ the type          │ the type         │ the type              │
	//  │ Assignment       │ Copies ALL elems  │ Copies header    │ Copies references     │
	//  │                  │ (expensive)      │ (cheap)          │ (shallow copy)        │
	//  │ Comparison       │ == works          │ slices.Equal()   │ == works (value cmp)  │
	//  │ Growth           │ Impossible        │ append()         │ .append() / .extend() │
	//  │ Slice semantics  │ N/A               │ VIEW (no copy)   │ COPY (new list)       │
	//  │ Capacity         │ = length          │ Explicit cap     │ Hidden (auto-managed) │
	//  │ Negative index   │ No                │ No               │ Yes (s[-1])           │
	//  │ Insert/delete    │ N/A               │ append(a[:i]...) │ .insert() / .pop()    │
	//  └──────────────────┴──────────────────┴──────────────────┴──────────────────────┘
	//
	// Bottom line: Go slices are the closest equivalent to Python lists,
	// but with significant differences:
	//   1. Slicing creates a VIEW (not a copy) — watch for shared mutation
	//   2. No negative indices — handle manually: s[len(s)-1] for last
	//   3. No built-in insert/pop at arbitrary index — use append tricks
	//   4. append() MUST be reassigned — it may reallocate
	//   5. No == comparison — use slices.Equal() or loop
	// ============================================================
}
