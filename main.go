package main

import "fmt"

func addStudent(students []string, student string) []string {
	return append(students, student)
}

func main() {
	students := []string{} // Empty slice
	result := addStudent(students, "Michael")
	result = addStudent(result, "Jennifer")
	result = addStudent(result, "Elaine")
	fmt.Println(result)
}
