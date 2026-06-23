package udev

import (
	"os"
	"path/filepath"

	types "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Types"
)

func NewUdevData(id types.UDevID) (*types.UdevData, error) {
	udevID := string(id)
	udevPath := filepath.Join("/run/udev/data", udevID)
	if _, err := os.Stat(udevPath); err != nil {
		return nil, err
	}

	udevData := &types.UdevData{}
	udevData.DevID = udevID
	udevData.Raw = EncodeUdevFile(udevID)

	if err := PopulateUdevObject(udevPath, udevData); err != nil {
		return udevData, err
	}

	return udevData, nil
}
