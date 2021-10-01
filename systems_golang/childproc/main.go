package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// exitStatus extracts the system information to get the exit status
// of a process
func exitStatus(state *os.ProcessState) int {
	status, ok := state.Sys().(syscall.WaitStatus)
	if !ok {
		return -1
	}
	return status.ExitStatus()
}

func main() {
	// create a command to run
	cmd := exec.Command("ls", "----al")

	// run command asynchronously
	if err := cmd.Run(); err != nil {

		// get process state and check the exit status
		if status := exitStatus(cmd.ProcessState); status == -1 {
			fmt.Println(err)
		} else {
			fmt.Println("Status", status)
		}
	}
}
