/*
Copyright © 2024 NAME HERE <AAVISION>
*/
package cmd

import (
	"log"
	"os"

	"github.com/AAVision/text-in-text/src"
	"github.com/AAVision/text-in-text/utils"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode the hidden text from the encoded text using password.",
	Long:  `Decode the hidden text from the encoded text using password.`,
	Run: func(cmd *cobra.Command, args []string) {

		path, _ := cmd.Flags().GetString("path")
		password, _ := cmd.Flags().GetString("password")
		if path != "" && password != "" {

			data, err := os.ReadFile(path)
			if err != nil {
				log.Fatal(err)
			}

			decodedText := src.Decode(string(data))
			output, _ := utils.Decrypt(decodedText, password)
			if output == "" {
				color.Red.Println("Un-resolved secret 🙂")
			} else {
				color.Green.Print("Secret: ")
				color.Cyan.Println(output)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)
	decodeCmd.Flags().String("path", "", "Path of file")
}
