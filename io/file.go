package io

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// File represent a file
type File struct {
	path string
}

// NewFile create a new file
func NewFile(path string) *File {
	return &File{
		path: path,
	}
}

// Path get file path
func (file *File) Path() string {
	return file.path
}

// Parent return parent folder
func (file *File) Parent() string {
	return filepath.Dir(file.path)
}

// Extension get file extension
func (file *File) Extension() string {
	return filepath.Ext(file.path)
}

// FileName get filename with extension
func (file *File) FileName() string {
	_, name := filepath.Split(file.path)

	return name
}

// FileNameWithoutExtension get filename without extension
func (file *File) FileNameWithoutExtension() string {
	_, filename := filepath.Split(file.path)
	extension := file.Extension()
	return strings.TrimSuffix(filename, extension)
}

// Size get file size
func (file *File) Size() int64 {
	info, err := os.Stat(file.path)

	if err != nil {
		return -1
	}

	return info.Size()
}

// Exists return a value whether the file exists
func (file *File) Exists() bool {
	info, err := os.Stat(file.path)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

// ReadAll read all content to text
func (file *File) ReadAll() string {
	raw, err := os.Open(file.path)
	if err != nil {
		return ""
	}
	defer raw.Close()

	bytes, err := ioutil.ReadAll(raw)
	return string(bytes)
}

// Append append content to file
func (file *File) Append(content string) error {
	raw, err := os.OpenFile(file.path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer raw.Close()

	_, err = raw.WriteString(content)
	return err
}

// Write write content to file
func (file *File) Write(content string) error {
	raw, err := os.OpenFile(file.path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer raw.Close()

	_, err = raw.WriteString(content)
	return err
}

// Delete delete give file
func (file *File) Delete() bool {
	if !file.Exists() {
		return true
	}

	err := os.Remove(file.path)

	return err == nil
}
