package linuxblocklib

import (
	"testing"

	devicebuilder "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/DeviceBuilder"
	res "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/DeviceBuilder"
	types "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Types"
	util "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Util"
)

func TestSD(t *testing.T) {
	res.ResolveBlockPath("/sys/devices/pci0000:00/0000:00:01.0/0000:01:00.0/host0/port-0:0/end_device-0:0/target0:0:0/0:0:0:0/block/sda")

	dev := &types.BlockDevice{
		SysFSBlockPath: "/sys/devices/pci0000:00/0000:00:01.0/0000:01:00.0/host0/port-0:0/end_device-0:0/target0:0:0/0:0:0:0/block/sda",
	}
	util.PrintObj(dev)
	err := devicebuilder.PopulateBlockDevice(dev)
	if err != nil {
		t.Error(err)
	}
	util.PrintObj(dev)
}
