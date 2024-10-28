package main

/*
Prime number validator implementation. It allows for computing numbers that are at most MaxUnit64 size
*/

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"

	"github.com/project-defiant/primes/lib"
)

func main() {

	var input string
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Could not find a number in the input parameters, provide one now:")
		fmt.Scanf("%s", &input)
	} else {
		input = args[1]
	}
	number, error := strconv.ParseUint(input, 10, 64)

	if error != nil {
		// ensure that the value does not surpass the math.MaxUint64
		integerSlice := []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
		for pos, char := range input {
			if !slices.Contains(integerSlice, char) {
				fmt.Printf("Provided incorrect character at pos %v\n", pos)
				os.Exit(1)
			}
		}
		fmt.Printf("Provided number is greater then the %v\n", strconv.FormatUint(math.MaxUint64, 10))
		os.Exit(2)
	}

	if !lib.IsPrimeNumber(number) {
		fmt.Printf("Input %s is not a prime number\n", input)
	} else {
		fmt.Printf("Input %s is a prime number\n", input)
	}

}
