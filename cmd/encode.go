/*
Copyright © 2024 NAME HERE <AAVISION>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/AAVision/text-in-text/src"
	"github.com/AAVision/text-in-text/utils"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	text     string
	password string
	secret   string
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

		now := time.Now()
		cipherSecret, _ := utils.Encrypt(hiddenText, password)
		encodedText := src.Encode(coverText, []byte(cipherSecret))
		color.Green.Println("Your text was added to file and you can share it with anyone!")
		err := os.WriteFile(strconv.Itoa(int(time.Now().Unix()))+".txt", []byte(encodedText), 0777)
		if err != nil {
			log.Fatal(err)
		}
		color.Cyan.Println("File written successfully!")
		fmt.Print("Finished in: ")
		color.BgHiGreen.Println(time.Since(now))

	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)

	encodeCmd.PersistentFlags().StringVar(&text, "text", "", "An encoded text to extract secret from it!")
	encodeCmd.PersistentFlags().StringVar(&password, "password", "", "A password to protect your text!")
	encodeCmd.PersistentFlags().StringVar(&secret, "secret", "", "A secret to be hidden!")

	encodeCmd.MarkPersistentFlagRequired("text")
	encodeCmd.MarkPersistentFlagRequired("password")
	encodeCmd.MarkPersistentFlagRequired("secret")
}
