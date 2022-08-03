package main

import (
	"fmt"

	"github.com/yenonn/go-utils/collection"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if collection.Contains[int](numbers, 2) {
		fmt.Println("Found!")
	}

	strings := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

	if collection.Contains[string](strings, "two") {
		fmt.Println("Found!")
	}

	floats := []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 1.0}

	if collection.Contains[float32](floats, 2.0) {
		fmt.Println("Found!")
	}
}