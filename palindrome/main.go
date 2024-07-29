package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func isValidInput(s string) error {
	validInput, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("failed to convert string %s to int", s)
	}

	if validInput < 10 {
		return fmt.Errorf("invalid input, please, provide number equal or more than 10")
	}

	return nil
}

func convertToIntArray(s string) ([]int, error) {
	intArray := make([]int, 0)

	for _, v := range strings.Split(s, "") {
		intConverted, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("failed to convert string %s to int", v)
		}
		intArray = append(intArray, intConverted)
	}

	return intArray, nil
}

func isPalindrome(a []int, palindromes *[]string) {
	if len(a)%2 != 0 {
		fmt.Printf("Length %v array is odd, odd numbers cannot be palindromes\n", a)
		return
	} else if len(a) == 0 {
		fmt.Printf("Don't compare slices with no elements\n")
		return
	}

	arrayMiddle := len(a) / 2
	firstHalf := a[0:arrayMiddle]
	secondHalf := a[arrayMiddle:]

	revertedSecondHalf := make([]int, 0)
	for i := 0; i < len(secondHalf); i++ {
		revertedSecondHalf = append(revertedSecondHalf, secondHalf[len(secondHalf)-1-i])
	}

	if reflect.DeepEqual(firstHalf, revertedSecondHalf) {
		fmt.Printf("%d is palindrome\n", a)
		newPalindrome := fmt.Sprint(a)
		if !isPalindromeAlreadyFound(*palindromes, newPalindrome) {
			*palindromes = append(*palindromes, newPalindrome)
		}
	} else {
		fmt.Printf("%d isn't palindrome\n", a)
	}
}

func isPalindromeAlreadyFound(palindromes []string, newPalindrome string) bool {
	for _, palindromeInSlice := range palindromes {
		if palindromeInSlice == newPalindrome {
			return true
		}
	}

	return false
}

func findPalindromes(a []int) []string {
	palindromes := make([]string, 0)

	for n := len(a); n >= 0; n-- {
		fmt.Printf("New run, n is %d\n", n)
		for i := range a {
			if i > n {
				continue
			}
			array := a[i:n]
			isPalindrome(array, &palindromes)
		}
	}

	return palindromes
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// when you hit enter - the data is read
	scanner.Scan()
	// this will read data from the scanner
	data := scanner.Text()

	err := isValidInput(data)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	intArray, err := convertToIntArray(data)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	palindromes := findPalindromes(intArray)

	fmt.Printf("In number %s unique palindromes are:\n", data)
	for _, palindrome := range palindromes {
		fmt.Println(palindrome)
	}
}
