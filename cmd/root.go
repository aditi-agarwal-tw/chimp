package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var OsExit = os.Exit

var rootCmd = &cobra.Command{
	Use:   "chimp",
	Short: "Hello I am Chimp, new version control cool kid on the block...",
	Long: `Hello I am Chimp, new version control cool kid on the block...
I can help you manage your files locally. 
If you have heard of Git, 
I am a its brother written in GoLang`,
}

// Execute is used to run the rootcmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		OsExit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
