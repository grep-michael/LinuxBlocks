package devicebuilder

import (
	"errors"
	"fmt"
	"path/filepath"

	sysfs "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/SysfsGathering"
	types "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Types"
	udev "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/UdevGathering"
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

	udevData, err := udev.NewUdevData(device.UDevId)
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
