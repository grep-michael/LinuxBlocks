package sysfs

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	types "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Types"
	util "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Util"
)

/*
 for the purpose of this project a sysfs attribute is a data value relating to a storage device thats available via reading sysfs files
 for example the files device/serial, device/model, device/vendor, etc
*/

// populates the types.BlockDevice fields that have a sysfs tag
func PopulateSYSFSAttributes(device *types.BlockDevice) error {
	objValue := reflect.ValueOf(device).Elem()
	objType := objValue.Type()

	if device.SysFSBlockPath == "" {
		return fmt.Errorf("Cant resolve sysfs attributes if theres no sysfs path")
	}

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		tag, ok := field.Tag.Lookup("sysfs")
		if !ok {
			continue
		}
		deviceField := objValue.Field(i)

		value, err := loadFieldValue(deviceField, filepath.Join(device.SysFSBlockPath, tag))
		if err != nil {
			//unsure whether is better to log or error if for say a vendor file is missing from a cd/rom or some jank device
			//errs = append(errs, fmt.Errorf("field %s (%s): %w", field.Name, tag, err))
			continue
		}
		setFieldValue(deviceField, value)
	}
	return nil
}

func loadFieldValue(deviceField reflect.Value, path string) (string, error) {
	switch deviceField.Addr().Interface().(type) {
	case *types.SerialNumber:
		serial, err := loadSerial(path)
		if err != nil {
			return "", err
		}

		return util.NormalizeSpaces(serial), nil
	case *types.UDevID:
		value, err := loadTextFile(path)
		if err != nil {
			return "", err
		}
		return "b" + util.NormalizeSpaces(value), nil
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

func loadSerial(sysfs_serial_path string) (string, error) {
	serial, err := loadTextFile(sysfs_serial_path)
	if err == nil {
		return serial, nil
	}
	raw, err := os.ReadFile(
		//we assume serial and vpd_page80 are in the same location
		filepath.Join(filepath.Dir(sysfs_serial_path), "vpd_pg80"),
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
