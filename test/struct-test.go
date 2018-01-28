package main

import "fmt"

// named structure
type Employee struct {
	firstName, lastName string
	age, salary         int
}

// anonymous structures
var employee struct {
	firstName, lastName string
	age, salary         int
}

func StructExample() {
	createNamedStructure()
	createAnonymousStructure()
	anonymousFields()
	promotedfields()
}

func createNamedStructure() {
	//creating structure using field names
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}

	//creating structure without using field names
	emp2 := Employee{"Thomas", "Paul", 29, 800}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)
}

func createAnonymousStructure() {
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}

	fmt.Println("Employee 3", emp3)
	emp4 := struct {
		firstName, lastName string
		age, salary         int
	}{"Andreah", "Nikola", 31, 5000}

	fmt.Println("Employee 4", emp4)

	employee.age, employee.firstName, employee.lastName, employee.salary = 1, "hello", "world", 300
	fmt.Println(employee)
}

type Person struct {
	string
	int
}

func anonymousFields() {
	p := Person{"Naveen", 50}
	fmt.Println(p)

	// by default the name of a anonymous field is the name of its type
	var p1 Person
	p1.string = "naveen"
	p1.int = 50
	fmt.Println(p1)
}

type Address struct {
	city, state string
}
type PersonInfo struct {
	name string
	age  int
	Address
}

func promotedfields() {
	var p PersonInfo
	p.name = "Naveen"
	p.age = 50
	p.Address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.city)   //city is promoted field
	fmt.Println("State:", p.state) //state is promoted field
}
