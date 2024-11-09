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

## Import a Module

- To import 3rd party module
  - \$ go mod tidy (When imported rsc.io/quote)
  - go mod tidy, it located and downloaded the rsc.io/quote module that contains the package you imported. By default, it downloaded the latest version
- \$ go run import-fun.go
