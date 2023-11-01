package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func luhnAlgorithm(arr []int) bool {
	sum := 0
	parity := len(arr) % 2
	for i := 0; i < (len(arr) - 1); i++ {
		if i%2 != parity {
			sum += arr[i]
		} else if int(arr[i]) > 4 {
			sum += 2*arr[i] - 9
		} else {
			sum += 2 * arr[i]
		}
	}
	return arr[len(arr)-1] == (10 - (sum % 10))
}
func checkArgsLength(args []string) (bool, error) {
	if len(args) < 1 {
		return false, errors.New("no argument provided")
	} else if len(args) > 1 {
		return false, errors.New("too many arguments, program wants only one argument")
	} else {
		return true, nil
	}
}

func main() {
	args := os.Args[1:]
	if _, err := checkArgsLength(args); err != nil {
		fmt.Println(err)
		return
	} else {
		if _, err := strconv.Atoi(args[0]); err != nil {
			fmt.Println("argument should only contain numbers")
			return
		} else {
			numberArray := make([]int, len(args[0]))
			for i, number := range args[0] {
				numberArray[i], _ = strconv.Atoi(string(number))
			}
			if luhnAlgorithm(numberArray) {
				fmt.Println("Valid number")
			} else {
				fmt.Println("Non-valid number")
			}
		}
	}
}
