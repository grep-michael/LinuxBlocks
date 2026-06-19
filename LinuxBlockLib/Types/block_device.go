package types

type BusType string

const (
	BusSCSI BusType = "scsi"
	BusNVMe BusType = "nvme"
	BusMMC  BusType = "MMC"
)

type BlockDevice struct {
	Name            string //sda, nvme, mmc, sr0, etc
	DevPath         string // /dev/<name>
	SysFSBlockPath  string // /sys/block/<name>
	SysFSDevicePath string // /sys/device/...

	Vendor string       `sysfs:"device/vendor"`
	Model  string       `sysfs:"device/model"`
	Serial SerialNumber `sysfs:"device/serial"`

	BlockCount    int64 `sysfs:"size"`
	PhysBlockSize int   `sysfs:"queue/physical_block_size"`
	LogiBlockSize int   `sysfs:"queue/logical_block_size"`
	SizeBytes     int64 // SectorCount * SectorSize
	Removable     bool  `sysfs:"removable"`
	Rotational    bool  `sysfs:"queue/rotational"`

	Bus     BusType
	Address BusAddress
}

type SerialNumber string
