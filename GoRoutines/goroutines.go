// ============================================================
//  GOROUTINES IN GO — Lightweight Concurrent Functions
// ============================================================
//
//  ┌───────────────────┬──────────────────────────────────────┐
//  │      Go           │              Python                   │
//  ├───────────────────┼──────────────────────────────────────┤
//  │ go fn()           │ import threading                     │
//  │ go func() { ... } │ t = threading.Thread(target=fn)     │
//  │                   │ t.start()                             │
//  │                   │                                       │
//  │ Goroutines are    │ Python threads are OS threads.        │
//  │ multiplexed onto  │ Heavy (MB stack). GIL limits CPU.     │
//  │ OS threads.       │                                       │
//  │ Stack starts at   │                                       │
//  │ ~2 KB, grows as   │                                       │
//  │ needed.           │                                       │
//  └───────────────────┴──────────────────────────────────────┘
//
// ============================================================

package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================================
//  1. BASIC GOROUTINE — "go" keyword
//  Think: Alex starts building WITHOUT waiting for Sam.
//  "go" = "start this, don't wait for it to finish"
// ============================================================

func sayHello() {
	fmt.Println("  👋 Hello from a goroutine!")
}

// ============================================================
//  2. WAITGROUP — Wait for all goroutines to finish
//  Think: Alex and Sam both building. Wait for BOTH to finish.
//  Without WaitGroup, main exits before goroutines run.
// ============================================================

func buildSpaceship(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this task as done when function exits
	fmt.Printf("  🚀 Builder %d: Building spaceship...\n", id)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("  🚀 Builder %d: Spaceship complete!\n", id)
}

// ============================================================
//  3. GOROUTINE CLOSURE CAPTURE — Common Gotcha
//  If you launch goroutines in a loop, the loop variable
//  changes while goroutines run. Always COPY the variable.
// ============================================================

// ============================================================
//  4. ANONYMOUS GOROUTINE — Inline function
//  Like an arrow function in JS / lambda in Python
// ============================================================

