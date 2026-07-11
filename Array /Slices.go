// ============================================================
//  SLICES IN GO
//  Slices are dynamically-sized, flexible views into arrays.
//  They are the most common data structure in Go (used more
//  than arrays in practice).
// ============================================================
//
//  SLICE INTERNALS (runtime representation):
//    ┌──────────┐
//    │  ptr     │──> underlying array (backing array)
//    │  len     │──> number of accessible elements
//    │  cap     │──> max elements before reallocation
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
    // ============================================================
    s = make([]string, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
    // Output: emp: [  ] len: 3 cap: 3

    // ============================================================
    //  SET & GET BY INDEX (same as arrays)
    // ============================================================
    // Index bounds are 0..len(s)-1. Accessing outside panics.
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
    //  SUMMARY: Slices vs Arrays
    // ============================================================
    //  ┌──────────────────┬──────────────────┬──────────────────┐
    //  │ Feature          │ Array [N]T       │ Slice []T        │
    //  ├──────────────────┼──────────────────┼──────────────────┤
    //  │ Size             │ Fixed (compile)   │ Dynamic (runtime)│
    //  │ Type identity    │ Size IS part of   │ Size NOT part of │
    //  │                 │ the type          │ the type         │
    //  │ Assignment       │ Copies ALL elems │ Copies header    │
    //  │                 │ (expensive)      │ (cheap)          │
    //  │ Comparison       │ == works         │ slices.Equal()   │
    //  │ Growth           │ Impossible       │ append()         │
    //  └──────────────────┴──────────────────┴──────────────────┘
    // ============================================================
}