package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Hello, World!")
	Command := "ls"
	cmd_obj := exec.Command(Command)
	cmd_obj.Stdout = os.Stdout
	cmd_obj.Stderr = os.Stderr
	err := cmd_obj.Run()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}
