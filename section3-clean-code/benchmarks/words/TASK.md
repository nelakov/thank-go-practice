# Optimize word search (section 3.3)

`Words` finds the index of a specific word in a string.

```go
s := "go is awesome, php is not"
w := MakeWords(s)

fmt.Println(w.Index("go"))      // 0
fmt.Println(w.Index("is"))      // 1
fmt.Println(w.Index("is not"))  // -1
fmt.Println(w.Index("python"))  // -1
```

The current `Words.Index()` allocates a lot. Each call re-splits the whole
string, so memory grows with the input:

```
Index-10:      314 ns      166 B      3 allocs
Index-100:     1917 ns     1798 B     3 allocs
Index-1000:    17398 ns    16390 B    3 allocs
Index-10000:   163084 ns   163847 B   3 allocs
```

## Goal

For strings of 1 to 10000 words, `Index` must use 1000 bytes or less. Also try
to keep its time from growing with the length of the string.

## Constraints

- You may change the fields of `Words` and the bodies of `MakeWords()` and
  `Index()`. Keep the signatures.
- `strings.Fields()` is a fine way to split the string into words.
- Benchmarks are added automatically by the checker (Stepik "Run code").

## Hint

The optimization target is `Words.Index()` specifically. If that means making
`MakeWords()` heavier, that is acceptable.
