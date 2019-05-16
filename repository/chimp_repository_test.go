package repository

import (
	"errors"
	"os"
	"testing"

	"github.com/aditi-agarwal-tw/chimp/utility"
)

var chimpRepoObj = ChimpRepository{}

func TestChimpRepositoryInitialize(t *testing.T) {
	dir, _ := os.Getwd()
	called := false
	path := ""
	dirName := ""
	force := true

	utility.CreateDirectory = func(
		pathLocal string,
		dirNameLocal string,
		forceLocal bool) error {
		called = true
		path = pathLocal
		dirName = dirNameLocal
		force = forceLocal
		return nil
	}

	chimpRepoObj.Initialize(dir)

	utility.AssertEqual(t, called, true, "Chimp Repository creation error")
	utility.AssertEqual(t, path, dir, "Chimp Repository creation error")
	utility.AssertEqual(t, dirName, ".chimp", "Chimp Repository creation error")
	utility.AssertEqual(t, force, false, "Chimp Repository creation error")
}

func TestChimpRepositoryInitializeOnError(t *testing.T) {
	utility.CreateDirectory = func(
		pathLocal string,
		dirNameLocal string,
		forceLocal bool) error {
		return errors.New("some-error-occurred")
	}

	err := chimpRepoObj.Initialize("some-dir")

	utility.AssertEqual(t, err != nil, true, "Chimp Repository creation error")

}
