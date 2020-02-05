package common

import (
	"gopkg.in/go-playground/validator.v9"
)

func ValidateNameTag(fl validator.FieldLevel) bool {
	// fmt.Println("FieldName:", fl.FieldName())
	// fmt.Println("StructFieldName", fl.StructFieldName())
	// fmt.Println("Parm:", fl.Param())
	value := fl.Field().String()
	err := ValidateName(value)
	if err != nil {
		return false
	}
	// fmt.Println(fl.Field())
	// fmt.Println(value)
	// fmt.Println(fl.GetStructFieldOKAdvanced2)
	return true
}
