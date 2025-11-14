package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:  "update [filepath] [content]",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return update(args)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func update(args []string) error {
	//file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0644) //os.O_TRUNC clears the file first — all old content is deleted.
	err := os.WriteFile(args[0], []byte(args[1]), 0644)
	if err != nil {
		return err
	}
	//defer utils.CloseFile(file)
	//_, err = file.WriteString(content)
	// if err != nil {
	// 	return err
	// }

	absPath, err := filepath.Abs(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("%s file is created.\n", absPath)
	return nil
}
