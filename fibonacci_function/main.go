package main

import "fmt"

func fibonacci() func() {
	prevPrevNumber := 0
	prevNumber := 1
	return func() {
		newNumber := prevPrevNumber + prevNumber
		fmt.Println(newNumber)
		prevPrevNumber = prevNumber
		prevNumber = newNumber
	}
}

func main() {
	function := fibonacci()
	for i := 0; i < 10; i++ {
		function()
	}
}
