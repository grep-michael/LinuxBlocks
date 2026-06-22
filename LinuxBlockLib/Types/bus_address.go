package types

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var NVME_REG = regexp.MustCompile(`nvme(\d+)n(\d+)`)
var HCTL_REG = regexp.MustCompile(`(\d):(\d):(\d):(\d)`)
var MMC_REG = regexp.MustCompile(`mmc(\d):(\d{4})`)

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
	Controller int
	Namespace  int
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

func NewBusAddress(address string) (BusAddress, error) {
	switch {
	case strings.HasPrefix(address, "nvm"):
		return buildNVME(address)
	case strings.HasPrefix(address, "mmc"):
		return buildMMC(address)
	default: //assume h:c:t:l
		return buildHCTL(address)
	}
}

func buildNVME(address string) (BusAddress, error) {
	matches := NVME_REG.FindStringSubmatch(address)
	if matches == nil {
		return nil, fmt.Errorf("%s didnt match nvme regex", address)
	}
	controller, _ := strconv.Atoi(matches[1])
	namespace, _ := strconv.Atoi(matches[2])
	return &NVMeAddress{
		Controller: controller,
		Namespace:  namespace,
	}, nil
}
func buildHCTL(address string) (BusAddress, error) {
	matches := HCTL_REG.FindStringSubmatch(address)
	if matches == nil {
		return nil, fmt.Errorf("%s didnt match hctl regex", address)
	}
	host, _ := strconv.Atoi(matches[1])
	channel, _ := strconv.Atoi(matches[2])
	target, _ := strconv.Atoi(matches[3])
	lun, _ := strconv.Atoi(matches[4])
	return &SCSIAddress{
		Host:    host,
		Channel: channel,
		Target:  target,
		Lun:     lun,
	}, nil

}
func buildMMC(address string) (BusAddress, error) {
	matches := MMC_REG.FindStringSubmatch(address)
	if matches == nil {
		return nil, fmt.Errorf("%s didnt match mmc regex", address)
	}
	controller, _ := strconv.Atoi(matches[1])
	cardAddress, _ := strconv.Atoi(matches[2])
	return &MMCAddress{
		Controller:  controller,
		CardAddress: cardAddress,
	}, nil
}
