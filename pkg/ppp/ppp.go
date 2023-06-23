package ppp

import (
	"fmt"
	"os/exec"
	"runtime"
)

func Vsam(input string) (result string) {

	var cmd *exec.Cmd

	//Clone Bankdemo
	fmt.Printf("*---\n Cloning BankDemo\n---*\n\n")
	switch runtime.GOOS {
	case "windows":
		//cmd.Dir = "C:\\Users\\PJennings\\git\\ppp"
		var repo = "https://github.com/MicroFocus/BankDemo.git"
		cmd = exec.Command("git", "clone", repo, "--progress")

	default: //Mac & Linux
		//cmd.Dir = "C:\\Users\\PJennings\\git\\ppp"
		var repo = "https://github.com/MicroFocus/BankDemo.git"
		cmd = exec.Command("git", "clone", repo, "--progress")
	}

	if err := cmd.Run(); err != nil {
		return err.Error()
	}

	//Run Python
	fmt.Printf("*---\n Running Python to deploy BankDemo\n---*\n\n")
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("python", "MF_Provision_Region.py", "vsam")

	default: //Mac & Linux
		cmd = exec.Command("python", "MF_Provision_Region.py", "vsam")
	}

	if err := cmd.Run(); err != nil {
		return err.Error()
	}

	return cmd.String()
}
