package deviceresolver

import (
	util "github.com/grep-michael/LinuxBlocks/LinuxBlockLib/Util"
	"path/filepath"
	"strings"
)

func isBlockPath(path string) bool {
	subsysPah := filepath.Join(path, "subsystem")
	eval := util.ReadSymlink(subsysPah)
	return strings.HasSuffix(eval, "block") //if subsystem resolves to /sys/class/block its a block device path
}
