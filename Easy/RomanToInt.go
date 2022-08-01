package main

import (
	"bufio"
	"fmt"
	"os"
)

var numerals = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

func FindNextIndex(index int, s string) int {
	var i int
	if index != 0 && numerals[string(s[index])] >= numerals[string(s[index-1])] {
		for i = index; i >= 0; i-- {
			if i > 0 && s[i] < s[i-1] {
				break
			}
		}
	} else {
		i = index
	}
	return (i)
}

func SumRoman(index int, s string) int {
	sum := 0
	if index != 0 && numerals[string(s[index])] >= numerals[string(s[index-1])] {
		for i := index; i >= 0; i-- {
			if i > 0 && s[i] < s[i-1] {
				break
			} else {
				sum += numerals[string(s[i])]
			}
		}
	} else {
		sum = numerals[string(s[index])]
	}
	return (sum)
}

func romanToInt(s string) int {
	sum := 0
	digit := len(s) - 1
	for digit >= 0 {
		r := numerals[string(s[digit])]
		if digit != 0 && numerals[string(s[digit])] > numerals[string(s[digit-1])] {
			sum += r - SumRoman(digit-1, s)
			digit = FindNextIndex(digit-1, s)
		} else {
			sum += r
		}
		digit--
	}
	return sum
}

func main() {
	var s string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Roman Number : ")
		scanner.Scan()
		s = scanner.Text()
		if len(s) != 0 {
			fmt.Printf("number equal to %d\n", romanToInt(s))
		} else {
			break
		}
	}
}
