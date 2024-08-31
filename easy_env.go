package easyenv

import (
	"errors"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvHelperInterface interface {
	GetKeyHelper() string
	Execute(value string, incoming_value string) (string, error)
}

type Env struct {
	Data map[string]string
}

var Halpers = []EnvHelperInterface{
	&EnvDefault{},
	&EnvRequired{},
}

func NewDecoder(filePath string) (*Env, error) {
	fileIsExist := checkFileExists(filePath)
	if !fileIsExist {
		return nil, errors.New("file not found")
	}

	data, err := godotenv.Read(filePath)
	if err != nil {
		return nil, err
	}

	return &Env{Data: data}, nil
}

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

func (e *Env) Load(target interface{}) error {
	targetType := reflect.TypeOf(target)
	if targetType.Kind() != reflect.Ptr {
		return errors.New("target must be a pointer")
	}

	targetStruct := targetType.Elem()
	targetStructValue := reflect.ValueOf(target).Elem()

	for i := 0; i < targetStruct.NumField(); i++ {
		field := targetStruct.Field(i)
		incoming_value := e.Data[field.Tag.Get("env")]

		for _, helper := range Halpers {

			value := field.Tag.Get(helper.GetKeyHelper())
			if value == "" {
				continue
			}

			helper_value, err := helper.Execute(value, incoming_value)
			if err != nil {
				return err
			}

			incoming_value = helper_value
		}

		convertValue, error := convertForString(field.Type, incoming_value)
		if error != nil {
			return error
		}

		targetStructValue.FieldByName(field.Name).Set(reflect.ValueOf(convertValue))
	}

	return nil
}

func convertForString(typeValue reflect.Type, value string) (any, error) {
	switch typeValue.String() {
	case "int":
		return strconv.Atoi(value)
	case "float64":
		return strconv.ParseFloat(value, 64)
	case "string":
		return value, nil
	default:
		return nil, errors.New("unsupported type")
	}
}
