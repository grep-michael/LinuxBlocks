package devicebuilder

import (
	"fmt"

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

	return device, err

}
