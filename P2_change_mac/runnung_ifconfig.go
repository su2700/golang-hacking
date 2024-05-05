package main

import (
	"flag"
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
	iface := flag.String( "iface","en0","Interface  you want to change MAC address" )
	newMac := flag.String( "newMac","","input new MAC address" )
    flag.Parse()
   executeCommand( "sudo",[]string{"ifconfig",*iface,"down"}  )
   executeCommand( "sudo",[]string{"ifconfig",*iface,"hw","ether",*newMac}  )
   executeCommand( "sudo",[]string{"ifconfig",*iface,"up"}  )


}