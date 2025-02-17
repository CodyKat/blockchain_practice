package cli

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/CodyKat/blockchain_practice/explorer"
	"github.com/CodyKat/blockchain_practice/rest"
)

func usage() {
	fmt.Println("hello my name is jaemjeon")
	fmt.Println("Please use the following flags :")
	fmt.Println("-port: 	Set the PORT of the server")
	fmt.Println("=mode: 	Start the REST API")
	runtime.Goexit()
}

func Start() {
	port := flag.Int("port", 4000, "Sets port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
	fmt.Println(*port, *mode)
}
