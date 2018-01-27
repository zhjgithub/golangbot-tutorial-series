package main

import (
	"fmt"
)

func SwitchExample() {
	multipleExpression()
	fallthroughExample()
}

func multipleExpression() {
	letter := "i"
	switch letter {
	case "a", "e", "i", "o", "u": //multiple expressions in case
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}
}

func number() int {
	num := 15 * 5
	return num
}

func fallthroughExample() {
	switch num := number(); { //num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
		// fallthrough should be the last statement in a case
	case num < 60:
		fmt.Printf("%d is greater than 60", num)
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}
}
