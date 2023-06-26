package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/micropj/gocmd/cmd/gocmd"
)

func main() {
	myFigure := figure.NewFigure("GOCMD", "", true)
	myFigure.Print()
	fmt.Printf("\n\n")
	gocmd.Execute()
}
