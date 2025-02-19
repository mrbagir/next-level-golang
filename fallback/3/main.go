package main

import (
	"cmp"
	"fmt"
	"slices"
)

type employee struct {
	name string
	age  int
}

func main() {
	employees := []employee{
		{name: "John", age: 25},
		{name: "Jane", age: 30},
		{name: "Doe", age: 35},
		{name: "Alex", age: 25},
	}

	fmt.Println("Employees before sorting:")
	fmt.Println(employees)

	fmt.Println("Employees after sorting:")
	fmt.Println(sortEmployees(employees))
}

func sortEmployees(employees []employee) []employee {
	sortedEmployees := employees

	slices.SortFunc(
		sortedEmployees,
		func(a, b employee) int {
			return cmp.Or(
				cmp.Compare(a.age, b.age),
				cmp.Compare(a.name, b.name),
			)
		},
	)

	return sortedEmployees
}
