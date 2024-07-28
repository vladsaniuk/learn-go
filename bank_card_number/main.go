package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func validate(data string) ([]string, error) {
	if len(data) < 19 {
		return nil, fmt.Errorf("please, provide valid card number in the format \"8787 9999 2345 7272\", pay attention to spaces")
	}

	card := strings.Split(data, " ")
	for _, section := range card {
		_, err := strconv.Atoi(section)
		if err != nil {
			return nil, fmt.Errorf("please, provide valid card number in the format \"8787 9999 2345 7272\", pay attention that it should be integers")
		}
	}

	return card, nil
}

func mask(card []string) []string {
	card[1] = "****"
	card[2] = "****"

	return card
}

func show(maskedCard []string) {
	fmt.Println(strings.Join(maskedCard, " "))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// when you hit enter - the data is read
	scanner.Scan()
	// this will read data from the scanner
	data := scanner.Text()
	card, err := validate(data)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	maskedCard := mask(card)
	show(maskedCard)
}
