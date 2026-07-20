// ============================================================
//  MAPS IN GO — with Python Comparisons
// ============================================================
//  Maps are Go's built-in key-value data structure (hash table).
//
// ┌──────────────────────┬────────────────────────────────────────────┐
// │        Go            │                 Python                     │
// ├──────────────────────┼────────────────────────────────────────────┤
// │ map[string]int       │ dict — key-value mapping                    │
// │                      │                                            │
// │ SYNTAX: make(map[K]V)│ SYNTAX: {} or dict()                       │
// │                      │                                            │
// │ m["k"] = 7           │ d["k"] = 7  (same syntax)                  │
// │                      │                                            │
// │ delete(m, "k")       │ del d["k"]  (similar)                      │
// │                      │                                            │
// │ v, ok := m["k"]      │ d.get("k", default) — check presence       │
// │ "comma ok" idiom     │   "k" in d — check only                    │
// │ separate return      │   d["k"] — raises KeyError if missing       │
// │ = two-value lookup   │                                            │
// └──────────────────────┴────────────────────────────────────────────┘
//
// KEY DIFFERENCES:
//   1. Go maps are TYPED — all keys must be the same type, all values
//      must be the same type. Python dicts can mix key/value types freely.
//   2. Accessing a missing key in Go returns the ZERO VALUE (no error).
//      Python raises KeyError. Use the "comma ok" idiom in Go to check.
//   3. Go has no dictionary comprehension. Use a for loop instead.
//   4. Go maps are REFERENCE types (like slices) — assigning copies
//      the reference, not the data. Python dicts work the same way.
// ============================================================

package main

import (
	"fmt"
	"maps"
)

