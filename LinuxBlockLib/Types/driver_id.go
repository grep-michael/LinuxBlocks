package types

type Driver struct {
	Name            string `json:"Name"`
	SYSFSDriverPath string `json:"Path"`
}

func LookUpDriverName(shortHand string) string {
	return DriverNameMap[shortHand]
}

var DriverNameMap = map[string]string{
	"nvme":   "NVMe",
	"mmcblk": "Multi-Media-Card Block Device",
	"sr":     "SCSI Read-Only",
	"sd":     "SCSI Disk",
}
