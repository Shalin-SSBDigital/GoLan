// ============================================================
//  strings.Builder — Efficient String Building
// ============================================================
//  strings.Builder is Go's answer to efficient string concatenation.
//  It internally uses a []byte buffer that grows as needed.
//
//  ┌────────────────────────┬────────────────────────────────┐
//  │          Go            │            Python              │
//  ├────────────────────────┼────────────────────────────────┤
//  │ strings.Builder        │ ''.join(list) or io.StringIO  │
//  │ sb.WriteString(s)      │ f.write(s)                    │
//  │ sb.WriteRune(r)        │ f.write(chr(r))               │
//  │ sb.WriteByte(b)        │ f.write(bytes([b]))            │
//  │ sb.String()            │ f.getvalue()                  │
//  │ sb.Grow(n)             │ Pre-allocate capacity          │
//  ├────────────────────────┼────────────────────────────────┤
//  │ s += x (BAD in loop)   │ s += x (BAD in loop)          │
//  │ sb.WriteString (GOOD)  │ ''.join(list) (GOOD)           │
//  └────────────────────────┴────────────────────────────────┘
// ============================================================

package main

import (
	"fmt"
	"strings"
	"time"
)

// =============================================================================
// 1. THE PROBLEM — Why += is Slow in Loops
// =============================================================================
// Every time you do s += x, Go creates a NEW string and copies
// ALL the old data. This is O(n²) — quadratic time.
//
// Python has the exact same problem: s += x in a loop is slow.

func demoProblem() {
	fmt.Println("=== The Problem: += in Loops ===")
	fmt.Println("Strings are IMMUTABLE — every += creates a NEW string")
	fmt.Println("and COPIES all previous data.")
	fmt.Println()
	fmt.Printf("%-30s %s\n", "Operation", "What happens")
	fmt.Printf("%-30s %s\n", "------", "------")
	fmt.Printf("%-30s %s\n", `s = ""`, "zero-length string")
	fmt.Printf("%-30s %s\n", `s += "a"`, "allocate 1 byte, write 'a'")
	fmt.Printf("%-30s %s\n", `s += "b"`, "allocate 2 bytes, copy 'ab'")
	fmt.Printf("%-30s %s\n", `s += "c"`, "allocate 3 bytes, copy 'abc'")
	fmt.Println()
	fmt.Println("Total work: 1 + 2 + 3 + ... + n = O(n²)")
	fmt.Println("With Builder: O(n) — no copies until final String()")
}

// =============================================================================
// 2. BENCHMARK — += vs Builder
// =============================================================================

func demoPerformance() {
	fmt.Println("\n=== Performance: += vs Builder ===")

	n := 100_000

	// Method 1: += (slow)
	start := time.Now()
	s := ""
	for i := 0; i < n; i++ {
		s += "x"
	}
	plusTime := time.Since(start)

	// Method 2: strings.Builder (fast)
	start = time.Now()
	var sb strings.Builder
	sb.Grow(n) // pre-allocate to avoid reallocations
	for i := 0; i < n; i++ {
		sb.WriteString("x")
	}
	result := sb.String()
	builderTime := time.Since(start)

	fmt.Printf("Concatenating %d strings:\n", n)
	fmt.Printf("  += operator:    %v (O(n²) — SLOW)\n", plusTime)
	fmt.Printf("  strings.Builder: %v (O(n) — FAST)\n", builderTime)
	fmt.Printf("  Builder is %.0fx faster!\n", float64(plusTime)/float64(builderTime))
	fmt.Printf("  Both produce length %d strings\n", len(s))
}

// =============================================================================
// 3. BUILDER METHODS
// =============================================================================

