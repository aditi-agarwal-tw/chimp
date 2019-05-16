package utility

import (
	"bytes"
	"os"
	"testing"
)

func TestDirectoryCreate(t *testing.T) {
	os.RemoveAll(".chimpy")

	path, _ := os.Getwd()
	dirName := ".chimpy"
	CreateDirectory(path, dirName, false)

	var chimpyPath bytes.Buffer
	chimpyPath.WriteString(path)
	chimpyPath.WriteString("/.chimpy")

	file, _ := os.Stat(chimpyPath.String())

	AssertEqual(t, file != nil, true,
		"Error creating root chimp directory")

	os.RemoveAll(".chimpy")
}

func TestDirectoryCreateTrailingPath(t *testing.T) {
	os.RemoveAll("some-dir")
	path := "./some-dir/"
	dirName := ".chimp"

	os.Mkdir("some-dir", 0777)

	CreateDirectory(path, dirName, false)

	file, _ := os.Stat("./some-dir/.chimp")

	AssertEqual(t, file != nil, true,
		"Error creating root chimp directory")

	os.RemoveAll("some-dir")
}

func TestDirectoryCreateOnError(t *testing.T) {
	path := "./some-dir/some-project"
	dirName := ".chimp"

	var err = CreateDirectory(path, dirName, false)

	AssertEqual(t, err != nil, true,
		"Error creating root chimp directory")
}
