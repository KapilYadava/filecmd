package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [filepath]",
	Short: "Delete file",
	RunE: func(cmd *cobra.Command, args []string) error {
		return delete(args)
	},
}

var deleteAllCmd = &cobra.Command{
	Use:   "deleteAll [filepath]",
	Short: "Delete all files or dirs",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteAll(args)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(deleteAllCmd)
}

func delete(args []string) error {
	for _, filepath := range args {
		// only delete empty folder
		err := os.Remove(filepath)
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteAll(args []string) error {
	for _, filepath := range args {
		// delete the folder and everything inside it
		err := os.RemoveAll(filepath)
		if err != nil {
			return err
		}
	}
	return nil
}
