package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kkumar/file/pkg/utils"
	"github.com/spf13/cobra"
)

var appendCmd = &cobra.Command{
	Use:   "append",
	Short: "Append text at end of file",
	RunE: func(cmd *cobra.Command, args []string) error {
		return append()
	},
}

func init() {
	appendCmd.Flags().StringVarP(&content, "content", "c", "", "Content to append to file")
	if err := appendCmd.MarkFlagRequired("content"); err != nil {
		panic(err)
	}
	rootCmd.AddCommand(appendCmd)
}

func append() error {
	//os.O_CREATE Create the file if it doesn’t exist, os.O_APPEND Always write at the end (no overwrite)
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer utils.CloseFile(file)
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	fmt.Printf("%s file is appended.\n", absPath)
	return nil
}
