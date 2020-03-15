package io

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// create a fake dir with files and sub-directories
func initFakeDirectory() *Directory {
	dir := NewDirectory(TestDir)
	dir.Create()

	file := NewFile(dir.Join(TestFileName))
	file.Write(TestFileContent)

	file1 := NewFile(dir.Join("mars.txt"))
	file1.Write(TestFileContent)

	subDir := NewDirectory(dir.Join("sub"))
	subDir.Create()

	file2 := NewFile(dir.Join("sub", "mars.txt"))
	file2.Write(TestFileContent)

	subdir1 := NewDirectory(subDir.Join("sub1"))
	subdir1.Create()

	file3 := NewFile(subdir1.Join("earth.txt"))
	file3.Write(TestFileContent)

	return dir
}

func cleanFakeDirectory(dir *Directory) {
	if dir.Exists() {
		dir.Delete()
	}
}

func TestNewDirectory(t *testing.T) {
	dir := NewDirectory("")

	assert.NotNil(t, dir)
	assert.Equal(t, "", dir.Path())

	dir = NewDirectory("test")
	assert.Equal(t, "test", dir.Path())
}

func TestDirectoryCreateDelete(t *testing.T) {
	dir := NewDirectory("./test")

	assert.False(t, dir.Exists())
	assert.True(t, dir.Create())
	assert.True(t, dir.Create())
	assert.True(t, dir.Exists())
	assert.True(t, dir.Delete())
	assert.False(t, dir.Exists())
}

func TestDirectoryGetFiles(t *testing.T) {
	dir := initFakeDirectory()
	defer cleanFakeDirectory(dir)

	actual := dir.GetFiles()

	assert.Equal(t, 2, len(actual))

	assert.Equal(t, filepath.Join(dir.Path(), "mars.txt"), actual[0].Path())
	assert.Equal(t, filepath.Join(dir.Path(), TestFileName), actual[1].Path())
}
func TestDirectoryGetFilesFolderNotExists(t *testing.T) {
	dir := NewDirectory("./not-exists")
	files := dir.GetFiles()

	assert.Empty(t, files)
}

func TestDirectoryGetFilesIncludeSubs(t *testing.T) {
	dir := initFakeDirectory()
	defer cleanFakeDirectory(dir)

	actual := dir.GetAllFiles()

	assert.Equal(t, 4, len(actual))

	assert.Equal(t, filepath.Join(dir.Path(), "mars.txt"), actual[0].Path())
	assert.Equal(t, filepath.Join(dir.Path(), TestFileName), actual[1].Path())
	assert.Equal(t, filepath.Join(dir.Path(), "sub", "mars.txt"), actual[2].Path())
	assert.Equal(t, filepath.Join(dir.Path(), "sub", "sub1", "earth.txt"), actual[3].Path())
}

func TestDirectoryGetDirectories(t *testing.T) {
	dir := initFakeDirectory()
	defer cleanFakeDirectory(dir)

	actual := dir.GetSubDirectories()

	assert.Equal(t, 1, len(actual))
	assert.Equal(t, filepath.Join(TestDir, "sub"), actual[0].Path())
}

func TestDirectoryGetDirectoriesFolderNotExists(t *testing.T) {
	dir := NewDirectory("./not-exists")
	files := dir.GetSubDirectories()

	assert.Empty(t, files)
}
func TestDirectoryGetDirectoriesIncludeSubs(t *testing.T) {
	dir := initFakeDirectory()
	defer cleanFakeDirectory(dir)

	actual := dir.GetAllSubDirectories()

	assert.Equal(t, 2, len(actual))
	assert.Equal(t, filepath.Join(TestDir, "sub"), actual[1].Path())
	assert.Equal(t, filepath.Join(TestDir, "sub", "sub1"), actual[0].Path())
}

func TestDirectoryJoinPath(t *testing.T) {
	dir := NewDirectory(TestDir)

	actual := dir.Join("sub")
	assert.Equal(t, filepath.Join(TestDir, "sub"), actual)
}
