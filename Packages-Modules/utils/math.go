// ============================================================
//  PACKAGE UTILS — second file in the same package
// ============================================================
//  This is part of the SAME `utils` package as strings.go.
//  They share the same namespace — functions defined here
//  can call functions from strings.go and vice versa.
//
//  In Python, math.go would be a SEPARATE module:
//    utils/math.py  →  from utils.math import is_even
//
//  In Go, both files are:
//    import "golan/packages-modules/utils"
//    utils.IsEven(5)
//    utils.Reverse("hello")
//
//  No sub-namespace — everything is `utils.XXX`.

package utils

import "fmt"

// ---------- Number utilities ----------

// IsEven returns true if n is even
func IsEven(n int) bool {
	return n%2 == 0
}

// Max returns the larger of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Unexported helper — only visible within utils package
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Clamp clamps a value between min and max using the unexported min function
func Clamp(value, low, high int) int {
	// Calls the unexported min function (same package!)
	if value < low {
		return low
	}
	return min(value, high)
}

func init() {
	fmt.Println("[utils/math] package initialized")
}
