# Go structs

## References: 
- [Go 101 Structs](https://go101.org/article/struct.html)
- [A Tour of Go Structs at golang.org](https://tour.golang.org/moretypes/2)
- [Golang Book: Structs and Interfaces](https://www.golang-book.com/books/intro/9)

## Struct Value Literals and Struct Manipulations

Go Structs are interesting because we have a certain flexibility with them.

```
type S struct {
	x int
	y bool
}
```

> Optional fields, default values, and zero values

- See [structs.go](./structs.go)   

- Go structs have zero values (zeroing out with default values if not explicitly defined when invoked):
```
type newS = S{}
```
- Optional fields (and order does not matter):
```
type newS = {x:1} // y is false by default.
type newNewS = {y:true, x:1}
```

> Tips and practices for working with Go structs

>> Avoid using a brittle approach with set number of fields.

- What if I use:
```
type TPayne struct {
    Nickname string
    Discography string
}
```
- and later invoke it with:
```
var bebehT = TPayne {'t', 'lots of records'}
```
- But then the developer of TPayne changes the struct to:
```
type TPayne struct {
    Nickname string
    Discography string
    Legitness int
}
```

- ^^ that means that TPayne can't add a field later, or this will break.
- So, instead, this won't break, as struct fields are optional, and have default values which vary per type.
```
var bebehT = TPayne{Nickname: 't', Discography: 'lots of records'}
```