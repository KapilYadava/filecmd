package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update file",
	RunE: func(cmd *cobra.Command, args []string) error {
		return update()
	},
}

func init() {
	updateCmd.Flags().StringVarP(&content, "content", "c", "This is my file conent", "file content (required)")
	if err := updateCmd.MarkFlagRequired("content"); err != nil {
		panic(err)
	}
	rootCmd.AddCommand(updateCmd)
}

func update() error {
	//file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0644) //os.O_TRUNC clears the file first — all old content is deleted.
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		return err
	}
	//defer utils.CloseFile(file)
	//_, err = file.WriteString(content)
	// if err != nil {
	// 	return err
	// }

	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	fmt.Printf("%s file is updated.\n", absPath)
	return nil
}
