package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type files struct {
	extensions []string
	references []os.FileInfo
}

func (f *files) collectFiles(fileInfos []os.FileInfo) {
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			f.extensions = append(f.extensions, filepath.Ext(fileInfo.Name()))
			f.references = append(f.references, fileInfo)
		}
	}
}

func (f *files) organize(currentDir string) {
	for _, file := range f.references {
		organizeFile(currentDir, file)
	}
}

func organizeFile(currentDir string, file os.FileInfo) {
	sourcePath := filepath.Join(currentDir, file.Name())
	destDir := filepath.Join(currentDir, getDestinationDir(file.Name()))
	destPath := filepath.Join(destDir, file.Name())

	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		if err := createDestinationDir(destDir); err != nil {
			fmt.Println(err)
			return
		}
	}

	if err := moveFile(sourcePath, destPath); err != nil {
		fmt.Println(err)
		return
	}
}

func getDestinationDir(fileName string) string {
	return strings.Replace(filepath.Ext(fileName), ".", "", -1)
}

func createDestinationDir(destDir string) error {
	return os.MkdirAll(destDir, os.ModePerm)
}

func moveFile(sourcePath, destPath string) error {
	return os.Rename(sourcePath, destPath)
}
