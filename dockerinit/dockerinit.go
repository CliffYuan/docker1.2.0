package main

import (
	_ "github.com/CliffYuan/docker1.2.0/daemon/execdriver/lxc"
	_ "github.com/CliffYuan/docker1.2.0/daemon/execdriver/native"
	"github.com/CliffYuan/docker1.2.0/reexec"
)

func main() {
	// Running in init mode
	reexec.Init()
}
