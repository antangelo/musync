package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "musync",
		Short: "A tool for synchronizing music between streaming services",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
