package main

import (
	"fmt"
)

type httpCode int

func main() {
	fmt.Println(`Hello from go!`)

	// var name string
	// name = "Gopher"

	var name string = "Gopher"
	fmt.Println(name)

	var age int = 14
	fmt.Println(age)

	fmt.Println(name, age, 46213947623, "jcshbdcjhbsdjkch") // \n

	var city, country, continent string = "Podgorica", "Montenegro", "Europe"
	fmt.Printf("City: %s, Country: %s, Continent: %s\n", city, country, continent)

	pi := 3.14
	isItWeekend := false
	its_not_weekend := true

	fmt.Println(pi, isItWeekend, its_not_weekend)
	fmt.Printf("Type: %T Value: %v\n", pi, pi)
	fmt.Printf("Type: %T Value: %v\n", isItWeekend, isItWeekend)

	// type year int

	// var y year = year(2023)

	var ok httpCode = httpCode(200)
	var created httpCode = httpCode(201)
	var notFound httpCode = httpCode(404)

	printCode(ok)
	printCode(created)
	printCode(notFound)

	if pi == 3.14 {
		fmt.Println("YES!")
	} else {
		fmt.Println("NO! :(")
	}

	num := 10

	if num < 10 {
		fmt.Println("less than")
	} else if num > 10 {
		fmt.Println("greater than")
	} else {
		fmt.Println("equals!")
	}

	switch country {
	case "Montenegro":
		fmt.Println("Zdravo!")
		fallthrough
	case "Macedonia":
		fmt.Println("Здраво!")
		fallthrough
	case "US":
		fmt.Println("Hello!")
	default:
		fmt.Println("***")
	}

	sampleString := "Sample string 123"
	fmt.Println(string(sampleString[3]))

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// foreach
	for k, v := range sampleString {
		fmt.Println(k, ":", string(v))
	}

	// for ;; {

	// }

}

func printCode(c httpCode) {
	fmt.Printf("HTTP Code: %d\n", c)
}
