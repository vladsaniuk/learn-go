package main

import (
	"flag"
	"fmt"
)

func printChessboard(height int, width int, symbol string) {
	for row := 0; row <= height; row++ {
		for column := 0; column <= width; column++ {
			if row%2 == 0 {
				if column == width {
					fmt.Printf("%v \n", symbol)
				} else {
					fmt.Printf("%v ", symbol)
				}
			} else {
				if column == width {
					fmt.Printf(" %v\n", symbol)
				} else {
					fmt.Printf(" %v", symbol)
				}
			}
		}
	}
}

func main() {
	var (
		height int
		width  int
		symbol string
	)

	flag.IntVar(&height, "height", 5, "chessboard height, default is 5")
	flag.IntVar(&width, "width", 10, "chessboard width, default is 10")
	flag.StringVar(&symbol, "symbol", "^", "symbol to print chessboard with, default is \"^\"")
	flag.Parse()

	printChessboard(height, width, symbol)
}
