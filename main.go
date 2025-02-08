package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Println("hello my name is jaemjeon")
	fmt.Println("Please use the following command :")
	fmt.Println("explorer : Start the HTML explorer")
	fmt.Println("rest : 	Start the REST API")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(0)
	}

	rest := flag.NewFlagSet("rest", flag.ExitOnError)

	portFlag := rest.Int("port", 4000, "Sets the port of this server")

	switch os.Args[1] {
	case "explorer":
	case "rest":
		rest.Parse(os.Args[2:])
		fmt.Println(*portFlag)
	default:
		usage()
	}
}
