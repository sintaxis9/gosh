package utils

import (
	"os"
	"runtime"
	"syscall"
)

func IsHidden(entry os.DirEntry) bool {
	name := entry.Name()

	if runtime.GOOS != "windows" {
		return len(name) > 0 && name[0] == '.'
	}

	info, err := entry.Info()
	if err != nil {
		return false
	}
	attrs := info.Sys().(*syscall.Win32FileAttributeData)
	return attrs.FileAttributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0
}
