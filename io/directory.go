package io

import (
	"os"
	"path/filepath"
)

// Directory represent a local folder/directory
type Directory struct {
	path string
}

// NewDirectory new instance of Directory
func NewDirectory(fullPath string) *Directory {
	return &Directory{
		path: fullPath,
	}
}

// Path get full path
func (dir *Directory) Path() string {
	return dir.path
}

// Exists return a value whether the folder exists, files in subdirectories are not included
func (dir *Directory) Exists() bool {
	_, err := os.Stat(dir.path)
	if err == nil {
		return true
	}

	return !os.IsNotExist(err)
}

// Create the folder if it doesn't exist
func (dir *Directory) Create() bool {
	if dir.Exists() {
		return true
	}

	err := os.MkdirAll(dir.path, os.ModePerm)
	return err == nil
}

// Delete the folder
func (dir *Directory) Delete() bool {
	err := os.RemoveAll(dir.path)
	return err == nil
}

// GetAllFiles Get files in this directory inclue files in subdirectories
func (dir *Directory) GetAllFiles() []*File {
	return dir.getFilesInternal(true)
}

// GetFiles get files without sub directories
func (dir *Directory) GetFiles() []*File {
	return dir.getFilesInternal(false)
}

// GetSubDirectories get sub directories in this level
func (dir *Directory) GetSubDirectories() []*Directory {
	return dir.getDirectoriesInternal(false)
}

// GetAllSubDirectories get sub directories include subdirectories
func (dir *Directory) GetAllSubDirectories() []*Directory {
	return dir.getDirectoriesInternal(true)
}

// Join compose the full path under current directory
func (dir *Directory) Join(parts ...string) string {
	return filepath.Join(dir.path, filepath.Join(parts...))
}

// getFilesInternal get files
func (dir *Directory) getFilesInternal(includeSub bool) []*File {
	files := make([]*File, 0)

	if !dir.Exists() {
		return files
	}

	folder, err := os.Open(dir.path)
	if err != nil {
		return files
	}
	defer folder.Close()

	names, _ := folder.Readdirnames(-1)

	for _, name := range names {
		fullPath := filepath.Join(dir.path, name)
		info, _ := os.Stat(fullPath)
		if info.IsDir() {
			if includeSub {
				subDir := NewDirectory(fullPath)
				subFiles := subDir.getFilesInternal(includeSub)
				files = append(files, subFiles...)
			}

			continue
		}

		files = append(files, NewFile(fullPath))
	}
	return files
}

// getDirectoriesInternal get sub directories
func (dir *Directory) getDirectoriesInternal(includeSub bool) []*Directory {
	dirs := make([]*Directory, 0)
	if !dir.Exists() {
		return dirs
	}

	folder, err := os.Open(dir.path)
	if err != nil {
		return dirs
	}
	defer folder.Close()

	names, _ := folder.Readdirnames(-1)

	for _, name := range names {
		fullPath := filepath.Join(dir.path, name)
		info, _ := os.Stat(fullPath)
		if !info.IsDir() {
			continue
		}

		if includeSub {
			subDir := NewDirectory(fullPath)
			subDirs := subDir.getDirectoriesInternal(includeSub)
			dirs = append(dirs, subDirs...)
		}

		dirs = append(dirs, NewDirectory(fullPath))
	}

	return dirs
}