func demoBuilderMethods() {
	fmt.Println("\n=== Builder Methods ===")

	var sb strings.Builder

	// WriteString — appends a string
	sb.WriteString("Hello")
	fmt.Println("After WriteString(\"Hello\"):", sb.String())

	// WriteRune — appends a single Unicode character
	sb.WriteRune(' ')
	sb.WriteRune('世') // multi-byte rune
	sb.WriteRune('界')
	fmt.Println("After WriteRune(' '), WriteRune('世'), WriteRune('界'):", sb.String())

	// WriteByte — appends a single byte
	sb.WriteByte('!')
	fmt.Println("After WriteByte('!'):", sb.String())

	// Write — appends a byte slice
	sb.Write([]byte(" How are you?"))
	fmt.Println("After Write([]byte):", sb.String())

	// Len — current length (in bytes, not characters)
	fmt.Printf("Length (bytes): %d\n", sb.Len())

	// Cap — current capacity
	fmt.Printf("Capacity: %d bytes\n", sb.Cap())

	// Reset — clear the builder (reuse without reallocating)
	var sb2 strings.Builder
	sb2.WriteString("temporary")
	fmt.Println("\nBefore Reset:", sb2.String())
	sb2.Reset()
	fmt.Println("After Reset:", sb2.String(), "(empty)")
	fmt.Println("  Reset keeps the underlying buffer — no new allocation")
}

// =============================================================================
// 4. Grow — PRE-ALLOCATION
// =============================================================================
// If you know the final size, call Grow(n) to pre-allocate.
// This avoids MULTIPLE reallocations as the buffer grows.

func demoGrow() {
	fmt.Println("\n=== Pre-allocation with Grow ===")

	n := 100_000

	// Without Grow (multiple reallocations)
	start := time.Now()
	var sb1 strings.Builder
	for i := 0; i < n; i++ {
		sb1.WriteString("x")
	}
	withoutGrow := time.Since(start)

	// With Grow (single allocation)
	start = time.Now()
	var sb2 strings.Builder
	sb2.Grow(n) // reserve space for n bytes
	for i := 0; i < n; i++ {
		sb2.WriteString("x")
	}
	withGrow := time.Since(start)

	fmt.Printf("Without Grow: %v\n", withoutGrow)
	fmt.Printf("With Grow:    %v\n", withGrow)
	fmt.Printf("Grow is %.0fx faster\n", float64(withoutGrow)/float64(withGrow))
	fmt.Println()
	fmt.Println("Best practice: if you know the final size, call Grow(n)")
	fmt.Println("  Otherwise, Builder doubles capacity automatically")
}

// =============================================================================
// 5. REAL-WORLD EXAMPLES
// =============================================================================

// BuildCSV builds a CSV row using Builder
func buildCSV(fields []string) string {
	var sb strings.Builder
	for i, f := range fields {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(f)
	}
	sb.WriteByte('\n')
	return sb.String()
}

// BuildQuery builds a simple SQL query
func buildQuery(table string, columns []string, conditions map[string]string) string {
	var sb strings.Builder

	sb.WriteString("SELECT ")
	for i, col := range columns {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(col)
	}

	sb.WriteString(" FROM ")
	sb.WriteString(table)

	if len(conditions) > 0 {
		sb.WriteString(" WHERE ")
		first := true
		for k, v := range conditions {
			if !first {
				sb.WriteString(" AND ")
			}
			sb.WriteString(k)
			sb.WriteString(" = ")
			sb.WriteByte('\'')
			sb.WriteString(v)
			sb.WriteByte('\'')
			first = false
		}
	}

	sb.WriteByte(';')
	return sb.String()
}

// FormatList formats a list with commas and "and"
func formatList(items []string) string {
	switch len(items) {
	case 0:
		return ""
	case 1:
		return items[0]
	case 2:
		return items[0] + " and " + items[1]
	default:
		var sb strings.Builder
		// Pre-allocate approximate size
		totalLen := len(items[0])*len(items) + (len(items)-1)*2 + 4
		sb.Grow(totalLen)

		for i, item := range items {
			if i == len(items)-1 {
				sb.WriteString(" and ")
			} else if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(item)
		}
		return sb.String()
	}
}

