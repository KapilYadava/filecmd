package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kkumar/file/pkg/utils"
	"github.com/spf13/cobra"
)

var appendCmd = &cobra.Command{
	Use:   "append [filepath] [content]",
	Short: "Append text at end of file",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return append(args)
	},
}

func init() {
	rootCmd.AddCommand(appendCmd)
}

func append(args []string) error {
	//os.O_CREATE Create the file if it doesn’t exist, os.O_APPEND Always write at the end (no overwrite)
	file, err := os.OpenFile(args[0], os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer utils.CloseFile(file)
	_, err = file.WriteString(args[1])
	if err != nil {
		return err
	}

	absPath, err := filepath.Abs(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("%s file is appended.\n", absPath)
	return nil
}
