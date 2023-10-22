package main

import (
	"fmt"

	"github.com/RomanBoegli/gostark"
)

func main() {
	// Get a greeting message and print it.
	message := gostark.Average(3, 5)
	fmt.Println(message)
	gostark.SayHello()
	gostark.SayWisper()
}
