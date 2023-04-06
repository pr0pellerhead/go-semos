package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Calculator struct {
	A         int    `json:"a" valid:"min=10"`
	B         int    `json:"b" valid:"max=10"`
	Operation string `json:"op" valid:"mustbe=+-%/*"`
}

// func Validate(c interface{}) error {
func Validate(c any) error {
	rv := reflect.ValueOf(c)
	for i := 0; i < rv.NumField(); i++ {
		t := rv.Type().Field(i)
		v := rv.Field(i).Interface()
		tn := t.Type.Name()
		fn := t.Name
		tag := t.Tag.Get("valid")
		fmt.Println(t, v, tn, fn, tag)

		s := strings.Split(tag, "=")
		fmt.Println(s)

		switch o := v.(type) {
		case int:
			if err := validateInt(o, s[0], s[1]); err != nil {
				return errors.New("invalid value for field:" + fn)
			}
		case string:
			if err := validateStr(o, s[0], s[1]); err != nil {
				return errors.New("invalid value for field:" + fn)
			}
		}
	}
	return nil
}

func validateInt(val int, valType string, valNum string) error {
	vn, err := strconv.Atoi(valNum)
	if err != nil {
		return err
	}
	switch valType {
	case "min":
		if val < vn {
			return errors.New("failed on min value")
		}
	case "max":
		if val > vn {
			return errors.New("failed on max value")
		}
	case "eq":
		if val == vn {
			return errors.New("failed on eq value")
		}
	}
	return nil
}

func validateStr(val string, valType string, valStr string) error {
	switch valType {
	case "mustbe":
		if !strings.Contains(valStr, val) {
			return errors.New("failed on mustbe value")
		}
	}
	return nil
}

func main() {
	// calc := Calculator{
	// 	A:         10,
	// 	B:         6,
	// 	Operation: "*",
	// }

	c := &Calculator{}

	input := `{
		"a": 10,
		"b": 6,
		"op": "*"
	}`

	if err := json.Unmarshal([]byte(input), c); err != nil {
		fmt.Println(err.Error())
	}

	if err := Validate(*c); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Data is VALID!")
	}
}
