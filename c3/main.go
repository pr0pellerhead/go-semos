package main

import (
	"encoding/json"
	"fmt"
	"semos/c3/pkg/blogpost"
	"time"

	"github.com/k0kubun/pp"
)

func main() {

	// conf := student.StudentConfig{
	// 	FirstName: "Ivana",
	// 	LastName:  "Ivanovic",
	// }

	// stu, err := student.NewStudent(conf)
	// if err != nil {
	// 	fmt.Println("Error!", err.Error())
	// }
	// fmt.Println(stu)

	//  **************************************

	// stu, err := student.NewStudent2(
	// 	student.WithFirstName("Goran"),
	// 	student.WithLastName("Goranovic"),
	// )
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(stu)

	//  **************************************

	// stu := student.NewStudent3("Bob", "Smith").SetEmail("bob@smith.com").SetGPA(9.3).SetYear(4)

	// stu2 := student.NewStudent3("John", "Doe").SetGPA(6.1).SetYear(2)

	// fmt.Println(stu)
	// fmt.Println(stu2)

	//  **************************************

	input := `{
		"title": "My sample title",
		"description": "The posts description",
		"post_content": "Some content blah blah blah...",
		"author": {
			"full_name": "John Doe",
			"email": "john@doe.com"
		},
		"publish_date": "2023-04-04T18:21:55.543Z",
		"tags": ["tag1", "tag2", "tag3"],
		"comments": [
			{
				"author": {
					"full_name": "Jane Doe",
					"email": "jane@doe.com"
				},
				"comment_body": "Comment blah blah blah"
			},
			{
				"author": {
					"full_name": "John Smith",
					"email": "john@smith.com"
				},
				"comment_body": "Comment blah blah blah balah..."
			}
		]
	}`

	bp := &blogpost.Blogpost{}

	err := json.Unmarshal([]byte(input), bp) // json_decode
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("TITLE:", bp.Title)
	fmt.Println("DESC:", bp.Description)
	fmt.Println("CONTENT:", bp.PostContent)
	fmt.Println("AUTHOR NAME:", bp.Author.FullName)
	fmt.Println("AUTHOR EMAIL:", bp.Author.Email)
	fmt.Println("TAGS:", bp.Tags)

	for _, c := range bp.Comments {
		fmt.Println("Person:", c.Author.FullName)
		fmt.Println("Comment:", c.CommentBody)
		fmt.Println("--------------------------------")
	}

	fmt.Println(bp.PublishDate.Format("Monday, January 2, 2006"))

	// fmt.Println(bp)

	pp.Println(bp)

	nbp := &blogpost.Blogpost{
		Title:       "A new blogpost to JSON",
		Description: "desc...",
		PostContent: "Content content content",
		Author: blogpost.Author{
			FullName: "Steven Johnson",
			Email:    "stephen@johnson.com",
		},
		Tags:        []string{"go", "course", "golang", "dates-are-weird"},
		PublishDate: time.Now(),
	}

	out, err := json.Marshal(nbp)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(out))
}