func demoRealWorld() {
	fmt.Println("\n=== Real-World Examples ===")

	// CSV builder
	csv := buildCSV([]string{"Alice", "30", "Engineer"})
	fmt.Printf("CSV row: %q\n", csv)

	// Query builder
	query := buildQuery("users",
		[]string{"id", "name", "email"},
		map[string]string{"active": "true", "role": "admin"},
	)
	fmt.Printf("SQL query: %s\n", query)

	// List formatter
	fmt.Printf("List [a]:          %q\n", formatList([]string{"a"}))
	fmt.Printf("List [a, b]:       %q\n", formatList([]string{"a", "b"}))
	fmt.Printf("List [a, b, c]:    %q\n", formatList([]string{"a", "b", "c"}))
	fmt.Printf("List [a,b,c,d]:    %q\n", formatList([]string{"a", "b", "c", "d"}))
}

// =============================================================================
// 6. strings.Join (Alternative)
// =============================================================================
// If you already have a []string slice, strings.Join is even simpler.
// It internally uses strings.Builder too.

func demoStringsJoin() {
	fmt.Println("\n=== strings.Join (Alternative) ===")

	words := []string{"Go", "is", "awesome"}
	sentence := strings.Join(words, " ")
	fmt.Printf("strings.Join(%v, \" \") = %q\n", words, sentence)
	fmt.Println()
	fmt.Println("When to use which:")
	fmt.Println("  strings.Join(slice, sep)  → when you ALREADY have a []string")
	fmt.Println("  strings.Builder           → when building INCREMENTALLY")
	fmt.Println("  sb.WriteString(s)         → adding strings one at a time")
	fmt.Println("  + operator                → only for 2-3 strings")
}

// =============================================================================
// 7. COMPLETE COMPARISON TABLE
// =============================================================================
//
// ┌───────────────────────────────┬───────────────────────┬──────────────────────────┐
// │          Concept              │         Go            │         Python            │
// ├───────────────────────────────┼───────────────────────┼──────────────────────────┤
// │ Efficient concat (has slice)  │ strings.Join(s, sep)  │ sep.join(list)            │
// │ Efficient concat (incremental)│ strings.Builder       │ io.StringIO               │
// │ Append string                 │ sb.WriteString(s)     │ f.write(s)                │
// │ Append rune / char            │ sb.WriteRune(r)       │ f.write(chr(r))           │
// │ Append byte                   │ sb.WriteByte(b)       │ f.write(bytes([b]))        │
// │ Final result                  │ sb.String()           │ f.getvalue()              │
// │ Pre-allocate                  │ sb.Grow(n)            │ No direct equivalent      │
// │ Reset for reuse               │ sb.Reset()            │ f.seek(0); f.truncate()   │
// │ Current length                │ sb.Len()              │ f.tell()                  │
// │ Loop concat (BAD)             │ s += x (O(n²))        │ s += x (O(n²))            │
// └───────────────────────────────┴───────────────────────┴──────────────────────────┘

// =============================================================================
// MAIN
// =============================================================================

func main() {
	demoProblem()
	demoPerformance()
	demoBuilderMethods()
	demoGrow()
	demoRealWorld()
	demoStringsJoin()
}

// =============================================================================
// CHEAT SHEET
// =============================================================================
//
// // 1. Create builder
// var sb strings.Builder
//
// // 2. Pre-allocate (optional, but recommended if you know the size)
// sb.Grow(1000)
//
// // 3. Write to it
// sb.WriteString("hello")
// sb.WriteRune(' ')
// sb.WriteByte('!')
// sb.Write([]byte("world"))
//
// // 4. Get result
// result := sb.String()
//
// // 5. Reuse (keeps buffer, no new allocation)
// sb.Reset()
//
// // 6. Or if you already have a slice
// result := strings.Join(slice, ", ")
