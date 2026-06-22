package sysfs

import (
	"fmt"
	"path/filepath"
	"strings"

	types "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Types"
	util "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Util"
)

// Populates a types.BlockDevice device with as many fields is it can
func PopulateBlockDevice(device *types.BlockDevice) error {

	device.Name = filepath.Base(device.SysFSBlockPath)

	err := ResolveBlockDeviceSYSFSAttributes(device)
	if err != nil {
		return err
	}

	populateSYSFSDevicePath(device)
	err = populateBusAddress(device)
	if err != nil {
		return err
	}

	return nil
}

/*
evals the symlink at /sys/block/<X>/device to get the devices /sys/device/<.....> path

unless its nvme cause nvme is special boy
*/
func populateSYSFSDevicePath(device *types.BlockDevice) {

	if strings.HasPrefix(device.Name, "nvme") {
		device.SysFSDevicePath = device.SysFSBlockPath
	} else {
		device.SysFSDevicePath = util.ReadSymlink(filepath.Join(device.SysFSBlockPath, "device"))
	}
}

// sets the devices bus addres from its /sys/device<....>/<busAddress> path
func populateBusAddress(device *types.BlockDevice) error {
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
