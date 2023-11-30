package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"golang.org/x/sys/unix"
)

func main() {

	os := runtime.GOOS
	arch := runtime.GOARCH

	fmt.Println("Operating System (runtime.GOOS):", os)
	fmt.Println("Architecture (runtime.GOARCH):", arch)

	osCmd, osErr := getOSInfo()
	archCmd, archErr := getArchInfo()

	if osErr == nil {
		fmt.Println("Operating System (command line):", osCmd)
	} else {
		fmt.Println("Error retrieving operating system:", osErr)
	}

	if archErr == nil {
		fmt.Println("Architecture (command line):", archCmd)
	} else {
		fmt.Println("Error retrieving architecture:", archErr)
	}

	unixOS, unixOSErr := getUnixOs()
	unixArch, unixArchErr := getUnixArch()

	if unixOSErr == nil {
		fmt.Println("Unix OS:", unixOS)
	} else {
		fmt.Println("unixOSErr Error:", unixOSErr)
	}

	if unixArchErr == nil {
		fmt.Println("Unix Arch:", unixArch)
	} else {
		fmt.Println("unixOSErr Error:", unixArchErr)
	}
}

func getOSInfo() (string, error) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "ver")
	case "linux", "darwin":
		cmd = exec.Command("uname", "-s")
	default:
		return "", fmt.Errorf("unsupported operating system")
	}

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

func getArchInfo() (string, error) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "echo %PROCESSOR_ARCHITECTURE%")
	case "linux", "darwin":
		cmd = exec.Command("uname", "-m")
	default:
		return "", fmt.Errorf("unsupported operating system")
	}

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

func getUnixOs() (string, error) {
	var utsname unix.Utsname

	err := unix.Uname(&utsname)
	return strings.TrimRight(string(utsname.Sysname[:]), "\x00"), err
}

func getUnixArch() (string, error) {
	var utsname unix.Utsname
	err := unix.Uname(&utsname)
	return strings.TrimRight(string(utsname.Machine[:]), "\x00"), err
}
