/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var path string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "file",
	Short: "file management",
	Long:  `Create, Read, Update, Delete, Copy, Move and Append a file. For example: file create -f <filename>`,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&path, "filepath", "f", "", "file path (required)")
	if err := rootCmd.MarkPersistentFlagRequired("filepath"); err != nil {
		panic(err)
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
