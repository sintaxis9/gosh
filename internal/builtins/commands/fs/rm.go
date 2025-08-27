package fs

import (
	"fmt"
	"os"
)

func Rm(args []string) {
	if len(args) != 1 {
		fmt.Println("usage: rm <file_path>")
		return
	}

	path := args[0]
	info, err := os.Stat(path)
	if err != nil {
		fmt.Printf("error accesing file %s: %v\n", path, err)
		return
	}

	if info.IsDir() {
		fmt.Printf("%s is a dir, instead use rmdir to remove it!\n", path)
		return
	}

	err = os.Remove(path)
	if err != nil {
		fmt.Printf("error removing file %s: %v\n", path, err)
		return
	}

	fmt.Printf("file removed: %s\n", path)
}
