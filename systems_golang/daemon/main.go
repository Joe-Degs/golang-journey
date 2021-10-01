package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// get the process id of this process
var pid = os.Getpid()

// fork parent process and start child process
func forkProcess() error {
	cmd := exec.Command(os.Args[0], "daemon")
	cmd.Stdout, cmd.Stderr, cmd.Dir = os.Stdout, os.Stderr, "/"
	return cmd.Start()
}

// release parent process resources so child process
// get adopted by the init process
func releaseResources() error {
	p, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	return p.Release()
}

// work done by the daemon
func runDaemon() {
	for {
		fmt.Println("[%d] Daemon mode\n", pid)
		time.Sleep(time.Second * 10)
	}
}

func main() {
	fmt.Printf("[%d] Start\n", pid)
	fmt.Printf("[%d] PPID: %d\n", pid, os.Getppid())
	defer fmt.Printf("[%d] Exit\n\n", pid)
	if len(os.Args) != 1 {
		// do the daemon work
		runDaemon()
		return
	}

	if err := forkProcess(); err != nil {
		fmt.Printf("[%d] Fork error: %s\n", pid, err)
		return
	}

	if err := releaseResources(); err != nil {
		fmt.Printf("[%d] Release error: %s\n", pid, err)
		return
	}
}
