package main

import (
	"fmt"
)

func InterfaceExample() {
	typeAssertion()
	emptyInterface()
}

//interface definition
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

//MyString implements VowelsFinder
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func typeAssertion() {
	name := MyString("Sam Anderson")
	fmt.Println(name)
	var v VowelsFinder
	v = name // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c", v.FindVowels())
	var s MyString
	var ok bool
	s, ok = v.(MyString) // use type assertion to fetch the underlying concrete value
	if ok {
		fmt.Println(s)
	}
}

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func emptyInterface() {
	s := "Hello World"
	describe(s)
	i := 55
	describe(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe(strt)
}
