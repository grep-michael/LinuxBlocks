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
)

func ResolveBlockPath(path string) error {
	if !isBlockPath(path) {
		return fmt.Errorf("%s is not a block path, No subsystem pointer to /sys/class/block", path)
	}

	return nil
}

func PopulateBlockDevice(device *types.BlockDevice) error {
	value := reflect.ValueOf(device).Elem()
	typ := value.Type()
	var errs []error
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag, ok := field.Tag.Lookup("sysfs")
		if !ok {
			continue
		}
		deviceField := value.Field(i)
		var deviceValue string
		switch deviceField.Addr().Interface().(type) {
		case *types.SerialNumber:
			serial, err := loadSerial(filepath.Join(device.SysFSBlockPath, tag))
			if err != nil {
				errs = append(errs, fmt.Errorf("field %s (%s): %w", field.Name, tag, err))
				continue
			}
			deviceValue = serial
		default:
			value, err := loadTextFile(filepath.Join(device.SysFSBlockPath, tag))
			if err != nil {
				errs = append(errs, fmt.Errorf("field %s (%s): %w", field.Name, tag, err))
				continue
			}
			deviceValue = value
		}

		//raw, err := os.ReadFile(filepath.Join(device.SysFSBlockPath, tag))
		//if err != nil {
		//	errs = append(errs, fmt.Errorf("field %s (%s): %w", field.Name, tag, err))
		//	continue
		//}
		//deviceValue := strings.TrimSpace(string(raw))

		switch deviceField.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			n, err := strconv.ParseInt(deviceValue, 10, 64)
			if err != nil {
				errs = append(errs, fmt.Errorf("field %s (%s): parsing %q: %w", field.Name, tag, deviceValue, err))
				continue
			}
			deviceField.SetInt(n)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			n, err := strconv.ParseUint(deviceValue, 10, 64)
			if err != nil {
				errs = append(errs, fmt.Errorf("field %s (%s): parsing %q: %w", field.Name, tag, deviceValue, err))
				continue
			}
			deviceField.SetUint(n)

		case reflect.Bool:
			deviceField.SetBool(deviceValue == "1")

		case reflect.String:
			deviceField.SetString(deviceValue)
		default:
			errs = append(errs, fmt.Errorf("field %s (%s): unsupported kind %s", field.Name, tag, deviceField.Kind()))
			continue

		}
	}
	return errors.Join(errs...)
}
func loadSerial(path string) (string, error) {
	//we assume serial and vpd_page80 are in the same location
	serial, err := loadTextFile(path)
	if err == nil {
		return serial, nil
	}
	raw, err := os.ReadFile(
		filepath.Join(filepath.Dir(path), "vpd_pg80"),
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
