package devicebuilder

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type DeviceBuilder struct {
	SysfsBlockPath string
}

func NewDeviceBuilder(sysfsBlockPath string) (*DeviceBuilder, error) {
	if !isBlockPath(sysfsBlockPath) {
		return nil, fmt.Errorf("%s is Not a block device path", sysfsBlockPath)
	}
	return &DeviceBuilder{
		SysfsBlockPath: sysfsBlockPath,
	}, nil
}

func (b *DeviceBuilder) FetchBlockCount() int64 {
	sizeString, err := b.fetchBlockPath("size")
	if err != nil {
		log.Printf("Failed to get size in %s\n\t%+v\n", b.SysfsBlockPath, err)
		return -1
	}
	s := strings.TrimSpace(string(sizeString))
	sizeInt, e := strconv.ParseInt(s, 10, 64)
	if e != nil {
		log.Printf("Failed to convert size (%s) to int: %+v\n", s, e)
		return -1
	}
	return sizeInt
}
func (b *DeviceBuilder) FetchLogicalBlockSize() int {
	sizeString, err := b.fetchBlockPath("queue/logical_block_size")
	if err != nil {
		log.Printf("Failed to get queue/logical_block_size in %s\n\t%+v\n", b.SysfsBlockPath, err)
		return -1
	}
	s := strings.TrimSpace(string(sizeString))
	sizeInt, e := strconv.Atoi(s)
	if e != nil {
		log.Printf("Failed to convert logical_block_size (%s) to int: %+v\n", s, e)
		return -1
	}
	return sizeInt
}

func (b *DeviceBuilder) fetchBlockPath(path string) (string, error) {
	sizePath := filepath.Join(b.SysfsBlockPath, path)
	text, err := os.ReadFile(sizePath)
	if err != nil {
		return "", err
	}
	return string(text), nil
}
