package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(incrementCmd)
}

var incrementCmd = &cobra.Command{
	Use:   "increment",
	Short: "Increment version",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}