package types

type BlockDevice struct {
	Name            string       `json:"Name"`
	Type            BlockDevType `json:"Type"`
	SYSFSDevicePath string       `json:"DevicePath"`
}

type BlockDevType string

const (
	NVMEType BlockDevType = "NVMe"
	SCSIType BlockDevType = "SCSI" //sata, sas, usb, some card readers, all handled by the scsi host, aka the sd driver
	MMCType  BlockDevType = "MMC"
)
