package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Ensure we have two numbers in input.
	if len(os.Args[1:]) != 2 {
		fmt.Println("Use: gcf <number1> <number2>")
		return
	}

	// Parse the numbers and ensure they are integers.

	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%s is not a number, please enter a correct number.\n", os.Args[1])
		return
	}

	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("%s is not a number, please enter a correct number.\n", os.Args[2])
		return
	}

	GCF := greatestCommonFactor(a, b)

	// If the function returns -1 it means the GCD is undefined.
	if GCF == -1 {
		fmt.Printf("Greatest common factor of %v and %v is undefined.\n", a, b)
		return
	}

	fmt.Printf("Greatest common factor of %v and %v is %v.\n", a, b, GCF)

}

func greatestCommonFactor(a, b int) int {
	var gcf int

	// GCF(0,0) is undefined. The function returns -1 if value is undefined.
	if a == 0 && b == 0 {
		return -1
	}

	// We set our counter at the smallest value between a and b.
	gcf = a
	if a > b {
		gcf = b
	}

	// Then we iterate in descending order and verify the condition.
	for gcf != 0 {
		if a%gcf == 0 && b%gcf == 0 {
			return gcf
		}
		gcf--
	}

	// if GCF reaches 0 means one of the two values is zero, we return the biggest number using the rule GCD(a,0) = GCD(0,a) = a.
	if a > b {
		return a
	}
	return b
}
