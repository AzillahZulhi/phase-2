package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type SuperHero struct {
	Name       string `required:"true"`
	Age        int    `required:"true" min:"18" max:"100"`
	SuperPower string `required:"true"`
}

func ValidateStruct(a interface{}) error {
	s := reflect.TypeOf(a)
	v := reflect.ValueOf(a)

	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		value := v.Field(i).Interface()
		tag := field.Tag

		if tag.Get("required") == "true" {
			if value == "" || value == nil {
				return fmt.Errorf("%s is required", field.Name)
			}
		}

		if minStr := tag.Get("min"); minStr != "" {
			min, err := strconv.Atoi(minStr)
			if err != nil {
				return fmt.Errorf("invalid min tag for %s: %v", field.Name, err)
			}
			if intValue, ok := value.(int); ok && intValue < min {
				return fmt.Errorf("%s must be greater than or equal to %d", field.Name, min)
			}
		}

		if maxStr := tag.Get("max"); maxStr != "" {
			max, err := strconv.Atoi(maxStr)
			if err != nil {
				return fmt.Errorf("invalid max tag for %s: %v", field.Name, err)
			}
			if intValue, ok := value.(int); ok && intValue > max {
				return fmt.Errorf("%s must be less than or equal to %d", field.Name, max)
			}
		}
	}
	return nil
}

func main() {
	IronMan := SuperHero{
		Name:       "Tony Stark",
		Age:        50,
		SuperPower: "Rich",
	}
	heroValidate := ValidateStruct(IronMan)
	fmt.Println(heroValidate)
}
