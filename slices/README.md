# Go slices

> Go slices are:
- [A pointer and a length](#A-pointer-and-a-length)
- [Immutable multi-word structures](#Immutable-multi-word-structures)

## A pointer and a length

[https://research.swtch.com/godata](https://research.swtch.com/godata)

> "A pointer and a length" is a common construct in Go. It makes slicing a lot easier to just move the pointer(s), and change the length.

## Immutable multi-word structures

### New data is created.

> "A string is represented in memory as a 2-word structure containing a pointer to the string data and a length. Because the string is immutable, it is safe for multiple strings to share the same storage, so slicing s results in a new 2-word structure with a potentially different pointer and length that still refers to the same byte sequence. This means that slicing can be done without allocation or copying, making string slices as efficient as passing around explicit indexes." [https://research.swtch.com/godata](https://research.swtch.com/godata)
