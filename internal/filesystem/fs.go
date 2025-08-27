package filesystem

import (
	"fmt"
	"os"
)

func Cd(args []string) {
	if len(args) == 0 {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("error dir home:", err)
			return
		}
		os.Chdir(home)
		return
	}

	if err := os.Chdir(args[0]); err != nil {
		fmt.Println("cd error:", err)
	}
}
