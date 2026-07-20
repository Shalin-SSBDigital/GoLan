// ============================================================
//  CONSTANTS IN GO — with Python Comparisons
//  Constants are fixed values that CANNOT be changed at runtime
// ============================================================
// Syntax: const CONSTNAME type = value
//
// Key rules:
//   - Use 'const' keyword (not 'var')
//   - Value must be known at compile time (no function calls)
//   - Reassignment is a COMPILE ERROR — constants are immutable
//   - Naming convention: PascalCase for exported, camelCase for
//     unexported (but ALL_CAPS is also common for constants)
//   - Types: numeric (int, float64), string, bool only
//
// ┌──────────────────────┬────────────────────────────────────────────┐
// │        Go            │                 Python                     │
// ├──────────────────────┼────────────────────────────────────────────┤
// │ const Pi float64     │ # No REAL constants in Python              │
// │     = 3.14           │ Pi = 3.14  # convention: ALL_CAPS          │
// │                      │                                            │
// │ const is COMPILE-    │ Python has NO compile-time constants       │
// │ TIME — the value     │ EVERYTHING can be reassigned.              │
// │ is baked into the    │ ALL_CAPS is just a CONVENTION — the        │
// │ binary               │ language does not enforce it.              │
// │                      │                                            │
// │ const is TRULY       │ You can reassign Pi = 2 at any time        │
// │ immutable — the      │ and Python won't complain.                 │
// │ compiler ENFORCES it │ Go will REJECT reassignment at compile     │
// │                      │ time.                                      │
// └──────────────────────┴────────────────────────────────────────────┘
// ============================================================

package variables

// ============================================================
//  CONSTANT DECLARATION
// ============================================================
// Pi is an exported constant (capital P) — accessible from
// other packages. Its type is float64 and value is 3.14.
//
// Trying to reassign:  Pi = 2
// would produce:       cannot assign to Pi (neither addressable
//                      nor a map index expression)
//
// Python:
//   Pi = 3.14           # just a variable named Pi
//   Pi = 2              # NO ERROR — Python allows this!
//   # The convention is ALL_CAPS for "constants" but nothing
//   # enforces it. Use MyPy's Final type if you want enforcement.
//   from typing import Final
//   PI: Final[float] = 3.14  # MyPy will flag reassignment
//   # But this is still a runtime variable, not compile-time.
// ============================================================
const Pi float64 = 3.14

// ============================================================
//  COMMON CONSTANT PATTERNS (Python comparison)
// ============================================================
// Go:
//   const (
//       StatusOK    = 200
//       StatusNotFound = 404
//       StatusError    = 500
//   )
//
// Python:
//   from enum import IntEnum
//   class Status(IntEnum):
//       OK = 200
//       NOT_FOUND = 404
//       ERROR = 500
//   # Python's IntEnum is MORE flexible than Go's const block
//   # (methods, iteration, type safety), but Go's const is
//   # simpler and more performant (no runtime overhead).
//
// Iota — auto-incrementing constant generator:
//   const (
//       Monday = iota  // 0
//       Tuesday        // 1
//       Wednesday      // 2
//   )
//
// Python:
//   from enum import IntEnum, auto
//   class Day(IntEnum):
//       MONDAY = auto()      # 1 (auto() starts at 1)
//       TUESDAY = auto()     # 2
//       WEDNESDAY = auto()   # 3
//   # Python's auto() starts at 1, Go's iota starts at 0.
//   # Both auto-increment, but Go's syntax is more compact.
//
// ┌──────────────────────┬──────────────────────┬──────────────────────────┐
// │ Feature              │ Go (const)           │ Python (convention)      │
// ├──────────────────────┼──────────────────────┼──────────────────────────┤
// │ Enforced immutability│ YES (compile-time)    │ NO (convention only)     │
// │                      │                       │ typing.Final + MyPy     │
// │ Compile-time value   │ YES                   │ NO (runtime)             │
// │ Types allowed        │ Numbers, strings,     │ Any type (but List, Dict│
// │                      │ bools only            │ can be mutated even if   │
// │                      │                       │ "constant")              │
// │ Block declaration    │ const ( ... )         │ No block form            │
// │ Auto-increment       │ iota                  │ enum.auto()              │
// │ Performance          │ Zero-cost (baked in)  │ Runtime variable         │
// │ Grouping             │ const block           │ enum class or module var │
// └──────────────────────┴──────────────────────┴──────────────────────────┘
//
// Bottom line: Go constants are TRULY immutable — enforced at compile
// time. Python has NO real constants — only naming conventions and
// third-party type checkers. Go's iota is a compact auto-increment
// generator; Python's enum.auto() serves a similar purpose but is
// more verbose.
// ============================================================
