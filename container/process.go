//go:build !(linux || darwin)

package container

import (
	"os/exec"
)

// NewParentProcess 是跨平台的公共接口
func NewParentProcess(tty bool, command string) *exec.Cmd {
	panic("unsupported platform")
}
