package main

import (
	"os"

	"github.com/Shopify/go-lua"
	"github.com/pkg/profile"
)

func main() {
	l := lua.NewState()
	lua.OpenLibraries(l)
	defer profile.Start().Stop()
	if err := lua.DoFile(l, os.Args[1]); err != nil {
		panic(err)
	}
}
