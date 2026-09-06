package main

import (
	"errors"
	"fmt"
)

func Factorial() func(int) int {
	cache := make(map[int]int)
	var calculate func(int) int
	calculate = func(num int) int {
		if num == 1 || num == 0 {
			return 1
		}
		if num < 0 {
			err := errors.New("Factorial is not defined for negative numbers")
			fmt.Println(err)
			return 0
		}
		if value, found := cache[num]; found {
			fmt.Println("From Cache")
			return value
		}
		cache[num] = num * calculate(num-1)
		return cache[num]
	}

	return calculate

}
func Palindrome() func(string) bool {
	return func(s string) bool {
		for i := 0; i < len(s)/2; i++ {
			if s[i] != s[len(s)-1-i] {
				return false
			}
		}
		return true
	}
}
func Anagrams() func(string, string) bool {
	return func(s1, s2 string) bool {
		if len(s1) != len(s2) {
			return false
		}
		counts := make(map[rune]int)
		for _, char := range s1 {
			counts[char]++
		}
		for _, char := range s2 {
			counts[char]--
		}
		for _, count := range counts {
			if count != 0 {
				return false
			}
		}
		return true
	}
}
func Average(nums ...float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}

func Calculator(a float64, b float64, op string) float64 {
	var result float64
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			err := errors.New("Division by zero is not allowed")
			fmt.Println(err)
			return 0
		}
		result = a / b
	default:
		err := errors.New("Invalid operator")
		fmt.Println(err)
		return 0
	}
	return result
}

func main() {

	// fact := Factorial()
	// fmt.Println(fact(3))
	// ex := Palindrome()
	// fmt.Println(ex("eye"))
	// an := Anagrams()
	// fmt.Println(an("listen", "silent"))
	// avg := Average(1, 2, 3, 4, 5)
	// fmt.Println(avg)
	avg2 := Average(10, 20, 30)
	fmt.Println(avg2)
	calc := Calculator(10, 5, "+")
	fmt.Println(calc)
	calc2 := Calculator(10, 0, "/")
	fmt.Println(calc2)
	calc3 := Calculator(10, 5, "%")
	fmt.Println(calc3)

}
