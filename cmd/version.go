package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func GetTag() string {
	out, err := exec.Command("git", "describe", "--tags").Output()
	if err != nil {
		return "unknown"
	}
	tag := strings.TrimSpace(string(out))
	return tag
}

func GetCommitHash() string {
	out, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
	if err != nil {
		return "unknown"
	}
	commitHash := strings.TrimSpace(string(out))
	return commitHash
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of lazycommit",
	Run: func(cmd *cobra.Command, args []string) {
		tag := GetTag()
		commitHash := GetCommitHash()
		fmt.Printf("lazycommit v%s (%s)\n", tag, commitHash)
	},
}
