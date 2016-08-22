package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
	"github.com/pepijnvanleeuwen/gopep/configuration"
	"github.com/pepijnvanleeuwen/gopep/twitter"
)

var opts struct {
	Module string `short:"m" long:"module" description:"The module to use. Required."`
	Action string `short:"a" long:"action" description:"The action to perform. Required."`
	Value  string `short:"v" long:"value" description:"The value to pass to the action."`
}

func main() {
	err := configuration.LoadConfig()
	if err != nil {
		panic(err)
	}

	flags.Parse(&opts)

	err = LoadModule()
	if err != nil {
		fmt.Println(err)
	}
}

// LoadModule loads the requested module and passes the action and value to
// that module.
func LoadModule() error {
	switch opts.Module {
	case "twitter", "t":
		return twitter.Load(opts.Action, opts.Value)
	case "":
		return fmt.Errorf("Please specify the module.")
	}

	return fmt.Errorf("Module '%s' is not supported", opts.Module)
}
