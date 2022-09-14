package main

import (
	"github.com/hansuh/kapicoin/explorer"
	"github.com/hansuh/kapicoin/rest"
)

// "github.com/hansuh/kapicoin/explorer"

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
