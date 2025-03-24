package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// We wanted to test a function that greets someone. First we had a function
// that prints to the terminal, but we couldn't test because we had no control
// to where it printed.
//
//	func Greet(name string) {
//		fmt.Printf("Hello, %s", name)
//	}
//
// We then changed the implementation so we could "inject" the place where the
// function prints. In this way we can have control over the writer and can
// test the function result.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// With the ability to pass in the writer we can also write to a response to a
// HTTP response.
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}
