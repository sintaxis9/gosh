package fs

import (
	"fmt"
	"os"
)

func Touch(args []string) {
	if len(args) != 1 {
		fmt.Println("usage: touch <file_path>")
		return
	}

	path := args[0]
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("error creating file %s: %v\n", path, err)
		return
	}
	defer f.Close()

	fmt.Printf("file created: %s\n", path)
}
