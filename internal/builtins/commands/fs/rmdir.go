package fs

import (
	"fmt"
	"os"
)

func Rmdir(args []string) {
	if len(args) != 1 {
		fmt.Println("usage: rmdir <folder_path>")
		return
	}

	path := args[0]
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Printf("error removing folder %s: %v\n", path, err)
		return
	}

	fmt.Printf("folder removed: %s\n", path)
}
