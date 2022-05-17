package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(readCmd)
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a configuration file.",
	Long:  `Read a configuration file. To use "tb read PATH"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("file %s\n", args)
	},
}
