package main

import (
	"fmt"
	"os"
	"net"
	"time"
)


func getIpArg() (string,error){

	type Arguments struct{
		target string
	}

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr,"Sytanx Error:\nUsage: dnsping <DOMAIN NAME> \nExample: dnsping facebook.com\n")
		os.Exit(0)
		return "",fmt.Errorf("Syntax Error: Wrong IP Address Argument")
	}

	operation := Arguments{os.Args[1]}

	return operation.target,nil

}

func resolver(inputString string,incr int) {
	start := time.Now()
	resolved,_ := net.ResolveIPAddr("ip",inputString)
	if resolved == nil {
	fmt.Printf("unable to resolve %v\n",inputString)

	} else {
		elapsed := time.Since(start)
		fmt.Printf("resolved %v to %v resolution_seq=%v time=%v\n",inputString,resolved,incr,elapsed)	
	}



}

func main() {
	target,_ := getIpArg()
	for x := 0; x <= 5; x++ {
		resolver(target,x)
	}
}