# GO Lang

- Go programming language is an open source project
- Its concurrency mechanisms make it easy to write programs that get the most out of multicore
- Go compiles quickly to machine code
- convenience of:
  - garbage collection
  - the power of run-time reflection
- fast, statically typed
- compiled language
- interpreted language

## Hello World

- Enable dependency tracking for your Module: is done by go.mod file, which tracks the deps of other modules in your module (just like package.json in node). Its package manager
- \$ touch go.mod
- \$ go mod init hello (go mod init command)
- \$ touch main.go
- \$ go run .

## Import a 3rd party Module

- All the 3rd party modules/libr are here- https://pkg.go.dev/
- To import 3rd party module
  - \$ go mod tidy (When imported rsc.io/quote)
  - go mod tidy, it located and downloaded the rsc.io/quote module that contains the package you imported. By default, it downloaded the latest version
  - Check there is go.sum file created automatically and also in go.mod file where you have direct & indirect deps/requires
- \$ go run import-fun.go

## Implement a module that others can use

- \$ mkdir greetings
- \$ cd greetings
- \$ go mod init example.com/greetings
  - Create the Package in greetings.go fie
  - Write your custom function in greetings.go file ==> Hello()
- \$ cd .. (root folder)
- \$ mkdir hello
- \$ cd hello
- \$ go mod init example.com/hello (inside create go.mod file)
- \$ touch hello.go
- \$ go mod edit -replace example.com/greetings=../greetings
- \$ go mod tidy (to synchronize the deps)
- \$ go run .

## Error handling

- Create a folder/module greetings2 which has go.mod and greetings.go
- \$ cd error-handling (for go module for error-handling, create file hello.go and go.mod)
- \$ go mod init tsabunkar.com/error-handling
- \$ go mod edit -replace tsabunkar.com/greetings2 => ../greetings2
- \$ go mod tidy (to synchronize the deps)
- \$ go run .

# Variable Casing Style

- Follow camelCase for private variables: Example: myVar, greetingMessage.
- Use PascalCase for exported variables: Example: MyPublicVar, ExportedFunction.
- Constants should use camelCase or PascalCase (depending on scope):
  - Private constant: defaultTimeout.
  - Public constant: MaxConnections.
- Avoid underscores: Go prefers camelCase for readability and consistency.
- Official Doc: https://go.dev/doc/effective_go#mixed-caps

# Go Slice Type

- Slice is a datatype
- A slice is like an array, except that its size changes dynamically as you add and remove items
- Defining an Array
  - `b := [2]string{"Penn", "Teller"}`
  - `b := [...]string{"Penn", "Teller"}`
- Defining an Slice
  - type specification for a slice is []T, where T is the type of the elements of the slice. Unlike an array type, a slice type has no specified length.
  - letters := []string{"a", "b", "c", "d"}
  - A slice can be created with the built-in function called make
    - func make([]T, len, cap) []T (T stands for the element type of the slice to be created)

# Return a random greeting

- Create greetings/greetings.go
- Create hello3/hello.go
- basics-go/greetings3 \$ go mod init xxx.com/greetings3
- basics-go/hello3 \$ go mod init example.com/hello3
- basics-go/hello3 \$ go mod edit -replace xxx.com/greetings3 => ../greetings3
- basics-go/hello3 \$ go mod tidy
- basics-go/hello3 \$ go run . (run multiple times)

# Unit Testing

-
