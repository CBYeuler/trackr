package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "tracker",
	Short: "A simple command line tracker",
	Long:  `A simple command line tracker for managing tasks and projects.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
