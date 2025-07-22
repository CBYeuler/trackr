package cmd

import (
	"github.com/CBYeuler/trackr/db"
	"github.com/spf13/cobra"
	//github.com/CBYeuler/trackr
)

var addCmd = &cobra.Command{
	Use:   "add <category> <value>",
	Short: "Add a new log entry to track something",
	Long:  `Adds a new log entry with a category (like mood, gym, prayer) and value.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Println("Usage: trackr add <category> <value>")
			return
		}

		category := args[0]
		value := args[1]

		dbConn, err := db.OpenDB()
		if err != nil {
			cmd.PrintErrf("Failed to open DB: %v\n", err)
			return
		}
		defer dbConn.Close()

		if err := db.InsertLog(dbConn, category, value); err != nil {
			cmd.PrintErrf("Failed to insert log: %v\n", err)
			return
		}

		cmd.Printf("Log added: [%s] %s\n", category, value)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

// AddCmd adds the add command to the root command.
// It allows users to add new tasks or projects to the tracker.
