//go:build linux

package container

import (
	"github.com/sirupsen/logrus"
	"os"
	"syscall"
)

func RunContainerInitProcess(command string, args []string) error {
	logrus.Infof("command %s", command)

	// 挂载 proc 文件系统（仅Linux）
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	if err := syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), ""); err != nil {
		logrus.Errorf("mount proc error: %v", err)
		return err
	}

	// 执行命令
	argv := []string{command}
	if err := syscall.Exec(command, argv, os.Environ()); err != nil {
		logrus.Errorf("exec error: %v", err)
		return err
	}
	return nil
}
