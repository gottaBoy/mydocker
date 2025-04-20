package main

import (
	"github.com/gottaboy/mydocker/container"
	log "github.com/sirupsen/logrus"
	"os"
)

func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	err := parent.Wait()
	if err != nil {
		return
	}
	os.Exit(-1)
}
