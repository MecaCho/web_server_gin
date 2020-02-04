package common

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"io"
	"net"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"web_server_gin/pkg/types"
)

// ValidateName check name string in request body
func ValidateName(name string) (err types.ErrorRespI) {
	ret, _ := regexp.MatchString(NameReg, name)
	if !ret {
		msg := fmt.Sprintf("Name (%s) invalid, resource name must consist of alphanumeric characters "+
			"or '-','_','/', length between 2-64 (e.g. 'my-name', 'ns/node1' or '123_abc'.", name)
		err = types.NewErrorResponse(NameInvalid, msg)
	}
	return
}

// ValidatePwd check name string in request body
func ValidatePwd(name string) (err types.ErrorRespI) {
	ret, _ := regexp.MatchString(PwdReg, name)
	if !ret {
		msg := fmt.Sprintf("Password invalid, must consist of alphanumeric characters " +
			"and length between 6 and 16.")
		err = types.NewErrorResponse(PwdInvalid, msg)
	}
	return
}

// ValidateAliasName check alias name
func ValidateAliasName(name string) (err types.ErrorRespI) {
	ret, _ := regexp.MatchString(AliasNameReg, name)
	if !ret {
		msg := fmt.Sprintf("Alias name (%s) invalid.", name)
		err = types.NewErrorResponse(AliasNameINvalid, msg)
	}
	return
}

// ValidateEmailAddr ...
func ValidateEmailAddr(email string) (err types.ErrorRespI) {
	ret, _ := regexp.MatchString(EmailReg, email)
	if !ret {
		msg := fmt.Sprintf("email (%s) invalid.", email)
		err = types.NewErrorResponse(EmailInvalid, msg)
	}
	return
}

// ValidateVersion check version...
func ValidateVersion(version string, isTag bool) (bool, types.ErrorRespI, string) {
	var respErr types.ErrorRespI
	var ret bool
	if isTag == true {
		ret, _ = regexp.MatchString(VersionTagReg, version)
	} else {
		ret, _ = regexp.MatchString(VersionReg, version)
	}
	if !ret {
		msg := fmt.Sprintf("Version (%s) invalid.", version)
		respErr = types.NewErrorResponse(VersionInvalid, msg)
	}
	return ret, respErr, version
}

// ValidateDesc check desc string...
func ValidateDesc(desc string) (bool, types.ErrorRespI, string) {
	var ret bool
	var resp types.ErrorRespI
	ret, _ = regexp.MatchString(DescReg, desc)
	if !ret {
		msg := fmt.Sprintf("Description (%s) invalid.", desc)
		resp = types.NewErrorResponse(DescInvalid, msg)
	}
	return ret, resp, desc
}

// ValidateIP ...
func ValidateIP(ip string) types.ErrorRespI {
	ret := net.ParseIP(ip)
	if ret == nil {
		msg := fmt.Sprintf("IP address (%s) invalid.", ip)
		return types.NewErrorResponse(IPInvalid, msg)
	}
	return nil
}

// ValidateRootRequestParams validate root request body params
func ValidateRootRequestParams(params interface{}) (err error) {
	valueInf := reflect.ValueOf(params)
	typeInf := reflect.TypeOf(params)

	// fmt.Println(valueInf.Type())
	// fmt.Println(typeInf.FieldByName("types.RootPutRequest"))

	for i := 0; i < typeInf.NumField(); i++ {
		fieldOfType := typeInf.Field(i)

		valueOfTypeInField := valueInf.FieldByName(fieldOfType.Name)
		fmt.Println(fieldOfType.Name, fieldOfType.Type, valueOfTypeInField.Type())

	}
	return
}

// ReadRequestBody read request body
func ReadRequestBody(requestBody io.Reader, modelPointer interface{}, allowUnknowField bool) error {
	decoder := json.NewDecoder(requestBody)
	if !allowUnknowField {
		decoder.DisallowUnknownFields()
	}
	err := decoder.Decode(modelPointer)
	if err != nil {
		glog.Errorf("Read request body error: %s.", err.Error())
		return err
	}
	return nil
}

// ValidateRequestBody ...
func ValidateRequestBody(requestBody io.Reader, modelPointer interface{}) types.ErrorRespI {

	decoder := json.NewDecoder(requestBody)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(modelPointer)

	//	decoder := NewDecoder(req.Request.Body)
	//	decoder.UseNumber()
	//	return decoder.Decode(v)

	// bytes.NewReader(byteData) convert byte to io.Reader

	// err := json.Unmarshal(requestBody, modelPointer)

	if err != nil {
		glog.Errorf("Decode request body error: %s.", err.Error())
		return types.NewErrorResponse(InvalidJSONFormat, fmt.Sprintf("Decode request body error : %s, please check the json format.", err.Error()))
	}

	return ValidateParams(modelPointer)
}

