package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func executeCommand(command string,args_arr []string ) (err error){
	args := args_arr
	cmd_obj := exec.Command(command,args...)
	cmd_obj.Stdout = os.Stdout
	cmd_obj.Stderr = os.Stderr
	err = cmd_obj.Run()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		
	}
	return nil
}

func main() {
	command := "ls"  
	executeCommand(command, []string{"-l"})
}
