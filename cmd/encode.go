/*
Copyright © 2024 NAME HERE <AAVISION>
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"strconv"
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
		password, err := cmd.Flags().GetString(PASSWORD_FLAG)
		fatalOnErr(err)
		hiddenText, err := cmd.Flags().GetString(SECRET_FLAG)
		fatalOnErr(err)

		if password == "" || hiddenText == "" {
			color.Red.Printf("%q and %q flags cannot be empty for %q sub-command!\n", SECRET_FLAG, PASSWORD_FLAG, ENCODE_COMMAND)
			os.Exit(1)
		}
		var text string
		var r io.Reader
		var providedInputFlags uint

		coverText, err := cmd.Flags().GetString(TEXT_FLAG)
		fatalOnErr(err)
		if coverText != "" {
			providedInputFlags++
		}

		stdin, err := cmd.Flags().GetBool(STDIN_FLAG)
		fatalOnErr(err)
		if stdin {
			providedInputFlags++
		}

		file, err := cmd.Flags().GetString(FILE_FLAG)
		fatalOnErr(err)
		if file != "" {
			providedInputFlags++
		}

		if providedInputFlags != 1 {
			color.Red.Printf("one input flag is required, please provide just one of: %q, %q, or %q\n", TEXT_FLAG, STDIN_FLAG, FILE_FLAG)
			os.Exit(1)
		}

		var needFileRead bool
		if coverText != "" {
			text = coverText
		} else if stdin {
			r = os.Stdin
			needFileRead = true
		} else if file != "" {
			f, err := os.Open(file)
			if err != nil {
				color.Red.Printf("error opening file %q: %s\n", file, err.Error())
				os.Exit(1)
			}
			defer f.Close()

			r = f
			needFileRead = true
		}

		if needFileRead {
			data, err := io.ReadAll(r)
			fatalOnErr(err)
			text = string(data)
		}

		now := time.Now()
		cipherSecret, _ := utils.Encrypt(hiddenText, password)
		encodedText := src.Encode(text, []byte(cipherSecret))
		color.Green.Println("Your text was added to file and you can share it with anyone!")
		err = os.WriteFile(strconv.Itoa(int(now.Unix()))+".txt", []byte(encodedText), 0777)
		fatalOnErr(err)
		color.Cyan.Println("File written successfully!")
		fmt.Print("Finished in: ")
		color.BgHiGreen.Println(time.Since(now))
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)
	encodeCmd.Flags().String(SECRET_FLAG, "", "A secret to be hidden!")

	encodeCmd.Flags().String(FILE_FLAG, "", "Path to file containing the input text")
	encodeCmd.Flags().Bool(STDIN_FLAG, false, "Input text from stdin")
	encodeCmd.Flags().String(TEXT_FLAG, "", "An encoded text to extract secret from it!")
}
