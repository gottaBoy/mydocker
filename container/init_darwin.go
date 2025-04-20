//go:build darwin

package container

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

func RunContainerInitProcess(command string, args []string) error {
	logrus.Infof("command %s", command)

	// 通过 Docker 模拟 Linux 环境
	dockerArgs := []string{"run", "--rm", "-it", "alpine", command}
	dockerArgs = append(dockerArgs, args...)
	cmd := exec.Command("docker", dockerArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		logrus.Errorf("docker run error: %v", err)
		return err
	}
	return nil
}
