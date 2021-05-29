# Iteration in Go

Slices, structs, ranges and more

## There is kinda only one type of iteration in Go: variations on the for loop

### for

```
for {
    ...
}
```

### for (C style loops)

```
for idx = 0; idx < 10; ++idx {
    ...
}
```

### for range (for each)

```
for idx, item := range ACollection.items {
    ...
}
```
