# 🎓 Thank Go Practice

[![Go](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![Course](https://img.shields.io/badge/Stepik-96832-green?logo=data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAyNCAyNCI+PHBhdGggZmlsbD0id2hpdGUiIGQ9Ik0xMiAyQzYuNDggMiAyIDYuNDggMiAxMnM0LjQ4IDEwIDEwIDEwIDEwLTQuNDggMTAtMTBTMTcuNTIgMiAxMiAyem0wIDE4Yy00LjQxIDAtOC0zLjU5LTgtOHMzLjU5LTggOC04IDggMy41OSA4IDgtMy41OSA4LTggNHoiLz48L3N2Zz4=)](https://stepik.org/course/96832)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Practice solutions for the [**"Thank Go!"**](https://stepik.org/course/96832) Stepik course — a hands-on journey through Go from basics to stdlib mastery.

## 🗂 Structure

Each section maps to a course module. Every exercise lives in its own package with a `main.go` containing the task description and solution.

```
├── section1-basics/
│   ├── basic-constructs/     # variables, types, control flow
│   ├── arrays-maps/          # slices, maps, strings
│   ├── functions-pointers/   # functions, closures, pointers
│   └── structs-methods/      # structs, methods, embedding
├── section2-basics2/
│   ├── interfaces/           # interfaces, polymorphism
│   ├── errors/               # error handling, wrapping
│   ├── iota/                 # enums with iota
│   └── generics/             # type parameters
├── section3-clean-code/
│   ├── packages-modules/     # project layout, modules
│   ├── tests/                # testing, table-driven tests
│   ├── benchmarks/           # performance benchmarks
│   └── profiler/             # pprof, profiling
├── section4-concurrency/
│   ├── goroutines/           # goroutines basics
│   ├── channels1/            # channel fundamentals
│   ├── channels2/            # advanced channel patterns
│   ├── sync/                 # sync primitives
│   ├── context/              # context propagation
│   ├── time/                 # timers, tickers
│   └── composition/          # concurrency composition
├── section5-stdlib/
│   ├── io/                   # readers, writers, streams
│   ├── text/                 # strings, regexp, templates
│   ├── json-xml-csv/         # encoding/decoding
│   ├── datetime/             # time, duration, formatting
│   ├── http/                 # HTTP client & server
│   └── sql/                  # database/sql
└── section6-final/           # capstone project
```

## 🚀 Progress

| Section | Topics | Exercises | Status |
|---------|--------|-----------|--------|
| **1. Basics** | Variables, types, control flow, slices, maps, functions, structs | 12 | ✅ In progress |
| ↳ basic-constructs | Variables, types, control flow | 4 | ✅ Done |
| ↳ arrays-maps | Slices, maps, strings | 3 | ✅ Done |
| ↳ functions-pointers | Functions, closures, pointers | 3 | ✅ Done |
| ↳ structs-methods | Structs, methods, embedding | 2 | ✅ Done |
| **2. Basics II** | Interfaces, errors, iota, generics | 5 | ✅ In progress |
| ↳ interfaces | Interfaces, polymorphism | 2 | ✅ Done |
| ↳ errors | Error handling, wrapping, custom types | 3 | ✅ Done |
| **3. Clean Code** | Packages, testing, benchmarks, profiling | 0 | ⬜ Not started |
| **4. Concurrency** | Goroutines, channels, sync, context | 0 | ⬜ Not started |
| **5. Stdlib** | IO, text, encoding, datetime, HTTP, SQL | 0 | ⬜ Not started |
| **6. Final** | Capstone project | 0 | ⬜ Not started |

## 🛠 Run

```bash
# Run a specific exercise
go run section1-basics/basic-constructs/duration/main.go

# Pass input via pipe
echo "1h30m" | go run section1-basics/basic-constructs/duration/main.go

# Run all tests
go test -race ./...

# Build all exercises
go build ./...
```

## 📝 License

MIT
