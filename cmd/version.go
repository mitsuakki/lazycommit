package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

type Commit struct {
	Hash string
}

var (
	version Version
	commit  Commit
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of lazycommit",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %d.%d.%d\n", version.Major, version.Minor, version.Patch)
		fmt.Printf("Commit: %s\n", commit.Hash)
	},
}
