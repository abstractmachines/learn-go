# learn-go
Learning golang ... WIP

## Go Data Structures and errata

- [https://research.swtch.com/godata](https://research.swtch.com/godata)

- [Strings](./strings/README.md)
- [Slices, data marshalling ...](./slices/README.md)
- [Structs](./structs/README.md)
- [Runes (WIP)](./runes/README.md)

## Go iteration with slices, structs, more
- [Iteration with slices and structs](./iteration-slice-struct/README.md)

## Go Typing System, var, const ...
- [Typing System, var, const)](./typing-system/README.md)

## Go Modules, examples, package management system, exports and identifiers
- [Modules, Packages, Exports, Encapsulation](./module-example/README.md)

## WIP : function signatures

Here's a Go function signature that takes in a value of any type `v interface {}`,
and returns a slice of bytes and an error (from package json):
```
func Marshal(v interface{}) ([]byte, error)
```


