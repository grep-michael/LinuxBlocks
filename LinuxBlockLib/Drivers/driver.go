package drivers

import types "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Types"

type DriverHandler interface {
	GetName() string
	LoadDrives() []*types.BlockDevice
}
