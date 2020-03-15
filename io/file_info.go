package io

import (	
	"os"
	"path/filepath"
	"strings"
)

// FileInfo Provides properties and instance methods for the creation, copying, deletion, moving, and opening of file
type FileInfo struct {
	fileName string
}

// NewFileInfo new FileInfo instance
func NewFileInfo(fileName string) *FileInfo {
	return &FileInfo{
		fileName: fileName,
	}
}

// FileName get file name
func (info *FileInfo) FileName() string {
	return info.fileName
}

// Extension get file extension
func (info *FileInfo) Extension() string {
	return filepath.Ext(info.fileName)
}

// Name get file name without directory/path
func (info *FileInfo) Name() string {
	_, file := filepath.Split(info.fileName)

	return file
}

// NameWithoutExtensions get file name without extension
func (info *FileInfo) NameWithoutExtensions() string {
	_, file := filepath.Split(info.fileName)
	extension := filepath.Ext(info.fileName)
	return strings.TrimSuffix(file, extension)
}

//DirNameExtension get directory, file name without extension and extension
func (info *FileInfo) DirNameExtension() (string, string, string) {
	dir, file := filepath.Split(info.fileName)
	extension := filepath.Ext(info.fileName)
	return dir, strings.TrimSuffix(file, extension), extension
}

// Directory get parent directory of the file
func (info *FileInfo) Directory() string {
	dir, _ := filepath.Split(info.fileName)
	return dir
}

// Length get file length
func (info *FileInfo) Length() int64 {
	info, err := os.Stat(file)

	if err != nil {
		return -1
	}

	return info.Size()
}

// Exists Gets a value indicating whether a file exists.
func (info *FileInfo) Exists() bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
