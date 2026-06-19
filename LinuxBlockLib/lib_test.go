package linuxblocklib

import (
	"fmt"
	"testing"

	res "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/DeviceResolver"
)

func TestSD(t *testing.T) {
	fmt.Println(res.ResolveBlockPath("/sys/devices/pci0000:00/0000:00:01.0/0000:01:00.0/host0/port-0:0/end_device-0:0/target0:0:0/0:0:0:0/block/sda"))
	fmt.Println(res.ResolveBlockPath("/sys/devices/pci0000:00/0000:00:01.0/0000:01:00.0/host0/port-0:0/end_device-0:0/target0:0:0/0:0:0:0"))
}
