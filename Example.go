package main

import (
	"fmt"
)

func test() {
	// Example of printing to the console
	fmt.Println("Hello, World!")

	// Basic math operations
	num1 := 10
	num2 := 5

	// Addition
	sum := add(num1, num2)
	fmt.Printf("Sum: %d\n", sum)

	// Subtraction
	difference := subtract(num1, num2)
	fmt.Printf("Difference: %d\n", difference)

	// Multiplication
	product := multiply(num1, num2)
	fmt.Printf("Product: %d\n", product)

	// Division with error handling
	quotient, err := divide(num1, num2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Quotient: %f\n", quotient)
	}

	// Example of using a loop
	for i := 0; i < 5; i++ {
		fmt.Printf("Loop iteration %d\n", i)
	}
}

// Function to add two integers
func add(a, b int) int {
	return a + b
}

// Function to subtract two integers
func subtract(a, b int) int {
	return a - b
}

// Function to multiply two integers
func multiply(a, b int) int {
	return a * b
}

// Function to divide two integers, returns an error for division by zero
func divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return float64(a) / float64(b), nil
}
