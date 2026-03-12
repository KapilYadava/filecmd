package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete file",
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := cmd.Flags().GetString("filepath")
		if err != nil {
			return err
		}
		return delete([]string{path})
	},
}

var deleteAllCmd = &cobra.Command{
	Use:   "deleteAll",
	Short: "Delete all files or dirs",
	RunE: func(cmd *cobra.Command, args []string) error {
		paths, err := cmd.Flags().GetStringArray("filepath")
		if err != nil {
			return err
		}
		return deleteAll(paths)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(deleteAllCmd)
}

func delete(paths []string) error {
	for _, filepath := range paths {
		// only delete empty folder
		err := os.Remove(filepath)
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteAll(paths []string) error {
	for _, filepath := range paths {
		// delete the folder and everything inside it
		err := os.RemoveAll(filepath)
		if err != nil {
			return err
		}
	}
	return nil
}
