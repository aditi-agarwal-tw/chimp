package utility

import (
	"bytes"
	"os"
)

// CreateDirectory us used to create dir
var CreateDirectory = func(path string, dirName string, force bool) {
	var newDirPath bytes.Buffer
	newDirPath.WriteString(path)
	newDirPath.WriteString("/")
	newDirPath.WriteString(dirName)

	if _, err := os.Stat(newDirPath.String()); os.IsNotExist(err) {
		os.Mkdir(newDirPath.String(), 0777)
	}
}
