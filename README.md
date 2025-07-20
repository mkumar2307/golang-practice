# 🧠 Go Practice Repository

Welcome to my Go (Golang) learning and practice repository! 🚀  
This repo tracks my journey into learning the Go programming language through examples, exercises, and small projects.

---

## 📆 Learning Plan

### Go Basics
- Variables, constants, data types
- Slices, maps, loops, conditionals
- Functions and basic CLI programs

### Structs, Interfaces, Errors
- Structs and receivers
- Interfaces and polymorphism
- Error handling best practices

### Concurrency
- Goroutines
- Channels
- Mutex and `sync` package

### HTTP & REST APIs
- Basic HTTP server with `net/http`
- JSON serialization
- RESTful design using routers (`chi`, `mux`)

### Testing & Tooling
- Unit testing with `testing`
- Benchmarks and coverage
- Using linters and Go modules

### Projects
- CLI tool
- REST API
- Real-world concurrency project

---

## 📚 Resources I'm Using

- [Go Tour](https://go.dev/tour/list)
- [Getting Started with Go](https://www.coursera.org/learn/golang-getting-started#modules)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Awesome Go](https://github.com/avelino/awesome-go)

---
## ▶️ Running & Building Go Code

To run a Go program, place the code in a file like `hello-world.go` and use:

```bash
$ go run hello-world.go
hello world
```

To build our programs into binaries. We can do this using go build.    

```bash
$ go build hello-world.go
$ ls
hello-world    hello-world.go
```

We can then execute the built binary directly.

```bash
$ ./hello-world
hello world
```

---

## 📌 Goals

- Become confident writing idiomatic Go code
- Learn Go's standard library and tools
- Understand concurrency patterns and memory safety
- Build real-world microservices and CLI tools

---

Feel free to fork or suggest ideas. Let's Go! 💪🐹