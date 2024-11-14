package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Short:        "AI-powered git hook to generate commit messages from the staged diff.",
	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
