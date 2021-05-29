# Go Encapsulation 

## Set privacy by exporting, not by `public` or `private` keyword

> Case determines what can be exported and what cannot.

```
type StructyPublicCapitalized struct {
	PublicField string
	privateField string
}
```
