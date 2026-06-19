package types

import (
	"fmt"
)

type BusAddress interface {
	String() string
	BusType() BusType
	Part(int) string
}

type SCSIAddress struct {
	Host    int
	Channel int
	Target  int
	Lun     int
}

func (a SCSIAddress) Part(i int) string {
	switch i {
	case 0:
		return fmt.Sprintf("%d", a.Host)
	case 1:
		return fmt.Sprintf("%d", a.Channel)
	case 2:
		return fmt.Sprintf("%d", a.Target)
	case 3:
		return fmt.Sprintf("%d", a.Lun)
	default:
		return a.String()
	}
}
func (a SCSIAddress) String() string {
	return fmt.Sprintf("%d:%d:%d:%d", a.Host, a.Channel, a.Target, a.Lun)
}

func (a SCSIAddress) BusType() BusType { return BusSCSI }

type NVMeAddress struct {
	Controller int //nvmeX
	Namespace  int //nvme0nY
}

func (a NVMeAddress) Part(i int) string {
	switch i {
	case 0:
		return fmt.Sprintf("nvme%d", a.Controller)
	case 1:
		return fmt.Sprintf("%d", a.Namespace)
	default:
		return a.String()
	}
}
func (a NVMeAddress) String() string {
	return fmt.Sprintf("nvme%d/%d", a.Controller, a.Namespace)
}

func (a NVMeAddress) BusType() BusType { return BusNVMe }

type MMCAddress struct {
	Controller  int
	CardAddress int
}

func (a MMCAddress) Part(i int) string {
	switch i {
	case 0:
		return fmt.Sprintf("mmc%d", a.Controller)
	case 1:
		return fmt.Sprintf("%04d", a.CardAddress)
	default:
		return a.String()
	}
}
func (a MMCAddress) String() string {
	return fmt.Sprintf("mmc%d/%04d", a.Controller, a.CardAddress)
}
func (a MMCAddress) BusType() BusType { return BusMMC }
