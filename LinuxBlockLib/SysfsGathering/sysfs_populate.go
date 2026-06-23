package sysfs

import (
	"path/filepath"
	"strings"

	types "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Types"
	util "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Util"
)

// Populates a types.BlockDevice device with sysfs fields
func PopulateBlockDevice(device *types.BlockDevice) error {

	device.Name = filepath.Base(device.SysFSBlockPath)

	device.DevPath = filepath.Join("/dev/", device.Name)

	err := PopulateSYSFSAttributes(device)
	if err != nil {
		return err
	}

	PopulateSYSFSDevicePath(device)

	device.SizeBytes = device.BlockCount * int64(device.PhysBlockSize)

	return nil
}

/*
evals the symlink at /sys/block/<X>/device to get the devices /sys/device/<.....> path

unless its nvme cause nvme is special boy
*/
func PopulateSYSFSDevicePath(device *types.BlockDevice) {
	var path string
	switch {
	case strings.HasPrefix(device.Name, "nvme"):
		path = device.SysFSBlockPath
	default:
		path = filepath.Join(device.SysFSBlockPath, "device")
	}
	device.SysFSDevicePath = util.ReadSymlink(path)
}
