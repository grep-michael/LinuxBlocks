package linuxblocklib

import (
	"testing"

	"github.com/grep-michael/LinuxBlocks/LinuxBlockLib/DeviceResolver/sd"
	util "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Util"
)

func TestSD(t *testing.T) {
	drive := sd.ResolveSDDevice("/sys/devices/pci0000:00/0000:00:01.0/0000:01:00.0/host0/port-0:0/end_device-0:0/target0:0:0/0:0:0:0/block/sda")
	util.PrintObj(drive)
}
