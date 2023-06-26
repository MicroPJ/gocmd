package gocmd

import (
	"fmt"
	"runtime"
)

func One(input string) (result string) {

	fmt.Printf("*---[One] Running\n")
	switch runtime.GOOS {
	case "windows":
		fmt.Printf("*---[One] Windows identified\n")

	default: //Mac & Linux
		fmt.Printf("*---[One] Linux identified\n")
	}

	fmt.Printf("*---[One] Completed\n")
	return "Done"
}
