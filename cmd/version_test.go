package cmd

import (
	"bytes"
	"testing"
	"os"
	"io"

	"github.com/aditi-agarwal-tw/chimp/utility"
)

func TestVersionCommandInitialization(t *testing.T) {
	utility.AssertEqual(t, versionCmd.Use, "version", "Invalid usage of version command")

	utility.AssertEqual(t, versionCmd.Long,
		"Show version of the existing chimp installation.",
		"Invalid long message for version command")

	utility.AssertEqual(t, versionCmd.Short,
		"Show version",
		"Invalid short message for version command")
}

func TestVersionCommandExecution(t *testing.T) {
	var myslice []string
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
  
	versionCmd.Run(versionCmd, myslice)
  
	w.Close()
	os.Stdout = old
  
	var buf bytes.Buffer
	io.Copy(&buf, r)

	utility.AssertEqual(t, buf.String(),
		"Chimp current version: v0.0.1 -- HEAD\n",
		"Invalid version message")
}
