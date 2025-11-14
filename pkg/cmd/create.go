package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kkumar/file/pkg/utils"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:  "create [filepath] [content]",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return create(args)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func create(args []string) error {
	_, err := os.Stat(args[0])
	if err == nil {
		return fmt.Errorf("file %s already exist", args[0])
	} else if os.IsNotExist(err) {
		file, err := os.Create(args[0])
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
		fmt.Printf("%s file is created.\n", absPath)
		return nil
	}
	return err
}
