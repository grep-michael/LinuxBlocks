package sd

import (
	drivers "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Drivers"
	types "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Types"
)

type SDHandler struct {
}

func (d *SDHandler) GetName() string {
	return "SCSI Disk"
}
func (d *SDHandler) LoadDrives() []*types.BlockDevice {
	return nil
}

func init() {
	drivers.RegisterHandler("sd", &SDHandler{})
}
