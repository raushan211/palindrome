package main

import "fmt"

func main() {
	var x int
	fmt.Println("enter no.")
	fmt.Scanln(&x)
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {

	r := 0
	sum := 0
	temp := x
	result := false
	for x > 0 {
		r = x % 10
		sum = (sum * 10) + r
		x = x / 10
	}
	if temp == sum {
		result = true
	} else {
		result = false
	}
	return result
}
