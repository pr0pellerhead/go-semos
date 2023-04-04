package student

import "errors"

type Student struct {
	FirstName string
	LastName  string
	Email     string
	GPA       float64
	Year      int
}

// Constructing an object with a configuration

type StudentConfig struct {
	FirstName string
	LastName  string
}

func NewStudent(config StudentConfig) (*Student, error) {
	s := &Student{
		FirstName: config.FirstName,
		LastName:  config.LastName,
	}
	return s, nil
}

// Contructing an object with options

type StudentOption func(s *Student) error

func NewStudent2(options ...StudentOption) (*Student, error) {
	s := &Student{}
	for _, o := range options {
		if e := o(s); e != nil {
			return nil, e
		}
	}
	return s, nil
}

func WithFirstName(name string) StudentOption {
	return func(s *Student) error {
		if len(name) == 0 {
			return errors.New("name cannot be empty")
		}
		s.FirstName = name
		return nil
	}
}

func WithLastName(name string) StudentOption {
	return func(s *Student) error {
		if len(name) == 0 {
			return errors.New("last name cannot be empty")
		}
		s.LastName = name
		return nil
	}
}

// Method chaining

// collection("students")
// 	.Find({student: "Marija"})
// 	.Limit(10)

func NewStudent3(fn string, ln string) *Student {
	s := &Student{
		FirstName: fn,
		LastName:  ln,
	}
	return s
}

func (s *Student) SetGPA(gpa float64) *Student {
	s.GPA = gpa
	return s
}

func (s *Student) SetEmail(eml string) *Student {
	s.Email = eml
	return s
}

func (s *Student) SetYear(y int) *Student {
	s.Year = y
	return s
}
