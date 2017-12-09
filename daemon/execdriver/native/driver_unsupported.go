// +build !linux

package native

import (
	"fmt"

	"github.com/CliffYuan/docker1.2.0/daemon/execdriver"
)

func NewDriver(root, initPath string) (execdriver.Driver, error) {
	return nil, fmt.Errorf("native driver not supported on non-linux")
}
