package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/kkumar/file/pkg/utils"
	"github.com/spf13/cobra"
)

var dest string

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy file from source to destination",
	RunE: func(cmd *cobra.Command, args []string) error {
		return copy()
	},
}

var mvCmd = &cobra.Command{
	Use:   "mv",
	Short: "Move file",
	RunE: func(cmd *cobra.Command, args []string) error {
		return move()
	},
}

func init() {
	copyCmd.Flags().StringVarP(&dest, "dest", "d", "", "destination file path")
	if err := copyCmd.MarkFlagRequired("dest"); err != nil {
		panic(err)
	}
	rootCmd.AddCommand(copyCmd)

	mvCmd.Flags().StringVarP(&dest, "dest", "d", "", "destination file path")
	if err := mvCmd.MarkFlagRequired("dest"); err != nil {
		panic(err)
	}
	rootCmd.AddCommand(mvCmd)
}

func copy() error {
	pathFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer utils.CloseFile(pathFile)

	dstFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer utils.CloseFile(dstFile)
	_, err = io.Copy(dstFile, pathFile)
	if err != nil {
		return err
	}

	abspathPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	absDstPath, err := filepath.Abs(dest)
	if err != nil {
		return err
	}

	fmt.Printf("%s file is copied to %s\n", abspathPath, absDstPath)
	return nil
}

func move() error {
	err := os.Rename(path, dest)
	if err != nil {
		return err
	}
	abspathPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	absDstPath, err := filepath.Abs(dest)
	if err != nil {
		return err
	}

	fmt.Printf("%s file is moved to %s\n", abspathPath, absDstPath)
	return nil
}
