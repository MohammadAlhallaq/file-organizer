package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	currentUser, err := user.Current()
	downloadsDir := filepath.Join(currentUser.HomeDir, "Downloads/test")

	files := files{
		extensions: []string{},
		references: []os.FileInfo{},
	}

	dir, err := os.Open(downloadsDir)
	if err != nil {
		return
	}
	defer func(dir *os.File) {
		err := dir.Close()
		if err != nil {
		}
	}(dir)

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	files.collectFiles(fileInfos)
	files.organize(downloadsDir)
}