// ValidateParams base check parameters in request body
func ValidateParams(modelPointer interface{}) (err types.ErrorRespI) {
	glog.Infof("Validate request body parameters : , type is %s. \n", reflect.TypeOf(modelPointer))

	defer func() {
		errRe := recover()
		if errRe != nil {
			glog.Errorf("request body validate error: %+v, detail: %+v.", errRe, modelPointer)
			err = types.NewErrorResponse(RequestBodyInvalid, fmt.Sprintf("request body invalid: %+v", errRe))
		}
	}()
	valueInf := reflect.ValueOf(modelPointer)
	typeInf := reflect.TypeOf(modelPointer)

	if reflect.TypeOf(modelPointer).Kind() == reflect.Ptr {
		valueInf = reflect.ValueOf(modelPointer).Elem()
		typeInf = reflect.TypeOf(modelPointer).Elem()
	}

	// fmt.Println("num field of interface type: ", typeInf.NumField())
	for i := 0; i < typeInf.NumField(); i++ {

		fieldOfType := typeInf.Field(i)
		valueOfTypeInField := valueInf.FieldByName(fieldOfType.Name)

		typeKind := fieldOfType.Type.Kind()

		if typeKind == reflect.Struct {
			err := ValidateParams(valueOfTypeInField.Interface())
			if err != nil {
				return err
			}
		} else if typeKind == reflect.String {
			err := ValidateBaseParams(fieldOfType, valueOfTypeInField)
			if err != nil {
				return err
			}
		}
		if isRequired, ok := fieldOfType.Tag.Lookup("required"); ok && (isRequired == "true") {
			err := ValidateRequiredValue(typeKind, valueOfTypeInField, fieldOfType)
			if err != nil {
				return err
			}
		}

		max, ok := fieldOfType.Tag.Lookup("length")
		if ok {
			err := ValidateValueLength(typeKind, valueOfTypeInField, fieldOfType, max)
			if err != nil {
				return err
			}
		}

		if enum, ok := fieldOfType.Tag.Lookup("enum"); ok {
			if typeKind == reflect.String && valueOfTypeInField.String() != "" {
				enums := strings.Split(enum, ",")
				for _, value := range enums {
					if valueOfTypeInField.String() == value {
						return nil
					}
				}
				return types.NewErrorResponse(EnumsNotSupport, fmt.Sprintf("Request body invalid: unsupport value (%s) of (%s)", valueOfTypeInField, fieldOfType.Name))
			}
		}
	}
	return
}

// ValidateValueLength validate string length
func ValidateValueLength(typeKind reflect.Kind, valueOfTypeInField reflect.Value, fieldOfType reflect.StructField, max string) (err error) {
	maxLen, err := strconv.Atoi(max)
	if err != nil {
		glog.Errorf("parse int from string (%s) error, check tag of request body.", max)
		return
	}
	switch typeKind {
	case reflect.String:
		if len(valueOfTypeInField.String()) > maxLen {
			return types.NewErrorResponse(StringMissing, fmt.Sprintf("Request body invalid: string (%s) length can not larger than %d.", fieldOfType.Name, maxLen))
		}
	case reflect.Int, reflect.Int64:
		if valueOfTypeInField.Int() > int64(maxLen) {
			return types.NewErrorResponse(IntMissing, fmt.Sprintf("Request body invalid: int (%s) can not max than %d.", fieldOfType.Name, maxLen))
		}
	}
	return
}

// ValidateRequiredValue check required value
func ValidateRequiredValue(typeKind reflect.Kind, valueOfTypeInField reflect.Value, fieldOfType reflect.StructField) (err error) {
	switch typeKind {
	case reflect.String:
		if valueOfTypeInField.String() == "" {
			return types.NewErrorResponse(StringMissing, fmt.Sprintf("Request body invalid: missing string (%s) or can not be ''.", fieldOfType.Name))
		}
	case reflect.Int, reflect.Int64:
		if valueOfTypeInField.Int() == 0 {
			return types.NewErrorResponse(IntMissing, fmt.Sprintf("Request body invalid: missing int (%s) or can not be 0.", fieldOfType.Name))
		}
	case reflect.Struct:
		err := ValidateParams(valueOfTypeInField.Interface())
		if err != nil {
			return err
		}
	case reflect.Slice, reflect.Array:
		if valueOfTypeInField.Len() == 0 {
			return types.NewErrorResponse(ArrayMissing, fmt.Sprintf("Request body invalid: missing arrary (%s)", fieldOfType.Name))
		}
		for j := 0; j < valueOfTypeInField.Len(); j++ {
			err := ValidateParams(valueOfTypeInField.Index(j).Interface())
			if err != nil {
				return err
			}
		}
	}
	return
}

// ValidateBaseParams validate root request body params
func ValidateBaseParams(fieldOfType reflect.StructField, valueOfTypeInField reflect.Value) (err error) {
	if fieldOfType.Name == "Name" && valueOfTypeInField.String() != "" {
		err := ValidateName(valueOfTypeInField.String())
		if err != nil {
			return err
		}
	}

	// if fieldOfType.Name == "Description" && valueOfTypeInField.String() != "" {
	// 	_, err, _ := ValidateDesc(valueOfTypeInField.String())
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	return
}
