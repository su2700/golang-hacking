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
	cmd_obj.Stdin = os.Stdin
	err = cmd_obj.Run()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		
	}
	return nil
}
func main(){
	// ifconfig eth0 down
	// ifconfig eth0 hw ether 00:
	// ifconfig eth0 up
   executeCommand( "sudo",[]string{"ifconfig","en0","down"}  )
   executeCommand( "sudo",[]string{"ifconfig","en0","hw","ether","00:00:00:00:00:00"}  )
   executeCommand( "sudo",[]string{"ifconfig","en0","up"}  )


}