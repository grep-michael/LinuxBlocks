package sd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	types "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Types"
	util "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Util"
)

func ResolveSDDevice(devicePath string) *types.BlockDevice {
	device := &types.BlockDevice{}
	device.Name = filepath.Base(devicePath)
	device.Hctl = filepath.Base(
		util.ReadSymlink(
			filepath.Dir(filepath.Dir(devicePath)),
		),
	)
	device.Type = lookupProcName(device.Hctl)
	device.SYSFSDevicePath = devicePath
	device.SYSFSDriverPath = util.ReadSymlink(
		filepath.Join(devicePath, "device/driver"),
	)

	return device
}

func lookupProcName(hctl string) types.BlockDevType {
	proc, err := os.ReadFile(fmt.Sprintf("/sys/class/scsi_host/host%c/proc_name", hctl[0]))
	if err != nil {
		return types.GENERICType
	}
	switch {
	case strings.Contains(string(proc), "sas"):
		return types.SASType
	case strings.Contains(string(proc), "ahci"):
		return types.SATAType
	default:
		log.Printf("Unhandled proc_name: %s\n", proc)
		return types.SCSIType
	}
}
