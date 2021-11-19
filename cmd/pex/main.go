package main

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml"
	"go.d34d.net/pex"
)

func usage() {
	fmt.Printf("%s <settings.toml>\n", os.Args[0])
}

func main() {
	if len(os.Args) == 2 {
		settings, err := toml.LoadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		pex.API(settings)
	} else {
		usage()
	}
}
