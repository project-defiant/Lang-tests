// this is the way to define a cli application that will make a fizzbuz
// it will require a parameter to pass to it otherwise should fail

package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	// define a simple CLI
	var input = flag.Int("n", 0, "Provide a number")
	flag.Parse()

	var mapping = map[int]string{
		3: "Fizz",
		5: "Buzz",
	}

	fmt.Printf("You have provided number: %v\n", *input)
	for i := range *input {
		msg := fizzBuzz(i, &mapping)
		fmt.Printf("Number %v: %v\n", i, msg)
	}
}

func fizzBuzz(number int, mapping *map[int]string) (msg string) {

	msg = ""
	for val, str := range *mapping {
		if number%val == 0 {
			msg += str
		}
	}
	if len(msg) == 0 {
		msg = strconv.Itoa(number)
	}

	return msg
}
