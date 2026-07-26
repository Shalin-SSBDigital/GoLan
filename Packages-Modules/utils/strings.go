// ============================================================
//  PACKAGE UTILS — demonstrates MULTI-FILE packages
// ============================================================
//  A package can span MULTIPLE .go files.
//  All files with `package utils` share the same namespace.
//
//  In Python, helpers might be in separate modules:
//    utils/
//    ├── __init__.py
//    ├── strings.py
//    └── numbers.py
//
//    from utils.strings import reverse
//    from utils.numbers import is_even
//
//  In Go, they're all the SAME package:
//    utils/
//    ├── strings.go    → package utils
//    └── numbers.go    → package utils
//
//    import "golan/packages-modules/utils"
//    utils.Reverse("hello")
//    utils.IsEven(5)
//
//  Everything in the package is shared — functions, types, variables.
//  Split files for ORGANIZATION, not namespace isolation.

package utils

import (
	"fmt"
	"strings"
)

// ---------- String utilities ----------

// Reverse returns a reversed version of s
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// CountWords returns the number of words in s
func CountWords(s string) int {
	if strings.TrimSpace(s) == "" {
		return 0
	}
	return len(strings.Fields(s))
}

// Unexported helper — only available within utils package
func logUsage(fn string) {
	fmt.Println("[utils] called:", fn)
}

func init() {
	fmt.Println("[utils/strings] package initialized")
}
