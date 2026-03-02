package utils

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func IsProcessRunningWindows(processName string) bool {
	cmd := exec.Command("wmic", "process", "get", "Caption,ProcessId,CommandLine")

	// Run the command and capture output
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	//fmt.Print(string(output))
	return strings.Contains(string(output), processName)
}

func ContainerIsRunning(container string) (bool, error) {
	res, er := ExecuteCommand("docker ps")
	if er != nil {
		return false, er
	}
	if strings.Contains(res, container) {
		return true, nil
	}
	return false, nil
}

func ExecuteCommand(cmdStr string) (res string, er error) {
	cmd := ExecuteCmd(cmdStr)
	if cmd == nil {
		return "", fmt.Errorf("couldn't execute Command %s ", cmdStr)
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), err
}

func ExecuteCmd(cmdStr string) *exec.Cmd {
	switch runtime.GOOS {
	case "windows":
		return exec.Command("cmd", "/C", cmdStr)
	case "linux":
		return exec.Command("sh", "-c", cmdStr)
	}
	return nil
}

func NowAsUnixMilli() int64 {
	return time.Now().UnixNano() / 1e6
}
