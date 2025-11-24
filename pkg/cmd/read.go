package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read [filepath]",
	Short: "Read file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return read(args)
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}

func read(args []string) error {
	filepath := args[0]
	content, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	fmt.Println(string(content))
	return nil
}
