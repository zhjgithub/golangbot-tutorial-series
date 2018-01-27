package main

import "fmt"

func ArrayExample() {
	declareWithDots()
	arrayType()
	slicesShareUnderlyingArray()
	appendSlice()
	copySliceMemoryOptimization()
}

func declareWithDots() {
	a := [...]int{1, 2, 3}
	fmt.Println(a)
}

func arrayType() {
	a := [3]int{5, 78, 8}
	var b [5]int
	// b = a //not possible since [3]int and [5]int are distinct types
	fmt.Println(a, b)
}

func arrayIsValueType() {
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a // a copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)

	// another example
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)
}

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)

}

func slicesShareUnderlyingArray() {
	numa := [3]int{78, 79, 80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	nums3 := nums1[0:2]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)
	nums3[1] = 1
	fmt.Println("array after modification to slice nums3", numa)
}

func appendSlice() {
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)
}

func copySliceMemoryOptimization() {
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy and the original array can be garbage collected
	return countriesCpy
}
