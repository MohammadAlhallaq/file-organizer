package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {

	directories := getDirectories()

	for _, directory := range directories {
		files := files{
			extensions: []string{},
			references: []os.FileInfo{},
		}

		dir, err := os.Open(directory)
		if err != nil {
			return
		}

		fileInfos, err := dir.Readdir(-1)
		if err != nil {
			fmt.Println("Error reading directory:", err)
			return
		}
		err = dir.Close()

		files.collectFiles(fileInfos)
		files.organize(directory)
	}
}

func getDirectories() []string {
	currentUser, _ := user.Current()
	downloadsDir := filepath.Join(currentUser.HomeDir, "Downloads")
	desktopDir := filepath.Join(currentUser.HomeDir, "Desktop")

	return []string{downloadsDir, desktopDir}
}
