package repository

import (
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
		forceLocal bool) {
		called = true
		path = pathLocal
		dirName = dirNameLocal
		force = forceLocal
	}

	chimpRepoObj.Initialize(dir)

	utility.AssertEqual(t, called, true, "Chimp Repository creation error")
	utility.AssertEqual(t, path, dir, "Chimp Repository creation error")
	utility.AssertEqual(t, dirName, ".chimp", "Chimp Repository creation error")
	utility.AssertEqual(t, force, false, "Chimp Repository creation error")
}
