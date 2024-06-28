/*
Copyright © 2024 NAME HERE <AAVISION>
*/
package main

import (
	"fmt"
	"time"

	"github.com/AAVision/text-in-text/cmd"
	"github.com/common-nighthawk/go-figure"
	"github.com/gookit/color"
)

func main() {
	myFigure := figure.NewFigure("Text-To-Text", "larry3d", true)
	myFigure.Print()
	fmt.Println()
	now := time.Now()
	cmd.Execute()
	fmt.Print("Finished in: ")
	color.BgHiGreen.Println(time.Since(now))
}
