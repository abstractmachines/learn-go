## Create a Go Module
> Step 1. Make your module that you'll export. It must have its own directory.

- Make a new directory named `fakemodule` (and `cd` into that dir).

> Step 2: Create a module out of those files/in that directory.

Create a file `fakemodule.go`:
```
package fakemodule

import "fmt"

func Fakemodule() string {
    message: := 
}
```

- In that directory, `go mod init fakemodule`.

- This creates `go.mod` which is similar to `package.json` - dependency manager.

To conclude so far, your directory will look like: 
```
/yourProject
    /fakemodule
        fakemodule.go
        mod.go
```

> Step 3. Now, make another directory which will contain our `main()` function, import the 
other module, and build our code.
```
/yourProject
    /mainfakemodule ---> create this directory and cd into it.
    /fakemodule
        fakemodule.go
        mod.go
```

- Create this file in `/mainfakemodule` dir:
```
package main

import (
	"fakemodule"
	"fmt"
)

func main() {
	message := fakemodule.Fakemodule()
	fmt.Println(message)
}
```

- `go mod init mainfakemodule` in this directory to create module.

To conclude so far, your directory will look like: 
```
/yourProject
    /mainfakemodule
        mainfakemodule.go
        mod.go
    /fakemodule
        fakemodule.go
        mod.go
```

> Step 4. Now create package and dependency management in mod.go in our MAIN - meaning the dir that has our "main" module.

`mod.go` is a lot like `package.json` in JS projects. It will look like:
```
module mainfakemodule

go 1.15
```
In `mod.go`, add the line `replace fakemodule => ../fakemodule` to add that _dependency._
```
module mainfakemodule

go 1.15

replace fakemodule => ../fakemodule
```


> Step 5. go build to create a binary, and, to create a hash in `go.mod` to refer to _package version._

After running `go build` in our directory `/mainfakemodule`, `go.mod` will be:
```
module mainfakemodule

go 1.15

replace fakemodule => ../fakemodule

require fakemodule v0.0.0-00010101000000-000000000000
```

> Step 6. Run the binary/executable.

```
> ./mainfakemodule
```
