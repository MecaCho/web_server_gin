package common

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"testing"
)

type Info struct {
	ID   int64  `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,nameValidate"`
}

func validateName(fl validator.FieldLevel) bool {
	fmt.Println("FieldName:", fl.FieldName())
	fmt.Println("StructFieldName", fl.StructFieldName())
	fmt.Println("Parm:", fl.Param())
	value := fl.Field().String()
	err := ValidateName(value)
	if err != nil {
		return false
	}
	fmt.Println(fl.Field())
	fmt.Println(value)
	// fmt.Println(fl.GetStructFieldOKAdvanced2)
	return true
}

func TestTimeValidate(t *testing.T) {
	var info Info
	info.ID = 123
	// info := Info{123, "test_name"}
	validatorHandle := validator.New()
	validatorHandle.RegisterValidation("nameValidate", validateName)
	err := validatorHandle.Struct(info)
	fmt.Println("without name result:", err)
	if err != nil {
		fmt.Println(err)
	}

	info.Name = "test_name"
	err = validatorHandle.Struct(info)
	fmt.Println("with name result:", err)
	if err != nil {
		fmt.Println(err)
	}

	info.Name = "test_name@#$$"
	err = validatorHandle.Struct(info)
	fmt.Println("with name result:", err)
	if err != nil {
		fmt.Println(err)
	}
}
