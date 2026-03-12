package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// change args to flags using cobra cmd here for read command

var readCmd = &cobra.Command{
	Use:   "read [filepath]",
	Short: "Read file",
	RunE: func(cmd *cobra.Command, args []string) error {
		return read()
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}

func read() error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	fmt.Println(string(content))
	return nil
}
