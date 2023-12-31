package main

import (
	"fmt"
)

func main() {
	var age int
	var name string
	var isStudent bool

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(isStudent)

	age = 35
	name = "Bob"
	isStudent = false

	fmt.Printf("Age: %v, Name: %v, Student: %v\n", age, name, isStudent)

	var firstname, lastname string
	firstname = "Rob"
	lastname = "Pike"

	fmt.Printf("Firstname : %v, Lastname: %v\n", firstname, lastname)

	var (
		a       int
		b       bool
		c, d, e float64
	)
	fmt.Println(a, b, c, d, e)

	var dogName string = "Hiro"
	fmt.Printf("Dog name : %v\n", dogName)

	var catName, catSurename string = "Felix", "Minou"
	fmt.Printf("Cat name : %v, cat surname : %v\n", catName, catSurename)

	var (
		f    int
		g, h bool = true, false
	)

	fmt.Println(f, g, h)

	var city = "Paris"

	fmt.Println(city)
	fmt.Printf("City : %v, type: %T\n", city, city)

	var cityA, cityB, ok = "A", "B", true
	fmt.Println(cityA, cityB, ok)

	var infA, infB = 8, 8.0
	fmt.Printf("Type A : %T, type B : %T\n", infA, infB)

	var minAge int8 = 7
	var minAge2 = minAge
	fmt.Printf("Type inferé : %T\n", minAge2)

	short := true
	fmt.Println(short)

	shortA, shortB, shortC := 1, 2, "ok"
	fmt.Println(shortA, shortB, shortC)
}
