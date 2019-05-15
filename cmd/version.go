package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long: `Show version of the existing chimp installation.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(os.Stdout, "Chimp current version: v0.0.1 -- HEAD")
	},
  }