// ============================================================
//  ARRAYS IN GO
//  Arrays are fixed-length sequences of the same type.
//  Key property: size is PART OF THE TYPE ([5]int != [3]int)
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
	// NOTE: b was [5]int — so [...] here resolves to [5]int
	// because the highest key index (3) plus computed positions
	// yields 5 elements.
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
	// ============================================================
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
	// Output: 2d:  [[1 2 3] [1 2 3]]

	// ============================================================
	//  SUMMARY: Array Key Points
	// ============================================================
	//  ┌──────────────────┬──────────────────────────────────────┐
	//  │ Property         │ Behavior                             │
	//  ├──────────────────┼──────────────────────────────────────┤
	//  │ Fixed size       │ Set at compile time — never changes  │
	//  │ Value type       │ Assigning copies ALL elements (cost) │
	//  │ len(arr)         │ Returns the compile-time length      │
	//  │ Zero values      │ All elements auto-initialized to 0   │
	//  │ Compare          │ arr == arr2 allowed (same type only) │
	//  └──────────────────┴──────────────────────────────────────┘
	// ============================================================
}