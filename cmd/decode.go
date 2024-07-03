/*
Copyright © 2024 NAME HERE <AAVISION>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AAVision/text-in-text/src"
	"github.com/AAVision/text-in-text/utils"
	"github.com/gookit/color"
	"github.com/jedib0t/go-pretty/v6/table"
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

		now := time.Now()
		data, err := os.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}

		decodedText := src.Decode(string(data))
		output, _ := utils.Decrypt(decodedText, password)

		t := table.NewWriter()
		rowHeader := table.Row{"#", "Path", "Secret", "Finished in:"}

		if output == "" {
			output = color.BgHiRed.Sprintf("Un-resolved secret 🙂  ")
		} else {
			output = color.BgHiGreen.Sprintf(output)
		}

		row := table.Row{"-", path, output, color.BgCyan.Sprintf("%s", time.Since(now))}
		t.AppendRows([]table.Row{row})
		t.AppendHeader(rowHeader)
		fmt.Println(t.Render())

	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)
	decodeCmd.Flags().String(PATH_FLAG, "", "Path of file")
	decodeCmd.MarkFlagRequired(PATH_FLAG)
}
