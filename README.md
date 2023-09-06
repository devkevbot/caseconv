# caseconv

`caseconv` is a library for converting strings to different cases.

More specfically, this library is a Go port of [this Rust-based string conversion library](https://github.com/devkevbot/str_case_conv).

## Documentation

Full documentation can be found on [pkg.go.dev](https://pkg.go.dev/github.com/devkevbot/caseconv#section-documentation).

## Features

This library provides functionality for converting strings to the following cases:

- `camelCase`
- `CONSTANT_CASE`
- `kebab-case`
- `PascalCase`
- `snake_case`
- `Start Case`

## Examples

A basic program which utilizes `caseconv` is as follows:

```go
package main

import (
	"fmt"

	"github.com/devkevbot/caseconv"
)

func main() {
	fmt.Println("camel case ->", caseconv.ToCamelCase("hello world"))
	fmt.Println("constant case ->", caseconv.ToConstantCase("hello world"))
	fmt.Println("kebab case ->", caseconv.ToKebabCase("hello world"))
	fmt.Println("pascal case ->", caseconv.ToPascalCase("hello world"))
	fmt.Println("snake case ->", caseconv.ToSnakeCase("hello world"))
	fmt.Println("start case ->", caseconv.ToStartCase("hello world"))
}
```

Which prints the following:

```
camel case -> helloWorld
constant case -> HELLO_WORLD
kebab case -> hello-world   
pascal case -> HelloWorld   
snake case -> hello_world   
start case -> Hello World
```
