package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func PrintObj(obj any) {
	js, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		fmt.Printf("Failed to marshel obj %T\n", obj)
	}
	fmt.Println(string(js))
}

func HasUdev() bool {
	if _, err := exec.LookPath("udevadm"); err == nil {
		return true
	}
	if fi, err := os.Stat("/run/udev/data"); err == nil && fi.IsDir() {
		return true
	}
	return false
}

func ReadSymlink(path string) string {
	ret, err := filepath.EvalSymlinks(path)
	if err != nil {
		//error handling is for cowards
		log.Printf("Failed to read symlink %s\n\t%+v\n", path, err)
	}
	return ret
}

func FindDevices(driverPath string) ([]string, error) {
	entries, err := os.ReadDir(driverPath)
	if err != nil {
		return nil, fmt.Errorf("reading driver path %s: %w", driverPath, err)
	}

	var devices []string
	for _, entry := range entries {
		if entry.Type()&os.ModeSymlink == 0 {
			continue
		}
		resolved, err := filepath.EvalSymlinks(filepath.Join(driverPath, entry.Name()))
		if err != nil {
			return nil, fmt.Errorf("resolving symlink for %s: %w", entry.Name(), err)
		}
		if strings.HasPrefix(resolved, "/sys/devices/") {
			devices = append(devices, resolved)
		}
	}

	return devices, nil
}

func ReadInt64File(file string) (int64, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return 0, err
	}

	s := strings.TrimSpace(string(data))
	return strconv.ParseInt(s, 10, 64)
}

func NormalizeSpaces(in string) string {
	return strings.TrimSpace(
		strings.ReplaceAll(in, "    ", " "),
	)
}
