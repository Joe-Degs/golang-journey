package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("help")
	}
}

// this is a code along to a go talk about building container with the go
// programming language. The talk was given by the amazing Liz Rice.

// Some things i've learnt from my time spent reading and watching videos
// about container things have given some insights into what makes containers
// in the first place.

// Containers are just some combination of linux subsystems that help to provide
// process isolation and resource mgmt and isolation etc
// this subsystems include cgroups, namespaces

// if you really understand the internals of the linux operating systems and its
// subsystems and know your shell commands very well, you could put together some
// script that will give you some sort of container

func run() {
	fmt.Printf("Running %v as pid %d\n", os.Args[2:], os.Getpid())

	// /proc/self/exe is a link to the currently executing process.
	// so we reinvoke the same process but this time with child argument.
	// at the end of the day we run this same process twice and we will
	// see the logs to show that.

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("Running %v as pid %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Sethostname([]byte("container")))

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
