// ============================================================
//  CONCURRENCY IN GO
// ============================================================
//
//  ┌─────────────────────────────────────────────────────────────┐
//  │  CONCURRENCY vs PARALLELISM                                  │
//  │                                                             │
//  │  CONCURRENCY = DEALING with many things AT ONCE             │
//  │  (structure, design pattern — one person juggling tasks)    │
//  │                                                             │
//  │  PARALLELISM = DOING many things AT ONCE                    │
//  │  (execution, hardware — multiple people, each doing one     │
//  │   thing simultaneously)                                     │
//  │                                                             │
//  │  Go is designed for CONCURRENCY.                            │
//  │  If hardware has multiple cores, concurrency CAN become     │
//  │  parallelism. Go doesn't guarantee it — the runtime decides.│
//  └─────────────────────────────────────────────────────────────┘
//
//  ┌───────────────┬──────────────────────────────────────────────────┐
//  │ Go            │ Python                                           │
//  ├───────────────┼──────────────────────────────────────────────────┤
//  │ Goroutines    │ threading (heavy), asyncio (cooperative)         │
//  │ Channel       │ queue.Queue, asyncio.Queue                       │
//  │ select        │ asyncio.wait() with FIRST_COMPLETED              │
//  │ sync.Mutex    │ threading.Lock                                   │
//  │ sync.WaitGroup│ threading.Thread.join() (one at a time)          │
//  │               │                                                  │
//  │ KEY: Go makes │ Python GIL prevents true parallel CPU work       │
//  │ concurrency   │ asyncio needs explicit await everywhere          │
//  │ SIMPLE.       │ Go = "just add go keyword"                      │
//  └───────────────┴──────────────────────────────────────────────────┘
// ============================================================

package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================================
//  THE LEGO ASSEMBLY LINE ANALOGY
//
//  ┌─────────────────────────────────────────────────────────┐
//  │  SEQUENTIAL (no concurrency):                           │
//  │   Alex builds: [Base] → [Wheels] → [Body] → [Roof]     │
//  │   (one step at a time, total = 4 minutes)               │
//  │                                                         │
//  │  CONCURRENT (goroutines + channels):                    │
//  │   Alex builds base ──┐                                  │
//  │                      ├─→ Sam attaches wheels            │
//  │   Sam prepares      ─┘   (as base arrives)             │
//  │   wheels              → Alex attaches body              │
//  │                      → Sam adds roof                    │
│                                                         │
//  │   Each person works when they HAVE work.               │
│   Total ~2 minutes (overlap!)                           │
//  │                                                         │
//  │  PARALLEL (multiple cores):                             │
//  │   Alex AND Sam AND Jamal all build cars AT THE SAME     │
//  │   TIME. Real parallelism needs multiple CPU cores.      │
//  └─────────────────────────────────────────────────────────┘
// ============================================================

// ============================================================
//  1. BASIC GOROUTINE — "go" keyword
//  Go's simplest concurrency primitive.
//  Think: "Hey Go, run this in the background while I continue"
// ============================================================

func cookRice() {
	fmt.Println("  🍚 Cooking rice (5 min)...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("  🍚 Rice done!")
}

func cutVegetables() {
	fmt.Println("  🥕 Cutting vegetables (3 min)...")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("  🥕 Veggies cut!")
}

// ============================================================
//  2. CHANNEL COMMUNICATION
//  sync.WaitGroup + channel = coordinated teamwork
// ============================================================

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("  👷 Worker %d: Processing job %d...\n", id, job)
		time.Sleep(200 * time.Millisecond)
		results <- job * 2 // Send result back
	}
}

// ============================================================
//  3. SELECT STATEMENT — Like a multiplexer
//  Wait on MULTIPLE channels at once.
//  The first channel that's ready wins (like a race).
//  Think: TV remote, press whichever button responds first
// ============================================================

func fetchFromSource(name string, ch chan<- string, delay time.Duration) {
	time.Sleep(delay)
	ch <- name
}

// ============================================================
//  4. MUTEX — Protect shared data
//  When goroutines share a variable, only one can touch it
//  at a time. Mutex = "exclusive access token"
//  Think: One bathroom key — only one person at a time
// ============================================================

type SharedCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SharedCounter) Increment() {
	c.mu.Lock()
	c.value++ // Only one goroutine here at a time!
	c.mu.Unlock()
}

