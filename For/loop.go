// ============================================================
//  FOR LOOPS IN GO — with Python Comparisons
// ============================================================
//  Go has ONE looping keyword: `for`. Everything is built from it.
//  There is NO while, no do-while, no until.
//
// ┌──────────────────────┬────────────────────────────────────────────────┐
// │        Go            │                 Python                         │
// ├──────────────────────┼────────────────────────────────────────────────┤
// │ for init; cond; post │ for i in range(n):  — counted iteration        │
// │ for condition        │ while condition:     — condition-only loop     │
// │ for { }              │ while True:          — infinite loop           │
// │ for _, v := range    │ for v in iterable:   — collection iteration    │
// │                      │ for i, v in enumerate(s):  — with index       │
// │                      │                                                │
// │ ONE keyword          │ for + while (two keywords)                     │
// └──────────────────────┴────────────────────────────────────────────────┘
// ============================================================

package main

import ("fmt")

func main() {
	// ============================================================
	//  TYPE 1: STANDARD THREE-PART FOR LOOP
	//  Syntax:  for init; condition; post { ... }
	// ============================================================
	// This is the classic C-style for loop. It has three parts:
	//   - init (i := 0)     -> runs once before the loop starts
	//   - condition (i < 5) -> checked before each iteration;
	//                          loop continues while this is true
	//   - post (i++)        -> runs after each iteration body
	//
	// Execution flow:
	//   1. init (i = 0)
	//   2. condition -> true?  ->  enter body
	//   3. print, then post (i++)
	//   4. condition -> true?  ->  enter body
	//   ... repeat until condition is false
	//
	// Python equivalent:
	//   for i in range(5):          # range(5) = [0, 1, 2, 3, 4]
	//       print("The value of i is:", i)
	//
	// Key differences:
	//   - Python's range() generates values (lazily)
	//   - Go's for-loop structure mirrors C/Java
	//   - In Python, you CANNOT modify i inside the loop to
	//     affect iteration count (range yields the next value).
	//     In Go, you CAN modify i (but shouldn't — it's confusing).
	// ============================================================
	for i := 0; i < 5; i++ {
		fmt.Println("The value of i is: ", i)
	}
	// Output:
	//   The value of i is:  0
	//   The value of i is:  1
	//   The value of i is:  2
	//   The value of i is:  3
	//   The value of i is:  4

	// ============================================================
	//  TYPE 2: FOR LOOP WITH break
	//  break  -> immediately exits the innermost loop
	// ============================================================
	// The 'break' keyword terminates the loop entirely and
	// execution resumes at the first statement after the loop.
	// Here, when j == 3, the loop stops — so only 0, 1, 2 print.
	//
	// Python:
	//   for j in range(5):
	//       if j == 3:
	//           break
	//       print("The value of j is:", j)
	//   # break works the same way in both languages!
	// ============================================================
	for j := 0; j < 5; j++ {
		if j == 3 {
			break   // <-- exits the for loop entirely
		}
		fmt.Println("The value of j is: ", j)
	}
	// Output:
	//   The value of j is:  0
	//   The value of j is:  1
	//   The value of j is:  2

	// ============================================================
	//  TYPE 3: FOR LOOP WITH continue
	//  continue -> skips rest of current iteration, goes to post
	// ============================================================
	// The 'continue' keyword skips the remaining body of the
	// current iteration and jumps to the post statement (i++),
	// then rechecks the condition. Here k == 2 is skipped entirely.
	//
	// Python:
	//   for k in range(5):
	//       if k == 2:
	//           continue
	//       print("The value of k is:", k)
	//   # continue works the same way in both languages!
	// ============================================================
	for k := 0; k < 5; k++ {
		if k == 2 {
			continue   // <-- skips the print below for k=2
		}
		fmt.Println("The value of k is: ", k)
	}
	// Output:
	//   The value of k is:  0
	//   The value of k is:  1
	//   The value of k is:  3
	//   The value of k is:  4

	// ============================================================
	//  TYPE 4: CONDITION-ONLY LOOP (while-style)
	//  Syntax:  for condition { ... }
	// ============================================================
	// Go has no 'while' keyword. Instead, you drop the init and
	// post parts, keeping only the condition. This behaves exactly
	// like a while loop in other languages.
	//
	// Flow:
	//   1. check condition (m < 5) -> if false, exit
	//   2. execute body
	//   3. go back to step 1
	// The variable m is declared and incremented manually inside
	// the body, giving full control over the loop variable.
	//
	// Python:
	//   m = 0
	//   while m < 5:
	//       print("The value of m is:", m)
	//       m += 1
	//   # Python uses a DIFFERENT keyword (while) for this pattern.
	//   # Go uses the SAME keyword (for) with fewer clauses.
	//   # This is Go's philosophy: one way to loop, many patterns.
	// ============================================================
	m := 0
	for m < 5 {
		fmt.Println("The value of m is: ", m)
		m++   // manual increment — without this, the loop runs forever
	}
	// Output:
	//   The value of m is:  0
	//   The value of m is:  1
	//   The value of m is:  2
	//   The value of m is:  3
	//   The value of m is:  4

	// ============================================================
	//  TYPE 5: INFINITE LOOP WITH break
	//  Syntax:  for { ... }
	// ============================================================
	// An empty 'for' clause (no init, no condition, no post) runs
	// forever. It must be terminated with 'break' (or 'return' or
	// 'panic'). This pattern is useful for:
	//   - server event loops (accepting connections)
	//   - worker queues (polling for jobs)
	//   - interactive menus (waiting for "quit" input)
	//
	// Here we manually break when n reaches 5 to avoid hanging.
	//
	// Python:
	//   n = 0
	//   while True:
	//       if n == 5:
	//           break
	//       print("The value of n is:", n)
	//       n += 1
	//   # Python uses `while True:` — Go uses `for { }`
	//   # Both express the same idea: "loop until something happens"
	// ============================================================
	n := 0
	for {
		if n == 5 {
			break   // <-- must have some exit condition!
		}
		fmt.Println("The value of n is: ", n)
		n++
	}
	// Output:
	//   The value of n is:  0
	//   The value of n is:  1
	//   The value of n is:  2
	//   The value of n is:  3
	//   The value of n is:  4

	// ============================================================
	//  TYPE 6: RANGE LOOP
	//  Syntax:  for index, value := range collection { ... }
	// ============================================================
	// 'range' iterates over slices, arrays, maps, strings, or
	// channels. It returns two values per iteration:
	//   - index (position in the collection, 0-based)
	//   - value (copy of the element at that index)
	//
	// Key behaviors:
	//   - The value is a COPY — modifying it does NOT change the
	//     original collection
	//   - Use _ (blank identifier) to discard either return value
	//   - For maps: range returns (key, value) pairs in random order
	//   - For strings: range iterates over runes, not bytes
	//
	// Python:
	//   for idx, val in enumerate(nums):       # with index
	//       print(f"Index: {idx}, Value: {val}")
	//   for val in nums:                        # value only
	//       print(val)
	//   # Python's enumerate() gives (index, value) like Go's range.
	//   # But Go's range works on MULTIPLE types natively.
	//   # Python needs different functions: enumerate, dict.items, etc.
	//   # Go uses `range` for ALL iterable types — uniform syntax.
	// ============================================================
	nums := []string{"a", "b", "c"}
	for idx, val := range nums {
		fmt.Printf("Index: %d, Value: %s\n", idx, val)
	}
	// Output:
	//   Index: 0, Value: a
	//   Index: 1, Value: b
	//   Index: 2, Value: c

	// ============================================================
	//  LABELED BREAK — exit from nested loops
	//  Go-specific feature — Python has no equivalent
	// ============================================================
	// Go allows you to LABEL a for loop and break/continue to
	// that specific label. This is useful for breaking out of
	// nested loops.
	//
	// Python:
	//   # Python has NO labeled break. You'd use:
	//   #   - A flag variable
	//   #   - A function + return
	//   #   - raise StopIteration (ugly)
	//   #   - for...else with break (limited)
	//   # Go's labeled break is cleaner for deep nesting.
	// ============================================================
	fmt.Println("\nLabeled break example:")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i*j > 2 {
				fmt.Printf("Breaking at i=%d, j=%d\n", i, j)
				break outer // <-- exits BOTH loops
			}
			fmt.Printf("  i=%d, j=%d\n", i, j)
		}
	}
	// Output:
	//   i=0, j=0
	//   i=0, j=1
	//   i=0, j=2
	//   i=1, j=0
	//   i=1, j=1
	//   i=1, j=2
	//   Breaking at i=2, j=0

	// ============================================================
	//  FOR-RANGE OVER MAP (keys only)
	// ============================================================
	// Python:  for k in my_dict:
	// ============================================================
	capitals := map[string]string{"Japan": "Tokyo", "France": "Paris", "India": "New Delhi"}
	fmt.Println("\nRange over map (keys only):")
	for country := range capitals { // value omitted — just keys
		fmt.Println(" ", country)
	}
	// Output order will VARY (maps are random order in Go!)
	// Python 3.7+ preserves insertion order.

	// ============================================================
	//  SUMMARY: For Loop Variants in Go (with Python comparison)
	// ============================================================
	//  ┌────────────────────┬────────────────────┬──────────────────────────────┐
	//  │ Go Form            │ Python Equivalent  │ Use case                     │
	//  ├────────────────────┼────────────────────┼──────────────────────────────┤
	//  │ for init; c; post  │ for i in range(n)  │ Counted iteration            │
	//  │ for condition      │ while cond:        │ Condition-only loop          │
	//  │ for { }            │ while True:        │ Infinite loop (+ break)      │
	//  │ for _, v := range  │ for v in iterable  │ Iterate over collections     │
	//  │ for i, v := range  │ for i,v in enum()  │ Iterate with index           │
	//  │ label: for {break} │ (no direct equiv)  │ Break from nested loops      │
	//  └────────────────────┴────────────────────┴──────────────────────────────┘
	//
	// Bottom line: Go does MORE with ONE keyword (for) than Python
	// does with TWO (for + while). Go adds labeled breaks for nested
	// loops (which Python lacks). But Go has no else clause on loops
	// (Python's for...else and while...else).
	// ============================================================
}
