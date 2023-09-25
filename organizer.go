package main

import (
	"os"
	"path/filepath"
)

type files struct {
	extensions []string
	references []*os.FileInfo
}

func (f *files) collectFiles(fileInfos []os.FileInfo) {
	for _, fileInfo := range fileInfos {
		f.extensions = append(f.extensions, filepath.Ext(fileInfo.Name()))
		f.references = append(f.references, &fileInfo)
	}
}

func (f *files) organize() {

}
