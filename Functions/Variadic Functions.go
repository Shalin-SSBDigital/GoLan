// ============================================================
//  VARIADIC FUNCTIONS IN GO — with Python Comparisons
// ============================================================
//  A variadic function accepts a variable number of arguments.
//
// ┌──────────────────────┬──────────────────────────────────────────────────┐
// │        Go            │                  Python                          │
// ├──────────────────────┼──────────────────────────────────────────────────┤
// │ func sum(nums ...int)│ def sum(*nums):                                 │
// │     int {            │     total = 0                                   │
// │     total := 0       │     for num in nums:                            │
// │     for _, n := range│         total += num                            │
// │         nums {       │     return total                                │
// │         total += n   │                                                  │
// │     }                │   Go: nums is a SLICE ([]int)                    │
// │     return total     │   Python: nums is a TUPLE                        │
// │ }                    │                                                  │
// │                      │                                                  │
// │ ...int means: take 0 │ *nums means: gather all positional args          │
// │ or more ints, packed │ into a tuple                                     │
// │ into a []int slice   │                                                  │
// │                      │                                                  │
// │ Pass slice with s... │ Pass list with *list                             │
// │ sum(slice...)        │ sum(*lst)                                        │
// └──────────────────────┴──────────────────────────────────────────────────┘
// ============================================================

package main

import "fmt"

// ============================================================
//  VARIADIC FUNCTION
//  Syntax: func name(paramName ...type) returnType { ... }
// ============================================================
// The ... before the type (int) makes this variadic.
// Inside the function, 'nums' is a []int slice.
//
// Python equivalent:
//   def sum(*nums: int) -> int:
//       print(nums)         # nums is a TUPLE
//       total = 0
//       for num in nums:
//           total += num
//       return total
//
// KEY DIFFERENCES:
//   - Go:  nums is a SLICE — you can use range, append, index, etc.
//   - Python: nums is a TUPLE — it's IMMUTABLE (no append, no modify)
//   - Go:  variable must be the LAST parameter (or only parameter)
//   - Python: *args can be ANYWHERE in the parameter list
//   - Go:  only ONE variadic parameter per function
//   - Python: you can have *args AND **kwargs (keyword variadic)
//          Go does NOT support keyword arguments at all!
// ============================================================
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	// ============================================================
	//  CALLING WITH MULTIPLE ARGUMENTS
	// ============================================================
	// You can pass any number of arguments (including zero).
	//
	// Python:
	//   sum(1, 2)      # nums = (1, 2)
	//   sum(1, 2, 3)   # nums = (1, 2, 3)
	// ============================================================
	sum(1, 2)
	// Output: [1 2] 3
	sum(1, 2, 3)
	// Output: [1 2 3] 6

	// ============================================================
	//  PASSING A SLICE TO A VARIADIC FUNCTION
	//  Syntax: funcName(slice...)
	// ============================================================
	// If you already have a slice, you can "expand" it with ...
	// This is like Python's * operator to unpack a list.
	//
	// Python:
	//   nums = [1, 2, 3, 4]
	//   sum(*nums)      # same: unpacks list into positional args
	//
	// Go:
	//   nums := []int{1, 2, 3, 4}
	//   sum(nums...)    # expands slice into individual args
	// ============================================================
	nums := []int{1, 2, 3, 4}
	sum(nums...)
	// Output: [1 2 3 4] 10

	// ============================================================
	//  MIXING REGULAR AND VARIADIC PARAMETERS
	// ============================================================
	// Go: the variadic parameter must be the LAST one.
	// func greet(prefix string, names ...string) { ... }
	//
	// Python: same rule — *args must come before **kwargs
	// def greet(prefix, *names): ...
	//
	// Go has NO equivalent of Python's **kwargs (keyword args).
	// If you need name-value pairs, use map[string]interface{}.
	// ============================================================

	// ============================================================
	//  SUMMARY: Variadic Functions
	// ============================================================
	//  ┌──────────────────────┬──────────────────────┬──────────────────────────┐
	//  │ Feature              │ Go                   │ Python                    │
	//  ├──────────────────────┼──────────────────────┼──────────────────────────┤
	//  │ Syntax               │ func f(nums ...int)  │ def f(*nums):             │
	//  │ Internal type        │ []int (SLICE)        │ tuple (IMMUTABLE)         │
	//  │ Unpack call          │ f(slice...)          │ f(*list)                  │
	//  │ Keyword args         │ NOT supported        │ **kwargs supported        │
	//  │ Must be last param   │ YES                  │ YES (for *args)           │
	//  │ Multiple variadic    │ NO (compiler error)  │ NO (syntax error)         │
	//  │ Default values       │ NOT supported        │ Supported (with care)     │
	//  │ Mixed named + var    │ Supported            │ Supported                 │
	//  │ Zero arguments       │ OK (empty slice)     │ OK (empty tuple)          │
	//  └──────────────────────┴──────────────────────┴──────────────────────────┘
	//
	// Bottom line: Variadic functions work similarly in both languages.
	// The key difference is Go uses SLICES (mutable, flexible) while
	// Python uses TUPLES (immutable, fixed). Go has no **kwargs equivalent
	// because it has no keyword arguments at all.
	// ============================================================
}
