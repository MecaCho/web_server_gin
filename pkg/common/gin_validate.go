package common

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/glog"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"sync"
)

func TimeValidate(v *validator.FieldLevel) bool {

	return true
}

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var NewStructValidator binding.StructValidator = &defaultValidator{}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {
		v.lazyinit()
		if err := v.validate.Struct(obj); err != nil {
			structType := GetTypeName(obj)
			// validationErrors := err.(validator.ValidationErrors)
			// // validationErrors.Translate("123")
			// for _, e := range validationErrors {
			// 	// can translate each error one at a time.
			// 	fmt.Println(e.Translate(trans))
			// }
			glog.Errorf("Validate %s error: %s.", structType, err.Error())
			return error(err)
		}
	}
	return nil
}

func (v *defaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")
		v.validate.RegisterValidation("nameReg", ValidateNameTag)
	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

func GetTypeName(value interface{}) (name string) {
	if t := reflect.TypeOf(value); t.Kind() == reflect.Ptr {
		name = "*" + t.Elem().Name()
	} else {
		name = t.Name()
	}
	fmt.Println(name)
	return
}
