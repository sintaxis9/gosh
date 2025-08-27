package fs

import (
	"fmt"
	"os"
)

func CheckExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func PrintError(action, path string, err error) {
	if err != nil {
		fmt.Printf("error %s %s: %v\n", action, path, err)
	}
}
