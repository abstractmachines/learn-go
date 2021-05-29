# Go Typing System

> Note that you may often receive the same address for data that's supposed to be assigned to a new memory location because of immutability -> creation of new data.

This is expected in Go.

## Mutable Go objects

- arrays and slices
- maps
- channels
- closures which are capturing at least 1 variable from the outer scope

## Immutable Go objects
- interfaces
- booleans, numeric values (including values of type int)
- strings
- pointers
- function pointers, and closures which can be reduced to function pointers
