/*
Copyright © 2024 NAME HERE <AAVISION>
*/
package main

import (
	"fmt"

	"github.com/AAVision/text-in-text/cmd"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	myFigure := figure.NewFigure("Text-To-Text", "larry3d", true)
	myFigure.Print()
	fmt.Println()
	cmd.Execute()
}
