package devicebuilder

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	sysfs "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/SysfsGathering"
	types "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Types"
	udev "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/UdevGathering"
	util "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Util"
)

func BuildNewBlockDevice(sysfs_block_path string) (*types.BlockDevice, error) {
	if !isBlockPath(sysfs_block_path) {
		return nil, fmt.Errorf("%s is not a block path", sysfs_block_path)
	}
	device := &types.BlockDevice{
		SysFSBlockPath: sysfs_block_path,
	}
	var errs []error

	err := sysfs.PopulateBlockDevice(device)
	errs = append(errs, err)

	err = PopulateBusAddress(device)
	errs = append(errs, err)

	udevData, err := NewUdevData(device.UDevId)
	errs = append(errs, err)
	device.Udev = udevData

	populateEmptySysfsWithUdev(device)

	return device, errors.Join(errs...)

}

func populateEmptySysfsWithUdev(device *types.BlockDevice) {
	if device.Serial == "" {
		device.Serial = types.SerialNumber(device.Udev.SerialShort)
	}
	if device.Model == "" {
		device.Model = device.Udev.Model
	}
}

// sets the devices bus addres from its /sys/device<....>/<busAddress> path
func PopulateBusAddress(device *types.BlockDevice) error {
	if device.SysFSDevicePath == "" {
		return fmt.Errorf("Attempted to get BusAddress but SysFSDevicePath is empty")
	}

	busAddy, err := types.NewBusAddress(filepath.Base(device.SysFSDevicePath))
	if err != nil {
		return err
	}
	device.Address = busAddy
	return nil
}

// build udev object from a udev id
func NewUdevData(id types.UDevID) (*types.UdevData, error) {
	if !util.HasUdev() {
		return nil, fmt.Errorf("No udev data directory")
	}
	udevID := string(id)
	udevPath := filepath.Join("/run/udev/data", udevID)
	if _, err := os.Stat(udevPath); err != nil {
		return nil, err
	}

	udevData := &types.UdevData{}
	udevData.DevID = udevID
	udevData.Raw = udev.EncodeUdevFile(udevID)

	if err := udev.PopulateUdevObject(udevPath, udevData); err != nil {
		return udevData, err
	}

	return udevData, nil
}
