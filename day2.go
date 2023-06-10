package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.

	var i2 int
	var d2 float64
	var s2 string

	// Read and save an integer, double, and String to your variables.
	fmt.Println("enter an integer")
	scanner.Scan()
	i2, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	fmt.Println("enter a float")
	scanner.Scan()
	d2, err = strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		return
	}

	fmt.Println("enter a string")
	scanner.Scan()
	s2 = scanner.Text()

	// Print the sum of both integer variables on a new line.
	fmt.Println(int(i) + i2)
	// Print the sum of the double variables on a new line.
	sum := d + d2
	fmt.Printf("%.2f\n", sum)
	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Printf("%s%s\n", s, s2)
}
