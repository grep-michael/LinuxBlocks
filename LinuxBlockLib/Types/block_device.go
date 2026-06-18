package types

type BlockDevice struct {
	Name            string       `json:"Name"`
	Type            BlockDevType `json:"Type"`
	Hctl            string       `json:"HCTL"`
	SYSFSDevicePath string       `json:"DevicePath"`
	SYSFSDriverPath string       `json:"DriverPath"`
}

type BlockDevType string

const (
	NVMEType    BlockDevType = "NVMe"
	SASType     BlockDevType = "SAS"
	SATAType    BlockDevType = "SATA"
	SCSIType    BlockDevType = "SCSI" //sata, sas, usb, some card readers, all handled by the scsi host, aka the sd driver
	MMCType     BlockDevType = "MMC"
	GENERICType BlockDevType = "Generic"
)
