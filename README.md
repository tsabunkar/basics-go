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

- Go's built-in support for unit testing
- `_test.go` tells the go test command that this file contains test functions
  - Added test file under /greetings3 folder
  - NOTE: [no test files] Error ==> Couldnt find the test file
    - fileName_test.go should be followed on the file which you are writting the test case
    - for ex: greetings3/greetings.go we want to write test case then create file called greetings_test.go
- To run Testing use command: `go test`
- basics-go/greetings3 \$ go test -v (Verbose)

# Go Commands

- go build command to compile the code into an executable
  - \$ go build
  - To run the executable in mac or linux: ./executable
  - basics-go/hello \$ go run .
  - It would generate executable as hello
  - basics-go/hello \$ ./hello
- Go List Commands
  - basics-go/hello \$ go list -f '{{.Target}}'
  - meaning that binaries are installed at this location /Users/tsabunkar/go/bin/hello

# Go Lang for REST API using GIN framework

- https://gin-gonic.com/docs/
- Gin is a HTTP web framework written in Go
- \$ mkdir web-service-gin
- \$ cd web-service-gin
- \$ go mod init example/web-service-gin (Creating modules to manage deps`)
- \$ touch main.go
- Added the following code with gin deps
- To downloaded this gin 3rd party deps-
  - basics-go/web-service-gin \$ go get .
- basics-go/web-service-gin\$ go run .
- open new terminal: curl http://localhost:8080/albums
- GET ALL: \$ curl http://localhost:8080/albums
- GET ONE: \$ curl http://localhost:8080/albums/1
- POST:
  - \$ curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
  - To verifiy: \$ curl http://localhost:8080/albums
- DELETE:
  - \$ curl http://localhost:8080/albums/1 \
    --request "DELETE"
- PUT:
  - \$ curl http://localhost:8080/albums/1 \
    --include \
    --header "Content-Type: application/json" \
    --request "PUT" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'

# How to install Go from the traditional installtion, which is present in /usr/local/go/bin/go

- \$ which go
- \$ sudo rm -rf /usr/local/go
- \$ sudo rm /etc/paths.d/go
- \$ which go (or) realpath
- \$ sudo rm -rf ./go (if present under: /tsabunkar/go)

# Installing go using homebrew

- \$ brew install go
- \$ go version (go version go1.23.5 darwin/arm64, Since its for apple silicon)

# Hot Reloading in Golang

- There is no hotloading in Golang, it has live reloading
  - Live reloading is when a background process observes changes in your source and recompiles your entire application automatically. Hot reloading is when you can recompile only part of your application and inject it into your already running process
- https://github.com/air-verse/air
- \$ cd /Users/tsabunkar/Developer/Workspace/vsc/basics-go/hello
- \$ go mod init main (If not initialized the module)
- \$ go install github.com/air-verse/air@latest
- \$ air -h
- If zsh command not found air
  - Your Go Global Modules are presents at- /Users/tsabunkar/go/bin
  - \$ cd \Users\tsabunkar\go\bin
  - \$ ls (you will find air here)
  - \$ code ~/.zshrc
  - # zsh to remeber all the global modules/packages of go-lang
    export PATH="$PATH:/Users/tsabunkar/go/bin"
  - Save above file, Terminate all the termnials and re=open
  - \$ air -v (anywhere)
- \$ cd /Users/tsabunkar/Developer/Workspace/vsc/basics-go/hello
- \$ go init (This will create .air.toml file which is config file for air-configuration for this directory/module)
- \$ air (Instead of running go run .)
