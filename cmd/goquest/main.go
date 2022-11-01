package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bnelsonjc/goquest/internal/goquest/engine"
	"github.com/bnelsonjc/goquest/internal/goquest/version"
)

const (

	// ExitSuccess indicates a successful program execution.
	ExitSuccess = iota

	// ExitGameError indicates an unsuccessful program execution due to a problem during the game.
	ExitGameError
)

var (
	returnCode  int   = ExitSuccess
	flagVersion *bool = flag.Bool("version", false, "Gives the version info")
)

func main() {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			// we are panicking, make sure we dont lose the panic just because we checked
			panic("unrecoverable panic occured")
		} else {
			os.Exit(returnCode)
		}
	}()

	flag.Parse()

	if *flagVersion {
		fmt.Printf("%s\n", version.Current)
		return
	}

	gameEng := engine.New(os.Stdin, os.Stdout)

	err := gameEng.RunUntilQuit()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
	}
}