func (c *SharedCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	fmt.Println("========================================")
	fmt.Println("  🧵  CONCURRENCY IN GO 🧵")
	fmt.Println("========================================")
	fmt.Println()

	// ─────────────────────────────────────────────────
	//  SCENARIO 1: Sequential vs Concurrent
	// ─────────────────────────────────────────────────
	fmt.Println("─── Scenario 1: Sequential vs Concurrent ───")
	fmt.Println()

	fmt.Println("  SEQUENTIAL (one at a time):")
	start := time.Now()
	cookRice()
	cutVegetables()
	fmt.Printf("  ⏱️  Total: %v\n\n", time.Since(start).Round(time.Millisecond))

	fmt.Println("  CONCURRENT (overlap with goroutines):")
	start = time.Now()
	go cookRice()
	go cutVegetables()
	fmt.Println("  👤 Main didn't wait! Both started in background.")
	time.Sleep(600 * time.Millisecond) // Wait for both to finish
	fmt.Printf("  ⏱️  Total: %v (faster!)\n", time.Since(start).Round(time.Millisecond))
	fmt.Println()

	// ─────────────────────────────────────────────────
	//  SCENARIO 2: Fan-out / Fan-in pattern
	//  Multiple workers handle jobs from a queue
	//  This is the classic concurrency pattern in Go
	// ─────────────────────────────────────────────────
	fmt.Println("─── Scenario 2: Fan-out / Fan-in ───")
	fmt.Println()

	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// Start 3 workers (Fan-out)
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send 5 jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // No more jobs — workers will exit after processing remaining

	// Wait for all workers to finish, then close results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results (Fan-in)
	fmt.Println("  📊 Results:")
	for result := range results {
		fmt.Printf("    → %d\n", result)
	}
	fmt.Println()

	// ─────────────────────────────────────────────────
	//  SCENARIO 3: Select — wait on multiple channels
	//  Like a TV remote: press whichever responds first
	// ─────────────────────────────────────────────────
	fmt.Println("─── Scenario 3: Select — race between sources ───")
	fmt.Println()

	ch1 := make(chan string)
	ch2 := make(chan string)

	go fetchFromSource("Cache", ch1, 200*time.Millisecond)
	go fetchFromSource("Database", ch2, 400*time.Millisecond)

	select {
	case msg := <-ch1:
		fmt.Printf("  🏆 Cache responded first: %s\n", msg)
	case msg := <-ch2:
		fmt.Printf("  🏆 Database responded first: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("  ⏰ Timeout! Both sources too slow")
	}

	// Wait a bit so goroutines can clean up
	time.Sleep(200 * time.Millisecond)
	fmt.Println()

	// ─────────────────────────────────────────────────
	//  SCENARIO 4: Mutex — shared counter
	//  Without mutex, two goroutines writing the same
	//  variable create a DATA RACE (corrupted value)
	// ─────────────────────────────────────────────────
	fmt.Println("─── Scenario 4: Mutex — safe shared counter ───")
	fmt.Println()

	counter := SharedCounter{}

	var wg2 sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			counter.Increment()
		}()
	}
	wg2.Wait()
	fmt.Printf("  🔢 Counter (1000 concurrent increments): %d\n", counter.Value())
	fmt.Println("  ✅ No data races — mutex kept increments safe!")
	fmt.Println()

	// ─────────────────────────────────────────────────
	//  SCENARIO 5: Concurrency visualization
	// ─────────────────────────────────────────────────
	fmt.Println("─── Scenario 5: Concurrency vs Parallelism Vis ───")
	fmt.Println()
	fmt.Println("  CONCURRENCY (Go's design):")
	fmt.Println("  ┌─────┐   ┌─────┐   ┌─────┐")
	fmt.Println("  │Task1│──→│Task2│──→│Task3│  (one goroutine")
	fmt.Println("  └─────┘   └─────┘   └─────┘   switches tasks)")
	fmt.Println("     ↓          ↓          ↓")
	fmt.Println("   [goroutine] [goroutine] [goroutine]")
	fmt.Println()
	fmt.Println("  PARALLELISM (hardware):")
	fmt.Println("  ┌─────┐                   ")
	fmt.Println("  │Task1│  [Core 0]          ")
	fmt.Println("  └─────┘                   ")
	fmt.Println("  ┌─────┐                   ")
	fmt.Println("  │Task2│  [Core 1]  (truly simultaneous)")
	fmt.Println("  └─────┘                   ")
	fmt.Println("  ┌─────┐                   ")
	fmt.Println("  │Task3│  [Core 2]          ")
	fmt.Println("  └─────┘                   ")

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("  🧵  CONCURRENCY DEMO COMPLETE 🧵")
	fmt.Println("========================================")
}

// ============================================================
//  KEY CONCEPTS SUMMARY
// ============================================================
//
//  ┌──────────────────────┬────────────────────────────────────┐
//  │ Concept              │ Meaning                            │
//  ├──────────────────────┼────────────────────────────────────┤
//  │ go func()            │ Launch goroutine (lightweight      │
//  │                      │ concurrent function)               │
//  ├──────────────────────┼────────────────────────────────────┤
//  │ channel              │ Typed pipe between goroutines      │
//  │                      │ ch <- val (send), <-ch (receive)  │
//  ├──────────────────────┼────────────────────────────────────┤
//  │ sync.WaitGroup       │ Wait for N goroutines to finish    │
//  │                      │ Add(N) → Done() N times → Wait()  │
//  ├──────────────────────┼────────────────────────────────────┤
//  │ select {             │ Wait on multiple channels at once  │
//  │ case <-ch1: ...      │ First ready case runs              │
//  │ case <-ch2: ...      │ Like asyncio.wait(FIRST_COMPLETED) │
//  │ }                    │                                    │
//  ├──────────────────────┼────────────────────────────────────┤
//  │ sync.Mutex           │ Exclusive access to shared data    │
//  │                      │ Lock() / Unlock()                 │
//  ├──────────────────────┼────────────────────────────────────┤
//  │ Fan-out / Fan-in     │ Distribute work to multiple        │
//  │                      │ workers, collect results           │
//  ├──────────────────────┼────────────────────────────────────┤
//  │ Concurrency          │ STRUCTURE: dealing with many       │
//  │                      │ things at once (design)            │
//  ├──────────────────────┼────────────────────────────────────┤
//  │ Parallelism          │ EXECUTION: doing many things       │
//  │                      │ at once (hardware)                 │
//  └──────────────────────┴────────────────────────────────────┘
//
//  ⚠️ Go Motto:
//  "Do not communicate by sharing memory;
//   instead, share memory by communicating."
// ============================================================