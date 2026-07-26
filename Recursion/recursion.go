// ============================================================
//  RECURSION IN GO — with Python Comparisons
// ============================================================
//  Recursion is when a function calls itself. Go supports it
//  fully, but with some important differences from Python.
//
//  ┌──────────────────────┬──────────────────────────────────┐
//  │         Go           │             Python               │
//  ├──────────────────────┼──────────────────────────────────┤
//  │ func fact(n int) int │ def fact(n):                     │
//  │     if n <= 1 {      │     if n <= 1:                   │
//  │         return 1     │         return 1                 │
//  │     }                │     return n * fact(n - 1)       │
//  │     return n*fact(n-1)│                                  │
//  ├──────────────────────┼──────────────────────────────────┤
//  │ No recursion limit   │ sys.getrecursionlimit() = 1000   │
//  │ (limited by stack)   │ (can be increased with          │
//  │                      │  sys.setrecursionlimit())       │
//  ├──────────────────────┼──────────────────────────────────┤
//  │ No tail-call         │ No tail-call optimization        │
//  │ optimization         │ (CPython)                        │
//  └──────────────────────┴──────────────────────────────────┘
// ============================================================

package main

import (
	"fmt"
	"runtime"
)

// =============================================================================
// 1. BASIC RECURSION — Factorial
// =============================================================================
// Python:
//   def factorial(n):
//       if n <= 1:
//           return 1
//       return n * factorial(n - 1)

func factorial(n int) int {
	// Base case: stop when n <= 1
	if n <= 1 {
		return 1
	}
	// Recursive case: n! = n * (n-1)!
	return n * factorial(n-1)
}

func demoFactorial() {
	fmt.Println("=== Factorial ===")
	for i := 0; i <= 10; i++ {
		fmt.Printf("%2d! = %d\n", i, factorial(i))
	}

	fmt.Println("\nExecution of factorial(5):")
	fmt.Println("  factorial(5)")
	fmt.Println("  = 5 * factorial(4)")
	fmt.Println("  = 5 * 4 * factorial(3)")
	fmt.Println("  = 5 * 4 * 3 * factorial(2)")
	fmt.Println("  = 5 * 4 * 3 * 2 * factorial(1)")
	fmt.Println("  = 5 * 4 * 3 * 2 * 1")
	fmt.Println("  = 120")
}

// =============================================================================
// 2. RECURSION — Fibonacci
// =============================================================================
// Python:
//   def fib(n):
//       if n <= 1:
//           return n
//       return fib(n-1) + fib(n-2)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func demoFibonacci() {
	fmt.Println("\n=== Fibonacci ===")
	for i := 0; i <= 10; i++ {
		fmt.Printf("fib(%2d) = %d\n", i, fibonacci(i))
	}

	fmt.Println("\n⚠️ WARNING: Simple recursion is SLOW for large n")
	fmt.Println("  Each call branches into 2 more calls → O(2ⁿ) time")
	fmt.Println("  fib(45) would take ~10 minutes!")
	fmt.Println("  Use iteration or memoization for performance:")
	fmt.Println()
	fmt.Println("  Iterative (fast):")
	fmt.Println("    func fib(n int) int {")
	fmt.Println("        a, b := 0, 1")
	fmt.Println("        for i := 0; i < n; i++ {")
	fmt.Println("            a, b = b, a + b")
	fmt.Println("        }")
	fmt.Println("        return a")
	fmt.Println("    }")
}

// =============================================================================
// 3. RECURSION — Tree/List Traversal
// =============================================================================
// Recursion is NATURAL for tree-like data structures.
// Python would use the same recursive approach.

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// SumTree recursively sums all values in a tree
func sumTree(node *TreeNode) int {
	if node == nil {
		return 0 // base case: empty subtree
	}
	// Recursive: value + left subtree + right subtree
	return node.Value + sumTree(node.Left) + sumTree(node.Right)
}

