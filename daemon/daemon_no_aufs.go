// +build exclude_graphdriver_aufs

package daemon

import (
	"github.com/CliffYuan/docker1.2.0/daemon/graphdriver"
)

func migrateIfAufs(driver graphdriver.Driver, root string) error {
	return nil
}
