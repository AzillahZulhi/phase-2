package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unicode/utf8"
)

type SuperHero struct {
	Name       string `required:"true" minLen:"3" maxLen:"50"`
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

		if minLenStr := tag.Get("minLen"); minLenStr != "" {
			minLen, err := strconv.Atoi(minLenStr)
			if err != nil {
				return fmt.Errorf("invalid minLen tag for %s: %v", field.Name, err)
			}
			if strValue, ok := value.(string); ok && utf8.RuneCountInString(strValue) < minLen {
				return fmt.Errorf("%s length must be greater than or equal to %d", field.Name, minLen)
			}
		}

		if maxLenStr := tag.Get("maxLen"); maxLenStr != "" {
			maxLen, err := strconv.Atoi(maxLenStr)
			if err != nil {
				return fmt.Errorf("invalid maxLen tag for %s: %v", field.Name, err)
			}
			if strValue, ok := value.(string); ok && utf8.RuneCountInString(strValue) > maxLen {
				return fmt.Errorf("%s length must be less than or equal to %d", field.Name, maxLen)
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
