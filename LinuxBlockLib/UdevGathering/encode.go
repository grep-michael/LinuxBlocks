package udev

import (
	"encoding/base64"
	"os"
	"path/filepath"
)

//b64 encode udev data file just for data retention

func EncodeUdevFile(udevID string) string {
	path := filepath.Join("/run/udev/data", udevID)
	raw, err := os.ReadFile(path)
	if err != nil {
		return err.Error()
	}
	return base64.StdEncoding.EncodeToString(raw)
}
