package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"syscall"

	_dbPackage "github.com/eminmuhammadi/memcache/db"
	ps "github.com/mitchellh/go-ps"
	"github.com/urfave/cli"
)

// The name of the executable
const EXECUTABLE_NAME = "memcache"

// Stops the memcache server.
func Stop() cli.Command {
	return cli.Command{
		Name:  "stop",
		Usage: "Stops the memcache server",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			return killServer()
		},
	}
}

// Kills the memcache server gracefully.
func killServer() error {
	processes, err := getServerProcess()
	// Find the pid of the server
	if err != nil {
		return err
	}

	// Kill the server gracefully
	for _, process := range processes {
		err = sendSignal(process)
		if err != nil {
			return err
		}
	}

	println(fmt.Sprintf("%s Memcache is shutting down", _dbPackage.TimeNowString()))

	return nil
}

// Returns the process of the memcache server.
func getServerProcess() ([]os.Process, error) {
	// The name of the executable
	var executable string = EXECUTABLE_NAME

	// The process of the server
	var memcacheProcess []os.Process

	if runtime.GOOS == "windows" {
		executable = fmt.Sprintf("%s.exe", executable)
	}

	// Find the pid of the server
	processes, err := ps.Processes()
	if err != nil {
		return memcacheProcess, err
	}

	for _, process := range processes {
		if process.Executable() == executable && process.Pid() != os.Getpid() {
			memcacheProcess = append(memcacheProcess, os.Process{Pid: process.Pid()})
		}
	}

	if len(memcacheProcess) == 0 {
		return memcacheProcess, fmt.Errorf("no memcache server is running")
	}

	return memcacheProcess, nil
}

// Sends a signal to the process.
func sendSignal(process os.Process) error {
	// TODO: Find a better way to kill the process
	if runtime.GOOS == "windows" {
		// cmd exec taskkill /F /PID 1234
		cmd := exec.Command("taskkill", "/F", "/PID", strconv.Itoa(process.Pid))
		return cmd.Run()
	}

	return process.Signal(syscall.SIGTERM)
}
