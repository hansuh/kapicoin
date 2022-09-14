package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/hansuh/kapicoin/explorer"
	"github.com/hansuh/kapicoin/rest"
)

func usage() {
	fmt.Printf("Welcome to KapiCoin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port:	Sets the port of the server\n")
	fmt.Printf("-mode:	Choose between 'html', 'rest', and 'all'\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Sets the port of the server")
	portalt := flag.Int("portalt", 5000, "Sets the port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	case "all":
		go rest.Start(*port)
		explorer.Start(*portalt)
	default:
		usage()
	}
}
