package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	// Demonstrating integer arithmetic using constants
	const maxInt16 int = 32767
	var intNum = maxInt16 + 1
	fmt.Println("Integer addition result:", intNum)

	// Working with float64 type variables
	var floatNum = 12345678.9
	fmt.Println("Float64 value:", floatNum)

	// Type casting and arithmetic operations between float32 and int32
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println("Float32 addition result:", result)

	// Demonstrating integer division and remainder calculations
	var intDividend = 3
	var intDivisor = 2
	fmt.Println("Integer division result:", intDividend/intDivisor)
	fmt.Println("Integer remainder result:", intDividend%intDivisor)

	// String concatenation example
	var myString = "Hello" + " " + "World"
	fmt.Println("Concatenated string:", myString)

	// Counting the number of runes in a Unicode string
	fmt.Println("Rune count:", utf8.RuneCountInString("💯"))

	// Working with rune data type
	var myRune rune = 'a'
	fmt.Println("Rune value:", myRune)

	// Demonstrating boolean data type
	var myBoolean = false
	fmt.Println("Boolean value:", myBoolean)

	// Showing the zero-value of an int data type
	var intNum3 int
	fmt.Println("Zero-value for int:", intNum3)

	// Short variable declaration syntax
	myVar := "text"
	fmt.Println("Short declaration:", myVar)

	// Declaring multiple variables in a single line
	var1, var2 := 1, 2
	fmt.Println("Multiple variable values:", var1, var2)

	// Working with constants
	const myConst = "const value needed here"
	fmt.Println("Constant value:", myConst)

	// Declaring an unused constant (will not throw an error)
	const pi = 3.1415926
	// Demonstrating a function call to print a string value
	var printValue string = "Hello World!"
	printMe(printValue)

	// Demonstrating integer division with a custom function and error handling
	var numerator int = 11
	var denominator int = 2
	var divisionResult, remainder, err = intDivision(numerator, denominator)
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v. ", divisionResult)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v. ", divisionResult, remainder)
	}

	// Using switch-case to classify the division result based on the remainder
	switch remainder {
	case 0:
		fmt.Println("The division was exact.")
	case 1, 2:
		fmt.Println("The division was close.")
	default:
		fmt.Println("The division was not close.")
	}
}

// Function to print a string value
func printMe(printValue string) {
	fmt.Println(printValue)
}

// Custom function for integer division that also handles division by zero
func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	return numerator / denominator, numerator % denominator, err
}
