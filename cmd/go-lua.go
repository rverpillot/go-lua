package main

import (
	"os"

	"github.com/pkg/profile"
	"github.com/rverpillot/go-lua"
)

func main() {
	l := lua.NewState()
	lua.OpenLibraries(l)
	defer profile.Start().Stop()
	if err := lua.DoFile(l, os.Args[1]); err != nil {
		panic(err)
	}
}
