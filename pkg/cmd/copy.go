package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/kkumar/file/pkg/utils"
	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:  "copy [srcpath] [dstpath]",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return copy(args)
	},
}

var mvCmd = &cobra.Command{
	Use:  "mv [srcpath] [dstpath]",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return move(args)
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
	rootCmd.AddCommand(mvCmd)
}

func copy(args []string) error {
	srcPath := args[0]
	dstPath := args[1]
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer utils.CloseFile(srcFile)

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer utils.CloseFile(dstFile)
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	absSrcPath, err := filepath.Abs(srcPath)
	if err != nil {
		return err
	}

	absDstPath, err := filepath.Abs(dstPath)
	if err != nil {
		return err
	}

	fmt.Printf("%s file is copied to %s\n", absSrcPath, absDstPath)
	return nil
}

func move(args []string) error {
	srcPath := args[0]
	dstPath := args[1]
	err := os.Rename(srcPath, dstPath)
	if err != nil {
		return err
	}
	absSrcPath, err := filepath.Abs(srcPath)
	if err != nil {
		return err
	}

	absDstPath, err := filepath.Abs(dstPath)
	if err != nil {
		return err
	}

	fmt.Printf("%s file is moved to %s\n", absSrcPath, absDstPath)
	return nil
}
