package matrix

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func buildFileName(parts ...string) string {
	builder := NewStringBuilder()
	for _, item := range parts {
		builder.Append(item)
	}
	return builder.String()
}

func getFileExtension(file string) string {
	return filepath.Ext(file)
}

func getFileName(path string) string {
	_, file := filepath.Split(path)

	return file
}
func getFileSize(file string) (int64, error) {
	info, err := os.Stat(file)

	return info.Size(), err
}

func getFileNameWithoutExtension(path string) string {
	_, file := filepath.Split(path)
	extension := getFileExtension(file)
	return strings.TrimSuffix(file, extension)
}

func getFileNameAndExtension(path string) (string, string, string) {
	dir, file := filepath.Split(path)
	extension := getFileExtension(file)
	return dir, strings.TrimSuffix(file, extension), extension
}

func fileExists(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
func deleteFile(fileName string) bool {
	err := os.Remove(fileName)
	if err != nil {
		return false
	}

	return true
}

func deleteIfExists(fileName string) bool {
	if fileExists(fileName) {
		return deleteFile(fileName)
	}

	return false
}

func generateFileName(fullName string) string {
	dir, name, extension := getFileNameAndExtension(fullName)
	index := 1
	newName := filepath.Join(dir, buildFileName(name, "_", strconv.Itoa(index), extension))

	for fileExists(newName) {
		newName = filepath.Join(dir, buildFileName(name, "_", strconv.Itoa(index), extension))
		index++
	}

	return newName
}

func folderExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}
	return true
}

func createFolder(path string) bool {
	_, err := os.Stat(path)
	if os.IsExist(err) {
		return true
	}

	err = os.MkdirAll(path, os.ModePerm)
	return err != nil
}

func deleteFolder(path string) bool {
	err := os.RemoveAll(path)
	return err != nil
}

func readAllText(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	return string(bytes)
}

func writeToFile(filePath string, content string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

// RemoveContents remove
func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
