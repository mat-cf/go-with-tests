package main

import "fmt"

const (
	spanish = "Spanish"
	portuguese = "Portuguese"

	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	portuguesePrefix = "Olá, "
)

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case portuguese:
		prefix = portuguesePrefix
	default:
		prefix = englishPrefix
	}
	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("", ""))
}
