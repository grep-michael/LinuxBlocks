package devicebuilder

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	types "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Types"
	util "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Util"
)

func PopulateBlockDevice(device *types.BlockDevice) error {
	objValue := reflect.ValueOf(device).Elem()
	objType := objValue.Type()
	var errs []error

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		tag, ok := field.Tag.Lookup("sysfs")
		if !ok {
			continue
		}
		deviceField := objValue.Field(i)

		value, err := loadFieldValue(deviceField, filepath.Join(device.SysFSBlockPath, tag))
		if err != nil {
			errs = append(errs, fmt.Errorf("field %s (%s): %w", field.Name, tag, err))
			continue
		}
		setFieldValue(deviceField, value)
	}
	return errors.Join(errs...)
}

func loadFieldValue(deviceField reflect.Value, path string) (string, error) {
	switch deviceField.Addr().Interface().(type) {
	case *types.SerialNumber:
		serial, err := loadSerial(path)
		if err != nil {
			return "", err
		}

		return util.NormalizeSpaces(serial), nil
	default:
		value, err := loadTextFile(path)
		if err != nil {
			return "", err
		}
		return util.NormalizeSpaces(value), nil
	}
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

func loadSerial(sysfs_device_path string) (string, error) {
	//we assume serial and vpd_page80 are in the same location
	serial, err := loadTextFile(sysfs_device_path)
	if err == nil {
		return serial, nil
	}
	raw, err := os.ReadFile(
		filepath.Join(filepath.Dir(sysfs_device_path), "vpd_pg80"),
	)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(
		strings.TrimSuffix(
			string(raw[5:]), "\n"),
	), nil
}

func loadTextFile(path string) (string, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("couldnt load %s", path)
	}
	raw_s := strings.TrimSuffix(string(raw), "\n")
	return strings.TrimSpace(raw_s), nil
}
