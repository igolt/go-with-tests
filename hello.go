package main

import "fmt"

const (
	spanish            = "Spanish"
	spanishHelloPrefix = "Hola,"
	englishHelloPrefix = "Hello,"
)

func Hello(name string, language string) string {
	prefix := englishHelloPrefix
	if name == "" {
		name = "World"
	}
	if language == spanish {
		prefix = spanishHelloPrefix
	}
	return fmt.Sprintf("%s %s!", prefix, name)
}

func main() {
	fmt.Println(Hello("world", ""))
}
