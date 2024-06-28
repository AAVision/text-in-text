/*
Copyright © 2024 NAME HERE <AAVISION>
*/
package cmd

import (
	"log"
	"os"
	"time"

	"github.com/AAVision/text-in-text/src"
	"github.com/AAVision/text-in-text/utils"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode the hidden text in a normal text using password.",
	Long:  `Encode the hidden text in a normal text using password.`,
	Run: func(cmd *cobra.Command, args []string) {
		coverText, _ := cmd.Flags().GetString("text")
		hiddenText, _ := cmd.Flags().GetString("secret")
		password, _ := cmd.Flags().GetString("password")

		if coverText != "" && hiddenText != "" && password != "" {
			cipherSecret, _ := utils.Encrypt(hiddenText, password)
			encodedText := src.Encode(coverText, []byte(cipherSecret))
			color.Green.Println("Your text was added to file and you can share it with anyone!")
			now := time.Now()
			err := os.WriteFile(now.Format(time.RFC3339)+".txt", []byte(encodedText), 0777)
			if err != nil {
				log.Fatal(err)
			}
			color.Cyan.Println("file written successfully.")
		} else {

		}
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)
	encodeCmd.Flags().String("secret", "", "A secret to be hidden!")
}
