package main

import (
	"fmt"
	//"go/constant"
	"net"
	//"os"
)


func main() {
   IP := "scanme.nmap.org"
   for i := 1; i<100 ; i++ {
   address := fmt.Sprintf(IP+ ":%d", i)
   connection, err := net.Dial("tcp" , address)
   if err == nil {
	   fmt.Printf(" [+] Connection established, port %v open ", i, connection.RemoteAddr().String())
	   
   }
}
}

//