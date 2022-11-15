package main

import (
	"os/exec"
)

func reboot() error {
	return exec.Command("cmd", "/C", "shutdown", "/r", "/f", "/t", "0").Run()
}
