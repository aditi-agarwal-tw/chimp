package utility

import (
	"bytes"
	"os"
	"testing"
)

func TestDirectoryCreate(t *testing.T) {
	path, _ := os.Getwd()
	dirName := ".chimpy"
	CreateDirectory(path, dirName, false)

	var chimpyPath bytes.Buffer
	chimpyPath.WriteString(path)
	chimpyPath.WriteString("/.chimpy")

	file, _ := os.Stat(chimpyPath.String())

	AssertEqual(t, file != nil, true,
		"Error creating root chimp directory")
}
