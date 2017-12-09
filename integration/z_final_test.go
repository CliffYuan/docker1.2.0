package docker

import (
	"github.com/CliffYuan/docker1.2.0/utils"
	"runtime"
	"testing"
)

func displayFdGoroutines(t *testing.T) {
	t.Logf("Fds: %d, Goroutines: %d", utils.GetTotalUsedFds(), runtime.NumGoroutine())
}

func TestFinal(t *testing.T) {
	nuke(globalDaemon)
	t.Logf("Start Fds: %d, Start Goroutines: %d", startFds, startGoroutines)
	displayFdGoroutines(t)
}
