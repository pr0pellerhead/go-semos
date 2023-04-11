package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/k0kubun/pp"
)

type Output struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Input struct {
	A  int    `json:"a"`
	B  int    `json:"b"`
	Op string `json:"op"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello World...")
		pp.Println(r)
		fmt.Fprintln(w, "<h1>Hello world... Again :|</h1>")
	})

	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		out := Output{
			FirstName: "John",
			LastName:  "Smith",
		}
		jsonOutput, err := json.Marshal(out)
		if err != nil {
			log.Println("JSON parse err in /json func, ", err.Error())
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(jsonOutput)
		// fmt.Println("Hello JSON...")
	})

	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			in, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println("JSON parse err in /json func, ", err.Error())
			}
			input := &Input{}
			if err := json.Unmarshal(in, input); err != nil {
				log.Println("JSON parse err in /json func, ", err.Error())
			}
			pp.Println(input)
		}
	})

	fmt.Println("Server is starting")
	http.ListenAndServe(":10000", nil)
}
