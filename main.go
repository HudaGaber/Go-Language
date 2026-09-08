package main

import (
	"fmt"
	"task1/functions"
)

func main() {
	fact := functions.Factorial()
	fmt.Println(fact(3))
	ex := functions.Palindrome()
	fmt.Println(ex("eye"))
	an := functions.Anagrams()
	fmt.Println(an("listen", "silent"))
	avg := functions.Average(1, 2, 3, 4, 5)
	fmt.Println(avg)
	avg2 := functions.Average(10, 20, 30)
	fmt.Println(avg2)
	calc := functions.Calculator(10, 5, "+")
	fmt.Println(calc)
	calc2 := functions.Calculator(10, 0, "/")
	fmt.Println(calc2)
	calc3 := functions.Calculator(10, 5, "%")
	fmt.Println(calc3)

}
