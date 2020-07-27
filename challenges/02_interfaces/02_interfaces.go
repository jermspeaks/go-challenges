package main

import "math/rand"

type Employee interface {
	Language() string
	Age() int
}

type Engineer struct {
	Name string
}

func (e Engineer) Language() string {
	return e.Name + " programs in Go"
}

func (e Engineer) Age() int {
	return 26
}

func (e Engineer) Random() int {
	return rand.Intn(100)
}

func main() {
	// This will throw an error
	var programmers []Employee
	elliot := Engineer{Name: "Elliot"}
	// Engineer does not implement the Employee interface
	// you'll need to implement Age() and Random()
	programmers = append(programmers, elliot)
}
