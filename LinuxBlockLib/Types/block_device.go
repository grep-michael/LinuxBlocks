package types

type BusType string

const (
	BusSCSI BusType = "scsi"
	BusNVMe BusType = "nvme"
	BusMMC  BusType = "mmc"
)

type BlockDevice struct {
	Name            string `json:"Name"`       //sda, nvme, mmc, sr0, etc
	DevPath         string `json:"DevPath"`    // /dev/<name>
	SysFSBlockPath  string `json:"BlockPath"`  // /sys/block/<name>
	SysFSDevicePath string `json:"DevicePath"` // /sys/device/...

	Vendor string       `json:"Vendor" sysfs:"device/vendor"`
	Model  string       `json:"Model" sysfs:"device/model"`
	Serial SerialNumber `json:"Serial" sysfs:"device/serial"`

	BlockCount    int64  `json:"BlockCount" sysfs:"size"`
	PhysBlockSize int    `json:"PhysBlockSize" sysfs:"queue/physical_block_size"`
	LogiBlockSize int    `json:"LogBlockSize" sysfs:"queue/logical_block_size"`
	SizeBytes     int64  `json:"SizeBytes"` // SectorCount * SectorSize
	Removable     bool   `json:"Removable" sysfs:"removable"`
	Rotational    bool   `json:"Rotation" sysfs:"queue/rotational"`
	UDevId        UDevID `json:"UDevID" sysfs:"dev"`

	Udev *UdevData `json:"UDevData"`

	Driver  Driver     `json:"Driver"`
	Address BusAddress `json:"BusAddress"`
}
type UDevID string
type SerialNumber string
