package udev

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func PopulateUdevObject(udevFile string, obj any) error {
	udevMap, err := generateUdevMap(udevFile)
	if err != nil {
		return err
	}
	return populateValueFromProps(reflect.ValueOf(obj).Elem(), udevMap)
}

func populateValueFromProps(value reflect.Value, props map[string]string) error {
	objType := value.Type()
	var errs []error
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		objField := value.Field(i)

		tag, ok := field.Tag.Lookup("udev")
		if ok {
			setFieldValue(objField, props[tag])
		} else {
			switch {
			//nested udev struct is a pointer
			case objField.Kind() == reflect.Ptr &&
				objField.Type().Elem().Kind() == reflect.Struct:
				if objField.IsNil() { //allocate if nil
					objField.Set(reflect.New(objField.Type().Elem()))
				}
				errs = append(errs, populateValueFromProps(objField.Elem(), props))
			//not a pointer but just a struct
			case objField.Kind() == reflect.Struct:
				errs = append(errs, populateValueFromProps(objField, props))
			}
		}
	}
	return errors.Join(errs...)
}

func generateUdevMap(udevFile string) (map[string]string, error) {
	file, err := os.Open(udevFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	properties := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		prefix, pair, ok := strings.Cut(scanner.Text(), ":")
		if prefix != "E" || !ok {
			continue
		}
		if key, value, ok := strings.Cut(pair, "="); ok {
			properties[key] = value

		}
	}
	return properties, nil
}

func setFieldValue(field reflect.Value, s string) error {
	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return fmt.Errorf("parsing %q: %w", s, err)
		}
		field.SetInt(n)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return fmt.Errorf("parsing %q: %w", s, err)
		}
		field.SetUint(n)
	case reflect.Bool:
		field.SetBool(s == "1")
	case reflect.String:
		field.SetString(s)
	default:
		return fmt.Errorf("unsupported kind %s", field.Kind())
	}
	return nil
}
