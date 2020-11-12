package main

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
)

func main() {
	platform, family, version, err := host.PlatformInformation()

	if err != nil {
		fmt.Printf("Error getting platform information: %v\n")
	} else {
		fmt.Printf("Platform: %s\n", platform)
		fmt.Printf("Family:   %s\n", family)
		fmt.Printf("Version:  %s\n", version)
	}

	kernelVersion, err := host.KernelVersion()

	if err != nil {
		fmt.Printf("Error getting kernel version info: %v\n", err)
	} else {
		fmt.Printf("Kernel Version: %s\n", kernelVersion)
	}

}
