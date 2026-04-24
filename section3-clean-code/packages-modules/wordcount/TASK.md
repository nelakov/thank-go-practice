# wordcount

CLI utility that counts the number of words in a sentence:

```
$ ./wordcount "go is awesome"
3
```

Whitespace is the only delimiter. A word is any sequence of characters not containing whitespace.

## Instructions

1. Clone `github.com/gothanks/wordcount` to your local machine.
2. Create a Go module with identifier `wordcount` (same way we did for `match`).
3. Implement the utility logic. Verify it counts words correctly.
4. Run tests: `go test`.
5. Tests print a sha1 hash — this is the Stepik answer code.

## Hint

`os.Args` helps read the input string. `strings.Fields` splits a sentence into words.

## Solution

Scaffolded as a separate Go module per course instructions:
[nelakov/wordcount](https://github.com/nelakov/wordcount)
