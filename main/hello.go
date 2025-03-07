package main

import (
	"example/hello/structs"
	"fmt"
)

const spanish = "spanish"
const french = "french"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	var prefix string
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix + name
}

func main() {
	// fmt.Println(Hello("Zain", ""))
	rectangle := structs.Rectangle{Width: 10, Height: 10}
	fmt.Println(structs.Area(rectangle))
}
