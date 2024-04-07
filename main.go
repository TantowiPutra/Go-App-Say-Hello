package main

import (
	"fmt"
	sample "github.com/TantowiPutra/Go-Module"
	sample2 "github.com/TantowiPutra/Go-Module/v2"
)

func main() {
	number, text := sample.SayHello("Tantowi")
	fmt.Println("Number:", number)
	fmt.Println("Text:", text)

	text2 := sample2.SayHello("Tantowi", "Putra")
	fmt.Println(text2)
}
