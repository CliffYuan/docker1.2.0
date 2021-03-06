package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os/exec"
	"testing"

	"github.com/CliffYuan/docker1.2.0/daemon"
)

func TestNetworkNat(t *testing.T) {
	iface, err := net.InterfaceByName("eth0")
	if err != nil {
		t.Skip("Test not running with `make test`. Interface eth0 not found: %s", err)
	}

	ifaceAddrs, err := iface.Addrs()
	if err != nil || len(ifaceAddrs) == 0 {
		t.Fatalf("Error retrieving addresses for eth0: %v (%d addresses)", err, len(ifaceAddrs))
	}

	ifaceIp, _, err := net.ParseCIDR(ifaceAddrs[0].String())
	if err != nil {
		t.Fatalf("Error retrieving the up for eth0: %s", err)
	}

	runCmd := exec.Command(dockerBinary, "run", "-d", "-p", "8080", "busybox", "nc", "-lp", "8080")
	out, _, err := runCommandWithOutput(runCmd)
	errorOut(err, t, fmt.Sprintf("run1 failed with errors: %v (%s)", err, out))

	cleanedContainerID := stripTrailingCharacters(out)

	inspectCmd := exec.Command(dockerBinary, "inspect", cleanedContainerID)
	inspectOut, _, err := runCommandWithOutput(inspectCmd)
	errorOut(err, t, fmt.Sprintf("out should've been a container id: %v %v", inspectOut, err))

	containers := []*daemon.Container{}
	if err := json.Unmarshal([]byte(inspectOut), &containers); err != nil {
		t.Fatalf("Error inspecting the container: %s", err)
	}
	if len(containers) != 1 {
		t.Fatalf("Unepexted container count. Expected 0, recieved: %d", len(containers))
	}

	port8080, exists := containers[0].NetworkSettings.Ports["8080/tcp"]
	if !exists || len(port8080) == 0 {
		t.Fatal("Port 8080/tcp not found in NetworkSettings")
	}

	runCmd = exec.Command(dockerBinary, "run", "-p", "8080", "busybox", "sh", "-c", fmt.Sprintf("echo hello world | nc -w 30 %s %s", ifaceIp, port8080[0].HostPort))
	out, _, err = runCommandWithOutput(runCmd)
	errorOut(err, t, fmt.Sprintf("run2 failed with errors: %v (%s)", err, out))

	runCmd = exec.Command(dockerBinary, "logs", cleanedContainerID)
	out, _, err = runCommandWithOutput(runCmd)
	errorOut(err, t, fmt.Sprintf("failed to retrieve logs for container: %v %v", cleanedContainerID, err))

	if expected := "hello world\n"; out != expected {
		t.Fatalf("Unexpected output. Expected: %s, recieved: -->%s<--", expected, out)
	}

	killCmd := exec.Command(dockerBinary, "kill", cleanedContainerID)
	out, _, err = runCommandWithOutput(killCmd)
	errorOut(err, t, fmt.Sprintf("failed to kill container: %v %v", out, err))
	deleteAllContainers()

	logDone("network - make sure nat works through the host")
}
