package helloworld

import "fmt"

const (
	french  = "French"
	spanish = "Spanish"

	frenchHelloPrefix  = "Bonjour,"
	spanishHelloPrefix = "Hola,"
	englishHelloPrefix = "Hello,"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s %s!", greetingPrefix(language), name)
}

// prefix is a name return
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
