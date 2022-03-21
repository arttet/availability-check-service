//go:build !windows

package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arttet/validator-service/cmd/validatord/cmd"
)

func main() {
	id, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0) // nolint:errcheck
	if id == 0 {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

		for {
			select {
			case <-quit:
				return
			case <-ticker.C:
				go cmd.Execute()
			}
		}
	} else {
		cmd.Execute()
	}
}
