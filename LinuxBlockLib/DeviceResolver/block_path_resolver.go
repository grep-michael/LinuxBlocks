package deviceresolver

import "fmt"

func ResolveBlockPath(path string) error {
	if !isBlockPath(path) {
		return fmt.Errorf("%s is not a block path, No subsystem pointer to /sys/class/block", path)
	}

	return nil
}
