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

	Vendor string
	Model  string
	Serial string

	SectorCount uint64 //from device/size
	SectorSize  int
	SizeBytes   uint64 // SectorCount * SectorSize
	Removable   bool
	Rotational  bool

	Bus     BusType
	Address BusAddress
}
