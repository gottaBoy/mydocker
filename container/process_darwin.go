//go:build darwin

package container

import (
	"os"
	"os/exec"
)

func NewParentProcess(tty bool, command string) *exec.Cmd {
	// 在 macOS 上通过 Docker 模拟命名空间隔离
	args := []string{"run", "--rm", "-it", "alpine", "sh", "-c", command}
	cmd := exec.Command("docker", args...)
	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd
}
