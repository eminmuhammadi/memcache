package cmd

import (
	"fmt"
	"os"
	"runtime"
	"syscall"

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
	process, err := getServerProcess()
	// Find the pid of the server
	if err != nil {
		return err
	}

	// Kill the server gracefully
	return sendSignal(process)
}

// Returns the process of the memcache server.
func getServerProcess() (os.Process, error) {
	// The name of the executable
	var executable string = EXECUTABLE_NAME

	// The process of the server
	var memcacheProcess os.Process

	if runtime.GOOS == "windows" {
		executable = fmt.Sprintf("%s.exe", executable)
	}

	// Find the pid of the server
	processes, err := ps.Processes()
	if err != nil {
		return memcacheProcess, err
	}

	for _, process := range processes {
		if process.Executable() == executable {
			memcacheProcess.Pid = process.Pid()
			break
		}
	}

	if memcacheProcess.Pid == 0 {
		return memcacheProcess, fmt.Errorf("%s is not running", executable)
	}

	return memcacheProcess, nil
}

// Sends a signal to the process.
func sendSignal(process os.Process) error {
	// TODO: Find a better way to kill the process
	if runtime.GOOS == "windows" {
		dll, err := syscall.LoadDLL("kernel32.dll")
		if err != nil {
			return err
		}

		proc, err := dll.FindProc("GenerateConsoleCtrlEvent")
		if err != nil {
			return err
		}

		_, _, err = proc.Call(uintptr(syscall.CTRL_C_EVENT), uintptr(process.Pid))

		return err
	}

	return process.Signal(syscall.SIGTERM)
}
