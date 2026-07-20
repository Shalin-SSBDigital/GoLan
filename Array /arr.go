// ============================================================
//  ARRAYS IN GO — with Python Comparisons
// ============================================================
//  Arrays are fixed-length sequences of the same type.
//  Key property: size is PART OF THE TYPE ([5]int != [3]int)
//
// ┌─────────────────────┬────────────────────────────────────────────┐
// │        Go           │                 Python                     │
// ├─────────────────────┼────────────────────────────────────────────┤
// │ [5]int              │ array.array('i', [0, 0, 0, 0, 0])          │
// │ Fixed size at       │ OR                                             │
// │ compile time        │ import numpy as np; np.array([0]*5)        │
// │                     │                                            │
// │ Size is PART OF     │ Python lists are dynamic — size is NOT     │
// │ the TYPE — you      │ part of the type. A list of 3 ints and     │
// │ CANNOT resize       │ a list of 5 ints are the SAME type.       │
// │                     │                                            │
// │ USED RARELY in Go   │ Python lists are used everywhere instead   │
// │ — slices are the    │ of array.array                           │
// │ idiomatic choice    │                                            │
// └─────────────────────┴────────────────────────────────────────────┘
// ============================================================

package main

import "fmt"

func main() {
	// ============================================================
	//  DECLARATION WITH DEFAULT ZERO VALUES
	//  Syntax: var name [size]type
	// ============================================================
	// Declares an array of 5 integers. All 5 elements are
	// automatically initialized to 0 (the zero-value for int).
	// No manual initialization is required.
	//
	// Memory layout: [0][0][0][0][0]
	// Indices:        0  1  2  3  4
	//
	// Python equivalent:
	//   import array
	//   a = array.array('i', [0, 0, 0, 0, 0])
	//   # But Python's array.array IS resizable — it's more like a slice.
	//   # True fixed-size arrays are only in numpy:
	//   import numpy as np
	//   a = np.zeros(5, dtype=np.int32)  # truly fixed size
	// ============================================================
	var a [5]int
	fmt.Println("emp:", a)
	// Output: emp: [0 0 0 0 0]

	// ============================================================
	//  SET & GET ELEMENTS BY INDEX
	//  Indexing is 0-based: arr[index] = value
	// ============================================================
	// Sets the element at index 4 (the 5th element) to 100.
	// Accessing an index outside [0..4] causes a compile error
	// or runtime panic.
	//
	// Python:
	//   a[4] = 100          # same syntax!
	//   a[5] = 200          # raises IndexError
	//   a[-1] = 100         # Go does NOT support negative indices
	//                       # Python counts from the end
	// ============================================================
	a[4] = 100
	fmt.Println("set:", a)
	// Output: set: [0 0 0 0 100]
	fmt.Println("get:", a[4])
	// Output: get: 100

	// ============================================================
	//  BUILT-IN len() FUNCTION
	// ============================================================
	// len() returns the compile-time known length of the array.
	// For arrays (not slices), len is always fixed.
	//
	// Python:  len(a) → 5  (same function name!)
	// Unlike Go, Python's len() for a list gives the CURRENT
	// number of elements, which can change.
	// ============================================================
	fmt.Println("len:", len(a))
	// Output: len: 5

	// ============================================================
	//  SHORT-HAND DECLARATION WITH LITERAL
	//  Syntax: name := [size]type{val1, val2, ...}
	// ============================================================
	// Declares and initializes b with values 1 through 5.
	// The values match the indices positionally.
	//
	// Memory layout: [1][2][3][4][5]
	//
	// Python:
	//   b = [1, 2, 3, 4, 5]   # creates a list
	//   # Python lists can grow and shrink — they're slices, not arrays
	//   b.append(6)            # OK in Python, IMPOSSIBLE in Go
	//   b.pop()                # OK in Python, IMPOSSIBLE in Go
	// ============================================================
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	// Output: dcl: [1 2 3 4 5]

	// ============================================================
	//  COMPILER-COUNTED SIZE WITH [...]
	//  Syntax: name := [...]type{val1, val2, ...}
	// ============================================================
	// The compiler counts the values and sets the size for you.
	// [...]int{1, 2, 3, 4, 5} is identical to [5]int{1, 2, 3, 4, 5}.
	// This is also the ONLY way to declare a constant-size array
	// literal without repeating the count.
	//
	// Python has no equivalent — list literals always create
	// dynamic-length lists. There is no "infer the size from
	// the literal" syntax because sizes aren't part of types.
	// ============================================================
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	// Output: dcl: [1 2 3 4 5]

	// ============================================================
	//  SPARSE (KEYED) INITIALIZATION
	//  Syntax: [...]type{index: value, ...}
	// ============================================================
	// You can specify explicit indices with the index:value syntax.
	// Indices not listed get zero values.
	//
	// Here: [0]=100, [3]=400, [4]=500
	// Result: [100, 0, 0, 400, 500]
	//
	// Python equivalent (using a list comprehension):
	//   b = [100 if i == 0 else 400 if i == 3 else 500 if i == 4 else 0
	//        for i in range(5)]
	//   # Go's syntax is much cleaner for this!
	// ============================================================
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)
	// Output: idx: [100 0 0 400 500]

	// ============================================================
	//  MULTI-DIMENSIONAL ARRAYS
	//  Syntax: var name [rows][cols]type
	// ============================================================
	// Declares a 2x3 array (2 rows, 3 columns). All elements
	// default to 0. Nested for loops fill each element with
	// the sum of its row and column indices.
	//
	// range 2  -> i = 0, 1  (rows)
	// range 3  -> j = 0, 1, 2 (columns)
	// twoD[1][2] = 1 + 2 = 3
	//
	// Python:
	//   twoD = [[0] * 3 for _ in range(2)]  # list of lists
	//   for i in range(2):
	//       for j in range(3):
	//           twoD[i][j] = i + j
	//   # Python lists of lists are jagged by nature (each row
	//   # can be a different length). Go's [2][3]int is RECTANGULAR —
	//   # every row MUST have exactly 3 columns.
	// ============================================================
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	// Output: 2d:  [[0 1 2] [1 2 3]]

	// ============================================================
	//  2D ARRAY LITERAL INITIALIZATION
	// ============================================================
	// You can also initialize a 2D array with nested literals.
	// Each inner { } represents one row.
	// The outer size [2][3] must match the literal's dimensions.
	//
	// Python:
	//   twoD = [[1, 2, 3], [1, 2, 3]]
	//   # No type-enforced shape — Python won't complain if rows
	//   # have different lengths. Go gives a COMPILE ERROR.
	// ============================================================
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
	// Output: 2d:  [[1 2 3] [1 2 3]]

	// ============================================================
	//  SUMMARY: Array Key Points (with Python comparison)
	// ============================================================
	//  ┌──────────────────┬──────────────────────┬──────────────────────────┐
	//  │ Property         │ Go                    │ Python                    │
	//  ├──────────────────┼──────────────────────┼──────────────────────────┤
	//  │ Fixed size       │ Set at compile time   │ list: dynamic;           │
	//  │                  │ — never changes      │ array.array: resizable   │
	//  │                  │                       │ numpy: fixed             │
	//  │ Value type       │ Assigning copies ALL  │ Assignment copies        │
	//  │                  │ elements (expensive)  │ the REFERENCE (shallow)  │
	//  │ len(arr)         │ Compile-time constant │ Runtime length            │
	//  │ Zero values      │ Auto-initialized     │ Must fill explicitly      │
	//  │ Compare          │ arr == arr2 works    │ [1] == [1] → True (list) │
	//  │                  │ (same type only)      │                           │
	//  │ Negative index   │ NOT supported         │ a[-1] = last element     │
	//  │ Type identity    │ Size is PART of type  │ Size is NEVER part of    │
	//  │                  │                       │ the type                 │
	//  └──────────────────┴──────────────────────┴──────────────────────────┘
	//
	// Bottom line: Go arrays are VALUE types with COMPILE-TIME fixed size.
	// Python lists are REFERENCE types with RUNTIME dynamic size.
	// They are fundamentally different tools. Go slices (not arrays)
	// are the closest equivalent to Python lists.
	// ============================================================
}
