package linuxblocklib

import (
	"os"
	"path/filepath"
	"testing"

	devicebuilder "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/DeviceBuilder"
	types "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Types"
	util "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Util"
)

func TestPpopulateBlockDevice(t *testing.T) {

	dev := &types.BlockDevice{
		SysFSBlockPath: "/sys/devices/pci0000:00/0000:00:01.0/0000:01:00.0/host0/port-0:0/end_device-0:0/target0:0:0/0:0:0:0/block/sda",
	}
	//util.PrintObj(dev)
	err := devicebuilder.PopulateBlockDevice(dev)
	if err != nil {
		t.Error(err)
	}
	util.PrintObj(dev)
}

func TestPpopulateBlockDeviceLoop(t *testing.T) {

	var devices []*types.BlockDevice

	entries, err := os.ReadDir("/sys/block/")
	if err != nil {
		t.Error(err)
	}

	for _, file := range entries {
		path := util.ReadSymlink(filepath.Join("/sys/block/", file.Name()))
		dev := &types.BlockDevice{
			SysFSBlockPath: path,
		}
		err := devicebuilder.PopulateBlockDevice(dev)
		if err != nil {
		}
		devices = append(devices, dev)
	}

	util.PrintObj(devices)
}
