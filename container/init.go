//go:build !(linux || darwin)

package container

import "github.com/sirupsen/logrus"

func RunContainerInitProcess(command string, args []string) error {
	logrus.Errorf("unsupported platform")
	return nil
}