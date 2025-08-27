package fs

import (
	"fmt"
	"os"
)

func Mkdir(args []string) {
	if len(args) != 1 {
		fmt.Println("usage: mkdir <folder_path>")
		return
	}

	path := args[0]
	err := os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Printf("error creating folder %s: %v\n", path, err)
		return
	}

	fmt.Printf("folder created: %s\n", path)
}
