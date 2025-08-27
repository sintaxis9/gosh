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
