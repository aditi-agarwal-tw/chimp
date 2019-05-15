package cmd

import (
	"testing"
	"errors"
	"github.com/spf13/cobra"

	"github.com/aditi-agarwal-tw/chimp/utility"
)

const NumberOfSubCommands int = 2

func TestRootCommandInitialization(t *testing.T) {
	utility.AssertEqual(t, rootCmd.Use, "chimp", "Invalid usage of version command")

	utility.AssertEqual(t, rootCmd.Short,
		"Hello I am Chimp, new version control cool kid on the block...",
		"Invalid long message for root command")

	utility.AssertEqual(t, rootCmd.Long,
		`Hello I am Chimp, new version control cool kid on the block...
I can help you manage your files locally. 
If you have heard of Git, 
I am a its brother written in GoLang`,
		"Invalid short message for root command")
}

func TestRootCommandInitFunc(t *testing.T) {
	commands := rootCmd.Commands()
	commandsSplice := make([]interface{}, len(commands))
	for i, v := range commands {
		commandsSplice[i] = v
	}
	utility.AssertContains(t, commandsSplice, versionCmd, "Version command not added")
}

func TestRootCommandExecution(t *testing.T) {
	addSubCommandsOnSuccess(t)
	exitOnFailure(t)
}

func addSubCommandsOnSuccess(t *testing.T) {
	Execute()
	utility.AssertEqual(t, len(rootCmd.Commands()), NumberOfSubCommands,
		"Invalid execution of root command")
}

func exitOnFailure(t *testing.T) {
	mockRunE := func(cmd *cobra.Command, args []string) error {
		return errors.New("Induced error")
	}
	rootCmd.RunE = mockRunE
	assertExitCode(t, 1, "No exit code returned")
}

func assertExitCode(t *testing.T,
	expectedExitCode int,
	errorMessage string) {
	oldOsExit := OsExit
	defer func() { OsExit = oldOsExit }()

	var actual int
	mockExit := func(code int) {
		actual = code
	}

	OsExit = mockExit
	Execute()

	if actual != expectedExitCode {
		t.Errorf("Test failed, got: %d, want: %d, message: %s.",
			actual, expectedExitCode, errorMessage)
	}
}