func main_() {

	// ============================================================
	//  CREATING A MAP
	//  Syntax: make(map[KeyType]ValueType)
	// ============================================================
	// make() creates an empty map and initializes its internal
	// hash table. The map is ready to use immediately.
	//
	// Python:
	//   m = dict()           # creates empty dict
	//   m = {}               # same, more common
	//   # Python has TWO syntaxes; Go has make() plus one literal form.
	// ============================================================
	m := make(map[string]int)

	// ============================================================
	//  SET & GET ELEMENTS
	// ============================================================
	// Setting and getting values uses the same syntax as Python.
	//
	// Python:
	//   m["k1"] = 7
	//   v1 = m["k1"]
	// ============================================================
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)
	// Output: map: map[k1:7 k2:13]

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	// Output: v1: 7

	// ============================================================
	//  ACCESSING A MISSING KEY — Zero Value
	// ============================================================
	// If the key doesn't exist, Go returns the ZERO VALUE for
	// the value type (0 for int, "" for string, false for bool).
	// It does NOT panic or error.
	//
	// Python:
	//   v3 = m["k3"]                # KeyError! Python panics.
	//   v3 = m.get("k3", 0)         # returns 0 (safely)
	//   v3 = m.get("k3")            # returns None (safely)
	//   # Python requires .get() for safe access. Go gives zero
	//   # value by default. This is both convenient AND dangerous —
	//   # you won't get an error for a typo in the key!
	// ============================================================
	v3 := m["k3"]
	fmt.Println("v3:", v3)
	// Output: v3: 0
	// ⚠️ Note: this prints 0, but we never set m["k3"]!
	// In Python, this would raise KeyError and crash.
	// In Go, it silently returns the zero value.

	// ============================================================
	//  len() — number of key-value pairs
	// ============================================================
	// Python: len(m)  — same function name!
	// ============================================================
	fmt.Println("len:", len(m))
	// Output: len: 2

	// ============================================================
	//  delete() — remove a key-value pair
	//  Syntax: delete(map, key)
	// ============================================================
	// delete() is a BUILT-IN function, not a method on the map.
	// Deleting a key that doesn't exist is a no-op (no error).
	//
	// Python:
	//   del m["k2"]           # deletes the key
	//   del m["nonexistent"]  # KeyError! Python raises error
	//   m.pop("k2")           # returns value AND deletes
	//   m.pop("k2", None)     # safe version with default
	//   # Go's delete() on missing key: no-op (safe, no error).
	//   # Python's del on missing key: KeyError.
	// ============================================================
	delete(m, "k2")
	fmt.Println("map:", m)
	// Output: map: map[k1:7]

	// ============================================================
	//  clear() — remove ALL entries (Go 1.21+)
	// ============================================================
	// clear() empties the map completely. After clear, len(m) == 0.
	// The map itself is still usable — you can add new keys.
	//
	// Python:
	//   m.clear()             # same! dict.clear() method
	//   m = {}                # creates a NEW empty dict (new object)
	// ============================================================
	clear(m)
	fmt.Println("map:", m)
	// Output: map: map[]

	// ============================================================
	//  THE "COMMA OK" IDIOM — checking if a key exists
	//  Syntax: value, ok := map[key]
	// ============================================================
	// When you access a map with TWO return values:
	//   - value: the value (or zero value if key is missing)
	//   - ok:    bool — true if the key existed, false if not
	//
	// This is the SAFE way to check for key presence, because
	// the single-value access silently returns the zero value
	// (which might be a valid stored value!).
	//
	// Python:
	//   if "k2" in m:                    # check presence only
	//       v = m["k2"]
	//   v = m.get("k2")                  # returns None if missing
	//   v = m.get("k2", default_value)   # with default
	//   # Python's "in" and .get() separate the concerns.
	//   # Go's "comma ok" combines them in one operation.
	// ============================================================
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	// Output: prs: false

	// ============================================================
	//  MAP LITERAL — declare and initialize
	//  Syntax: map[K]V{key1: val1, key2: val2, ...}
	// ============================================================
	// Creates a map with two pre-populated key-value pairs.
	// No make() needed — the literal allocates the hash table.
	//
	// Python:
	//   n = {"foo": 1, "bar": 2}    # same syntax!
	//   # Python uses the exact same curly-brace syntax.
	// ============================================================
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
	// Output: map: map[bar:2 foo:1]

	// ============================================================
	//  maps.Equal() — compare two maps (Go 1.21+)
	//  Package: "maps"
	// ============================================================
	// Maps CANNOT be compared with == (compile error).
	// Use maps.Equal() which compares key-by-key, value-by-value.
	//
	// Python:
	//   n == n2     # True! Python dicts support == directly.
	//   # Go deliberately omits == for maps because "what is
	//   # map equality?" is ambiguous (same references? same entries?).
	//   # maps.Equal() makes it explicit: value equality.
	// ============================================================
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
	// Output: n == n2

	// ============================================================
	//  ITERATING OVER A MAP — UNORDERED!
	// ============================================================
	// for k, v := range m { ... } iterates over all key-value pairs.
	// The order is DELIBERATELY RANDOMIZED in Go — you get a
	// different order each time (to prevent reliance on order).
	//
	// Python:
	//   for k, v in m.items():       # method call, not range
	//       print(k, v)
	//   # Python 3.7+ preserves INSERTION ORDER.
	//   # This is a FUNDAMENTAL difference:
	//   #   Go:  intentionally random order
	//   #   Py:  guaranteed insertion order
	//   # NEVER rely on map iteration order in Go!
	// ============================================================
	fmt.Println("Iterating n (order will vary!):")
	for k, v := range n {
		fmt.Printf("  %s -> %d\n", k, v)
	}

	// ============================================================
	//  SUMMARY: Go Maps vs Python Dicts
	// ============================================================
	//  ┌────────────────────┬──────────────────────┬──────────────────────────┐
	//  │ Feature            │ Go                   │ Python                    │
	//  ├────────────────────┼──────────────────────┼──────────────────────────┤
	//  │ Type safety        │ Fully typed (K, V)   │ Dynamic (any key/value)  │
	//  │ Creation           │ make(map[K]V)        │ {} or dict()             │
	//  │ Missing key        │ Returns ZERO value   │ KeyError (or .get())     │
	//  │ Key existence      │ v, ok := m[k]        │ "k" in m                 │
	//  │ Delete             │ delete(m, k)         │ del d[k]                 │
	//  │ Iteration order    │ RANDOM (intentional) │ Insertion order (3.7+)   │
	//  │ Comparison         │ maps.Equal(a, b)     │ == operator works        │
	//  │ Comprehension      │ Not supported        │ {k:v for k in ...}       │
	//  │ Reference type     │ Yes (assign = share) │ Yes (assign = share)     │
	//  │ Zero value         │ nil (usable! no make)│ None (unusable)          │
	//  └────────────────────┴──────────────────────┴──────────────────────────┘
	//
	// Bottom line: Go maps are typed hash tables with deliberate
	// non-deterministic iteration. Python dicts are dynamic mappings
	// with guaranteed insertion order. The syntax is similar, but
	// the safety guarantees are very different.
	// ============================================================
}
