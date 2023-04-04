package student

import "errors"

type Student struct {
	Name      string
	LastName  string
	GPA       float64
	Graduated bool
	Location
}

type Location struct {
	Country string
	City    string
}

func NewStudent(name string, lastName string) (*Student, error) {
	if name == "" {
		return nil, errors.New("no student name assigned")
	}
	if lastName == "" {
		return nil, errors.New("no student last name assigned")
	}
	s := &Student{
		Name:     name,
		LastName: lastName,
		GPA:      5,
		Location: Location{
			City:    "Podgorica",
			Country: "Montenegro",
		},
	}
	return s, nil
}

func (s *Student) AddGPA(gpa float64) {
	s.GPA += gpa
	s.calculateGPA()
}

func (s *Student) calculateGPA() {
	s.GPA -= s.GPA / 10
}

// try {
// 	comman1()
// 	comman2()
// 	comman2()
// 	comman2()
// 	comman2()
// 	comman2()
// 	comman2()
// 	comman2()
// 	comman2()
// 	comman2()
// } catch($err) {
// 	$err->msg()
// }
