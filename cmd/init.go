package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/aditi-agarwal-tw/chimp/repository"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a chimp project",
	Long: `Initializes a chimp project for version control in your root directory.
Will setup the basic chimp configuration and setup the algorithms`,
	Run: executeInit,
}

// ChimpRepoObj is the singleton for chimp repo
var ChimpRepoObj repository.ChimpRepository

// ChimpRepInitializeFn to make the function an interface
var ChimpRepInitializeFn = ChimpRepoObj.Initialize

func executeInit(cmd *cobra.Command, args []string) {
	dir, _ := os.Getwd()
	fmt.Println("Initializing chimp repo at: ", dir)
	ChimpRepInitializeFn(dir)
}
