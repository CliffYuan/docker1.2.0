// +build !linux

package lxc

import "github.com/CliffYuan/docker1.2.0/daemon/execdriver"

func setHostname(hostname string) error {
	panic("Not supported on darwin")
}

func finalizeNamespace(args *execdriver.InitArgs) error {
	panic("Not supported on darwin")
}