func main() {
	fmt.Println("========================================")
	fmt.Println("  🧵  GOROUTINES DEMO 🧵")
	fmt.Println("========================================")
	fmt.Println()

	// ─────────────────────────────────────────────────
	//  SCENARIO 1: Basic goroutine
	//  Starting a function with "go" keyword
	//  The goroutine runs CONCURRENTLY with main
	// ─────────────────────────────────────────────────
	fmt.Println("─── Scenario 1: Basic goroutine ───")
	fmt.Println()

	// Start sayHello as a goroutine — it runs in background
	go sayHello()

	// Main continues immediately WITHOUT waiting for sayHello
	fmt.Println("  👤 Main: 'I started a goroutine! Continuing...'")
	fmt.Println("  👤 Main: 'But I need to sleep so the goroutine can finish'")
	time.Sleep(100 * time.Millisecond) // Give goroutine time to run
	fmt.Println()

	// ─────────────────────────────────────────────────
	//  SCENARIO 2: Multiple goroutines with WaitGroup
	//  WaitGroup = team leader counting hands raised
	//  wg.Add(N) = N people on the team
	//  wg.Done() = "I'm done!" (lowers one hand)
	//  wg.Wait() = wait until ALL hands are down
	// ─────────────────────────────────────────────────
	fmt.Println("─── Scenario 2: Multiple builders with WaitGroup ───")
	fmt.Println()

	var wg sync.WaitGroup

	// Start 3 builders working simultaneously
	for i := 1; i <= 3; i++ {
		wg.Add(1)                     // "One more worker joining"
		go buildSpaceship(i, &wg)     // Start builder in background
	}

	// Main waits HERE until all 3 call wg.Done()
	wg.Wait()
	fmt.Println("  👤 Main: 'All 3 builders finished! Team complete!'")
	fmt.Println()

	// ─────────────────────────────────────────────────
	//  SCENARIO 3: Loop closure gotcha
	//  WRONG: goroutines capture the SAME loop variable
	//  RIGHT: pass it as a parameter or copy inside loop
	// ─────────────────────────────────────────────────
	fmt.Println("─── Scenario 3: Loop closure gotcha ───")
	fmt.Println()

	fmt.Println("  ❌ WRONG way — all goroutines see LAST value:")
	var wgWrong sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wgWrong.Add(1)
		go func() {
			defer wgWrong.Done()
			// BUG: i is captured by reference — by the time
			// this goroutine runs, i has moved on!
			// In Go 1.22+, this is fixed. In older versions,
			// all 3 print "Builder 3"!
			fmt.Printf("    Builder %d\n", i)
		}()
	}
	wgWrong.Wait()

	fmt.Println()
	fmt.Println("  ✅ RIGHT way — pass variable as argument:")
	var wgRight sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wgRight.Add(1)
		go func(id int) {  // Copy i by passing as argument
			defer wgRight.Done()
			fmt.Printf("    Builder %d\n", id)
		}(i)  // Pass CURRENT value of i
	}
	wgRight.Wait()
	fmt.Println()

	// ─────────────────────────────────────────────────
	//  SCENARIO 4: Anonymous goroutine inline
	//  Like Python's lambda but with "go" keyword
	// ─────────────────────────────────────────────────
	fmt.Println("─── Scenario 4: Anonymous goroutine (inline) ───")
	fmt.Println()

	var wg2 sync.WaitGroup

	wg2.Add(1)
	go func() {
		defer wg2.Done()
		fmt.Println("  🏃 Anonymous goroutine: 'I have no name but I work!'")
		time.Sleep(50 * time.Millisecond)
	}()

	wg2.Add(1)
	go func(name string) {
		defer wg2.Done()
		fmt.Printf("  🏃 %s: 'I have a parameter!' \n", name)
	}("NamedRunner")

	wg2.Wait()
	fmt.Println()

	// ─────────────────────────────────────────────────
	//  SCENARIO 5: Goroutines vs sequential
	//  Compare sequential vs concurrent execution time
	// ─────────────────────────────────────────────────
	fmt.Println("─── Scenario 5: Sequential vs Concurrent ───")
	fmt.Println()

	work := func(id int) {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("    Task %d done\n", id)
	}

	// Sequential — 3 tasks take 600ms
	start := time.Now()
	for i := 1; i <= 3; i++ {
		work(i)
	}
	fmt.Printf("  Sequential: %v\n", time.Since(start).Round(time.Millisecond))

	// Concurrent — 3 tasks take ~200ms (overlapped!)
	var wg3 sync.WaitGroup
	start = time.Now()
	for i := 1; i <= 3; i++ {
		wg3.Add(1)
		go func(id int) {
			defer wg3.Done()
			work(id)
		}(i)
	}
	wg3.Wait()
	fmt.Printf("  Concurrent: %v (3x faster!)\n", time.Since(start).Round(time.Millisecond))

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("  🧵  GOROUTINES DEMO COMPLETE 🧵")
	fmt.Println("========================================")
}

// ============================================================
//  KEY CONCEPTS SUMMARY
// ============================================================
//
//  ┌──────────────────┬──────────────────────────────────────┐
//  │ Concept          │ Explanation                          │
//  ├──────────────────┼──────────────────────────────────────┤
//  │ go fn()          │ Start fn() as a goroutine            │
//  │                  │ Returns immediately, runs in bg      │
//  ├──────────────────┼──────────────────────────────────────┤
//  │ go func() { }()  │ Start an anonymous goroutine         │
//  │                  │ Like lambda + start in Python        │
//  ├──────────────────┼──────────────────────────────────────┤
//  │ sync.WaitGroup   │ Counter for goroutine completion     │
//  │                  │ wg.Add(N) → wg.Done() → wg.Wait()   │
//  ├──────────────────┼──────────────────────────────────────┤
//  │ Goroutine stack  │ Starts at ~2KB, grows as needed      │
//  │                  │ Python thread stack: ~8MB fixed      │
//  ├──────────────────┼──────────────────────────────────────┤
//  │ Closure capture  │ Loop variables are shared by ref!    │
//  │                  │ Pass as arg to avoid bugs            │
//  ├──────────────────┼──────────────────────────────────────┤
//  │ main exits →     │ All goroutines DIE immediately       │
//  │ goroutines die   │ Use WaitGroup to wait for them       │
//  └──────────────────┴──────────────────────────────────────┘
//
//  ⚠️ Common mistakes:
//  1. Forgetting wg.Add(1) — program panics or skips wait
//  2. Forgetting wg.Done() — forever hang (deadlock!)
//  3. Not copying loop variables — all goroutines see last value
//  4. Main exits before goroutines finish — use WaitGroup, not sleep
// ============================================================