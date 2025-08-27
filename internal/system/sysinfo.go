package system

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
)

func Sysinfo(args []string) {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("error user info:", err)
		return
	}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	fmt.Println("====================== SysInfo ======================")
	fmt.Printf("user: %s\n", usr.Username)
	fmt.Printf("home dir: %s\n", usr.HomeDir)
	fmt.Printf("hostname: %s\n", hostname)
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("arch: %s\n", runtime.GOARCH)
	fmt.Printf("go ver: %s\n", runtime.Version())
	fmt.Println("====================================================")
}
