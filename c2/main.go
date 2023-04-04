package main

import (
	"fmt"
	"semos/c2/calc"
	"semos/c2/student"
)

func main() {
	// fmt.Println("class two")

	names := [6]string{
		"Brian",
		"John",
		"Joe",
		"Mary",
		"Sandra",
	}
	fmt.Println(names)

	cities := []string{
		"New York",
		"Los Angeles",
		"San Francisco",
	}
	fmt.Println(cities)

	for k, v := range cities {
		fmt.Printf("cities[%d]: %s\n", k, v)
	}

	cities = append(cities, "Chicago")
	fmt.Println(cities)

	firstTwo := cities[0:2]
	fmt.Println(firstTwo[0:1])

	fmt.Println(cities[2:])
	fmt.Println(cities[:2])
	fmt.Println(cities[:1])

	sentence := `"Lorem ipsum dolor sit amet"`
	fmt.Println(len(sentence))
	fmt.Println(sentence[1 : len(sentence)-1])

	fmt.Println("capacity: ", cap(firstTwo)) // what is the capacity of the slice
	fmt.Println("len: ", len(firstTwo))      // what is the current size

	firstTwo = append(firstTwo, "Skopje")

	fmt.Println("capacity: ", cap(firstTwo)) // what is the capacity of the slice
	fmt.Println("len: ", len(firstTwo))      // what is the current size

	firstTwo = append(firstTwo, "Podgorica")

	fmt.Println("capacity: ", cap(firstTwo)) // what is the capacity of the slice
	fmt.Println("len: ", len(firstTwo))      // what is the current size

	firstTwo = append(firstTwo, "Paris")
	firstTwo = append(firstTwo, "Belgrade")

	firstTwo = append(firstTwo, "Zagreb") // a new (bigger) slice is being created

	fmt.Println("capacity: ", cap(firstTwo)) // what is the capacity of the slice
	fmt.Println("len: ", len(firstTwo))      // what is the current size

	abc := []int{1, 3, 4}
	abc = append(abc, 7)

	fmt.Println(cap(abc))

	var countries []string = make([]string, 3)

	countries = append(countries, "Montenegro")
	fmt.Println(countries)

	emails := map[string]string{
		"Bojan":      "bojan@beyondbasics.co",
		"Ivan":       "ican@ivan.com",
		"Aleksandra": "alex@alex.com",
	}

	for name, email := range emails {
		fmt.Printf("Person %s and the email %s\n", name, email)
	}

	fmt.Println(emails)

	gpa := map[string]float64{
		"George": 7.2,
		"Sarah":  9.7,
		"John":   6.4,
	}

	gpa["Gopher"] = 10.0

	var val float64

	if val, ok := gpa["Gopher"]; ok {
		fmt.Println("The value exists and it's", val)
	} else {
		fmt.Println("The value does not exist")
	}

	fmt.Println(val)

	res := Add(3, 6)
	fmt.Println(res)

	//

	res = calc.Add(5, 10)
	fmt.Println(res)

	res = calc.Multiply(4, 7)
	fmt.Println(res)

	fmt.Println(calc.First, calc.Second)

	// *************************************************
	// structs

	s := student.Student{
		Name: "Ivana",
	}
	s.LastName = "Ivanovic"
	s.GPA = 8
	s.City = "Podgorica"
	s.Country = "Montenegro"

	fmt.Println(s)
	fmt.Println(s.Name)
	fmt.Println(s.City)
	fmt.Println(s.Location)
	fmt.Println(s.Location.Country)

	stu, err := student.NewStudent("Goran", "")
	if err != nil {
		fmt.Println("ERROR!!!!!", err.Error())
	}
	stu.AddGPA(3)
	stu.AddGPA(1.5)

	// stu.calculateGPA() // not going to work... not exported

	fmt.Println(stu)

	age := 30
	ageRef := &age
	fmt.Println(ageRef)
	fmt.Println(*ageRef)
}
