/*
Copyright © 2024 NAME HERE <AAVISION>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "text-in-text",
	Short: "Hide text inside a text.",
	Long: `This tool was created to hide text inside a text using text-plain steganography. 
The hidden secret text is protected by a password and using AES symmetric key.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var (
	// Global flag variable
	text     string
	password string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.text-in-text.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringVar(&text, "text", "", "An encoded text to extract secret from it!")
	rootCmd.PersistentFlags().StringVar(&password, "password", "", "A password to protect your text!")

}
