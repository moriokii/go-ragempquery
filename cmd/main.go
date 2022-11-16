package main

import (
	"fmt"

	"github.com/moriokii/go-ragempquery"
)

func main() {
	a, err := ragempquery.NewClient("Server:port")
	if err != nil {
		print(err.Error())
	}

	r, err := a.Info()
	if err != nil {
		panic(err)
	}
	fmt.Print(r.Gamemode)
}
