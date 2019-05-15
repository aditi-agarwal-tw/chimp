package cmd

import (
	"testing"
	"os"

	"github.com/aditi-agarwal-tw/chimp/utility"
)

func TestInitCommandDeclaration(t *testing.T) {
	utility.AssertEqual(t, initCmd.Use, "init", "Invalid usage of init command")

	utility.AssertEqual(t, initCmd.Long,
		`Initializes a chimp project for version control in your root directory.
Will setup the basic chimp configuration and setup the algorithms`,
		"Invalid long message for init command")

	utility.AssertEqual(t, initCmd.Short,
		"Initializes a chimp project",
		"Invalid short message for init command")
}

func TestInitCommandExecution(t *testing.T) {
	argsSplice := []string{}
	called := false
	path := ""
	ChimpRepInitializeFn = func(path_local string) {
		called = true
		path = path_local
	}

	initCmd.Run(initCmd, argsSplice)
	dir, _ := os.Getwd()

	utility.AssertEqual(t, called, true, "Chimp Repository not initialized")
	utility.AssertEqual(t, path, dir, "Chimp Repository not initialized")
}
