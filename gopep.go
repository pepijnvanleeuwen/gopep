package main

import (
	"fmt"
	"log"

	"github.com/jessevdk/go-flags"
	"github.com/pepijnvanleeuwen/gopep/configuration"
	"github.com/pepijnvanleeuwen/gopep/twitter"
)

var opts struct {
	Module string `short:"m" long:"module" description:"The module to use."`
	Action string `short:"a" long:"action" description:"The action to perform."`
	Value  string `short:"v" long:"value" description:"The value to pass to the selected action."`
}

func main() {
	log.Println("Loading config.json")
	err := configuration.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Parsing flags")
	flags.Parse(&opts)

	err = LoadModule()
	if err != nil {
		log.Fatalln(err)
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
