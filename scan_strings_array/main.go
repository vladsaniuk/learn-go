package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(data string) []string {
	return strings.Split(data, ",")
}

func convert(strArray []string) []int {
	intArray := make([]int, 0)
	for n := 0; n < len(strArray); n++ {
		i, err := strconv.Atoi(strArray[n])
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		intArray = append(intArray, i)
	}

	return intArray
}

func countEvenPositives(intArray []int) {
	var counter int

	for _, n := range intArray {
		if n > 0 && n%2 == 0 {
			counter++
		}
	}

	fmt.Printf("Total number of even positive numbers is %d\n", counter)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// when you hit enter - the data is read
	scanner.Scan()
	// this will read data from the scanner
	data := scanner.Text()
	strArray := parse(data)
	intArray := convert(strArray)
	countEvenPositives(intArray)
}
