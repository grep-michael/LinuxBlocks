package devicebuilder

import (
	"fmt"
	"path/filepath"

	sysfs "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/SysfsGathering"
	types "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Types"
)

func BuildNewBlockDevice(sysfs_block_path string) (*types.BlockDevice, error) {
	if !isBlockPath(sysfs_block_path) {
		return nil, fmt.Errorf("%s is not a block path", sysfs_block_path)
	}
	device := &types.BlockDevice{
		SysFSBlockPath: sysfs_block_path,
	}
	err := sysfs.PopulateBlockDevice(device)

	err = PopulateBusAddress(device)
	if err != nil {
		return nil, err
	}

	return device, err

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
