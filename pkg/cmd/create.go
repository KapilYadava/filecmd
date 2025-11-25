package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kkumar/file/pkg/utils"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create file",
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := cmd.Flags().GetString("filepath")
		if err != nil {
			return err
		}
		content, err := cmd.Flags().GetString("content")
		if err != nil {
			return err
		}
		return create(path, content)
	},
}

func init() {
	createCmd.Flags().StringP("content", "c", "This is my file conent", "file content (required)")
	if err := createCmd.MarkFlagRequired("content"); err != nil {
		panic(err)
	}

	rootCmd.PersistentFlags().StringP("filepath", "f", "", "file path (required)")
	if err := rootCmd.MarkPersistentFlagRequired("filepath"); err != nil {
		panic(err)
	}
	rootCmd.AddCommand(createCmd)
}

func create(path, content string) error {
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
