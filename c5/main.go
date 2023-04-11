package main

import "fmt"

func calculate(a, b int, op string) int {
	switch op {
	case "add":
		return a + b
	case "substract":
		return a - b
	case "multiply":
		return a * b
	}
	return 0
}

func vocals(text string) bool {
	vcs := "aouie"
	count := 0
	for _, l := range text {
		for _, v := range vcs {
			if l == v {
				count++
			}
		}
	}
	return count%2 == 0
}

func main() {
	res := calculate(2, 4, "multiply")
	fmt.Println(res)
}
