package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kkumar/file/pkg/utils"
	"github.com/spf13/cobra"
)

var content string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create file",
	RunE: func(cmd *cobra.Command, args []string) error {
		return create()
	},
}

func init() {
	createCmd.Flags().StringVarP(&content, "content", "c", "This is my file conent", "file content (required)")
	if err := createCmd.MarkFlagRequired("content"); err != nil {
		panic(err)
	}
	rootCmd.AddCommand(createCmd)
}

func create() error {
	_, err := os.Stat(path)
	if err == nil {
		return fmt.Errorf("file %s already exist", path)
	} else if os.IsNotExist(err) {
		file, err := os.Create(path)
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
		fmt.Printf("%s file is created.\n", absPath)
		return nil
	}
	return err
}
