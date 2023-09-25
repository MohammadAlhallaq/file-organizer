package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	currentUser, err := user.Current()
	downloadsDir := filepath.Join(currentUser.HomeDir, "Downloads")

	extensions := fileExt{}
	files := files{}

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

	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name())
		extensions = append(extensions, filepath.Ext(fileInfo.Name()))
		files = append(files, &fileInfo)
	}
}