// MaxDepth finds the maximum depth of the tree
func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := maxDepth(node.Left)
	right := maxDepth(node.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

func demoTreeTraversal() {
	fmt.Println("\n=== Tree Traversal ===")

	// Build a tree:
	//        1
	//       / \
	//      2   3
	//     / \
	//    4   5
	root := &TreeNode{
		Value: 1,
		Left: &TreeNode{
			Value: 2,
			Left:  &TreeNode{Value: 4},
			Right: &TreeNode{Value: 5},
		},
		Right: &TreeNode{Value: 3},
	}

	fmt.Println("     Tree:")
	fmt.Println("       1")
	fmt.Println("      / \\")
	fmt.Println("     2   3")
	fmt.Println("    / \\")
	fmt.Println("   4   5")
	fmt.Printf("Sum of all values: %d\n", sumTree(root))
	fmt.Printf("Maximum depth: %d\n", maxDepth(root))
}

// =============================================================================
// 4. RECURSION LIMITS — Go vs Python
// =============================================================================
// Python has sys.getrecursionlimit() (default 1000).
// Go has NO recursion limit — but the stack is finite (~1 GB per goroutine).
// You CAN crash Go with deep recursion, but it takes much more.

func demoRecursionLimits() {
	fmt.Println("\n=== Recursion Limits ===")

	// Measure the current goroutine's stack
	var stackSize int
	for i := 0; i < 100; i++ {
		stackSize = guessStackDepth()
	}
	fmt.Printf("Go stack: starts at ~2 KB, grows as needed (default max ~1 GB)\n")
	fmt.Printf("Go has NO sys.setrecursionlimit() equivalent\n")
	fmt.Println()
	fmt.Println("Python: default recursion limit = 1000 calls")
	fmt.Println("Go:     practical recursion limit = ~1,000,000+ calls")
	fmt.Println("         (limited by RAM, not by a hardcoded limit)")
	fmt.Println()
	fmt.Println("⚠️ Neither language does tail-call optimization (TCO)")
	fmt.Println("  A deeply recursive function still builds a call stack")
	fmt.Println("  Go's stack is GROWABLE — Python's is FIXED SIZE")

	// Demonstrate Go's stack measurement
	fmt.Printf("\nCurrent goroutine stack: roughly %d KB\n", stackSize)
}

func guessStackDepth() int {
	// Use runtime to demonstrate stack awareness
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return int(m.StackInuse / 1024)
}

// =============================================================================
// 5. RECURSION vs ITERATION
// =============================================================================
// Go encourages iteration over recursion for simple cases.
// Go has for loops and NO while loops — all iteration uses `for`.

func factorialIterative(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func fibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func demoIterative() {
	fmt.Println("\n=== Recursion vs Iteration ===")
	fmt.Println("Factorial(10):")
	fmt.Printf("  Recursive: %d\n", factorial(10))
	fmt.Printf("  Iterative: %d\n", factorialIterative(10))
	fmt.Println()
	fmt.Println("Fibonacci(10):")
	fmt.Printf("  Recursive:       %d (O(2ⁿ) — SLOW for n > 40)\n", fibonacci(10))
	fmt.Printf("  Iterative:       %d (O(n) — fast)\n", fibonacciIterative(10))
	fmt.Println()
	fmt.Println("When to use recursion:")
	fmt.Println("  ✅ Tree traversal (natural fit)")
	fmt.Println("  ✅ Directory walking")
	fmt.Println("  ✅ Divide-and-conquer algorithms (quicksort, mergesort)")
	fmt.Println("  ❌ Simple loops (use iteration)")
	fmt.Println("  ❌ Large depths (risk stack overflow)")
}

// =============================================================================
// 6. COMPLETE COMPARISON TABLE
// =============================================================================
//
// ┌───────────────────────┬──────────────────────────┬──────────────────────────┐
// │       Feature         │          Go              │         Python           │
// ├───────────────────────┼──────────────────────────┼──────────────────────────┤
// │ Recursion limit       │ No hard limit            │ sys.getrecursionlimit()  │
// │                       │ (stack is growable)      │ (default 1000)           │
// ├───────────────────────┼──────────────────────────┼──────────────────────────┤
// │ Tail-call optimization│ ❌ No                    │ ❌ No (CPython)          │
// ├───────────────────────┼──────────────────────────┼──────────────────────────┤
// │ Multiple recursion    │ func f(n) { f(n-1) }    │ def f(n): f(n-1)         │
// │ (mutual)              │ func g(n) { f(n) }      │ def g(n): f(n)           │
// ├───────────────────────┼──────────────────────────┼──────────────────────────┤
// │ Anonymous recursion   │ var f func(int)          │ f = lambda n: ...        │
// │                       │ f = func(n int) { f(n) } │ (need to declare first)  │
// ├───────────────────────┼──────────────────────────┼──────────────────────────┤
// │ Default recursion     │ for loop (Go has no      │ while or for loop        │
// │ pattern               │ while keyword)           │                          │
// └───────────────────────┴──────────────────────────┴──────────────────────────┘

// =============================================================================
// MAIN
// =============================================================================

func main() {
	demoFactorial()
	demoFibonacci()
	demoTreeTraversal()
	demoRecursionLimits()
	demoIterative()
}

// =============================================================================
// CHEAT SHEET
// =============================================================================
//
// // Recursive factorial
// func fact(n int) int {
//     if n <= 1 { return 1 }
//     return n * fact(n - 1)
// }
//
// // Iterative (preferred for simple cases)
// func fact(n int) int {
//     r := 1
//     for i := 2; i <= n; i++ { r *= i }
//     return r
// }
//
// // Tree recursion (natural fit)
// func sum(n *Node) int {
//     if n == nil { return 0 }
//     return n.Value + sum(n.Left) + sum(n.Right)
// }
