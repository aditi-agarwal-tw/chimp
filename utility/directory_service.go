package utility

import (
	"bytes"
	"os"
)

// CreateDirectory us used to create dir
var CreateDirectory = func(path string, dirName string, force bool) error {
	var newDirPath bytes.Buffer
	newDirPath.WriteString(path)
	newDirPath.WriteString("/")
	newDirPath.WriteString(dirName)

	var err error
	if _, err = os.Stat(newDirPath.String()); os.IsNotExist(err) {
		err = os.Mkdir(newDirPath.String(), 0777)
	}

	return err
}
