package main

import (
	"fmt"
	"sort"
)

type Stringer = interface {
	String() string
}

type Ordered interface {
	~int | ~float64 | ~string
}

type Integer int

func (i Integer) String() string {
	return fmt.Sprintf("%d", i)
}

type String string

func (s String) String() string {
	return string(s)
}

type Student struct {
	Name string
	ID   int
	Age  float64
}

func (s Student) String() string {
	return fmt.Sprintf("%s %d %0.2f", s.Name, s.ID, s.Age)
}

func addStudent[T any](students []T, student T) []T {
	return append(students, student)
}

type OrderedSlice[T Ordered] []T

func (s OrderedSlice[T]) Len() int {
	return len(s)
}

func (s OrderedSlice[T]) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s OrderedSlice[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type SortType[T any] struct {
	slice   []T
	compare func(T, T) bool
}

func (s SortType[T]) Len() int {
	return len(s.slice)
}

func (s SortType[T]) Less(i, j int) bool {
	return s.compare(s.slice[i], s.slice[j])
}

func (s SortType[T]) Swap(i, j int) {
	s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
}

func PerformSort[T any](slice []T, compare func(T, T) bool) {
	sort.Sort(SortType[T]{slice, compare})
}

func main() {
	students := []string{} // Empty slice
	result := addStudent[string](students, "Michael")
	result = addStudent[string](result, "Jennifer")
	result = addStudent[string](result, "Elaine")
	sort.Sort(OrderedSlice[string](result))
	fmt.Println(result)

	students1 := []int{}
	result1 := addStudent[int](students1, 45)
	result1 = addStudent[int](result1, 64)
	result1 = addStudent[int](result1, 78)
	fmt.Println(result1)

	students2 := []Student{}
	result2 := addStudent[Student](students2, Student{"John", 213, 17.5})
	result2 = addStudent[Student](result2, Student{"James", 111, 18.75})
	result2 = addStudent[Student](result2, Student{"Marsha", 110, 16.25})
	PerformSort[Student](result2, func(s1, s2 Student) bool {
		return s1.Age < s2.Age
	})
	fmt.Println(result2)
}
