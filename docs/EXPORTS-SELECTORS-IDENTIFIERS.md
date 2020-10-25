# Golang exports, selectors, and identifiers

The Go package management ecosystem is a bit like that of npm.

Modules and packages are not exported explicitly. They are exported by inference via the identifier.

## Identifiers

> Case determines what can be exported and what cannot.

## Packages
- [all available golang packages, including source code](https://golang.org/pkg/)
- [GoDoc.org is a great place to search for Go packages too!](https://godoc.org/)

## Selectors and expressions vs "qualified (package) identifiers"

- `foo.bar` is a selector (or "selector expression") **if** it's describing a "local/type"
- `foo.Bar` is a _qualified identifier_, **if** `Foo` is an exported package.

There are also "depths of selectors" - which really just means, nested data.