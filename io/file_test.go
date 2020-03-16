package io

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFile(t *testing.T) {
	file := NewFile(TestFileName)

	assert.Equal(t, TestFileName, file.Path())
}

func TestFileExists(t *testing.T) {
	file := NewFile(TestFileName)

	assert.False(t, file.Exists())

	file.Write(TestFileContent)
	assert.True(t, file.Exists())
	actual := file.Delete()

	assert.True(t, actual)
}

func TestFileAppend(t *testing.T) {
	file := NewFile(TestFileName)
	file.Write(TestFileContent)
	file.Append(TestFileContent)

	actual := file.ReadAll()

	assert.Equal(t, TestFileContent+TestFileContent, actual)
	file.Delete()
}

func TestFileWrite(t *testing.T) {
	file := NewFile(TestFileName)
	file.Write(TestFileContent)
	actual := file.ReadAll()

	assert.Equal(t, TestFileContent, actual)
	file.Delete()
}

func TestFileDelete(t *testing.T) {
	file := NewFile(TestFileName)
	file.Write(TestFileContent)

	actual := file.Delete()

	assert.True(t, actual)
}
func TestFileDeleteNotExist(t *testing.T) {
	notExistFile := NewFile("not-exists.xx")
	assert.True(t, notExistFile.Delete())
}

func TestFileReadAll(t *testing.T) {
	file := NewFile(TestFileName)
	file.Write(TestFileContent)
	actual := file.ReadAll()

	assert.Equal(t, TestFileContent, actual)
	file.Delete()
}
func TestFileReadAllNotExist(t *testing.T) {
	notExistFile := NewFile("not-exists.xx")
	assert.Empty(t, notExistFile.ReadAll())
}

func TestFileExtension(t *testing.T) {
	file := NewFile(TestFileName)
	assert.Equal(t, ".txt", file.Extension())
}

func TestFileGetName(t *testing.T) {
	file := NewFile(TestFileName)
	assert.Equal(t, TestFileName, file.FileName())
}

func TestFileGetNameWithoutExtension(t *testing.T) {
	file := NewFile(TestFileName)
	assert.Equal(t, "starwar", file.FileNameWithoutExtension())
}

func TestFileGetSize(t *testing.T) {
	file := NewFile(TestFileName)
	file.Write(TestFileContent)

	defer file.Delete()

	assert.True(t, file.Size() > 0)

	notExistFile := NewFile("not-exists.xx")

	assert.Equal(t, int64(-1), notExistFile.Size())
}

func TestFileParent(t *testing.T) {
	file := NewFile(TestFileName)
	parent := file.Parent()

	assert.NotEmpty(t, parent)
}
