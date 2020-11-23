package cmd

import (
	"fmt"
	"os"

	//homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	//"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "semver-cli",
	Short: "Semver is versioning standard that works well with git.",
	Long: `Semver stands for Semantic versioning. 
It is originated by creators of git. 
For more info and guidelines visit https://semver.org`,
	Run: func(cmd *cobra.Command, args []string) {
	  // Do Stuff Here
	},
  }
  
  func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Fprintln(os.Stderr, err)
	  os.Exit(1)
	}
  }