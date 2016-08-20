// Copyright (c) 2016 Pepijn van Leeuwen
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
	"github.com/pepijnvanleeuwen/gopep/configuration"
	"github.com/pepijnvanleeuwen/gopep/twitter"
)

var opts struct {
	Module string `short:"m" long:"module" description:"The module to use. Required." required:"true"`
	Action string `short:"a" long:"action" description:"The action to perform. Required." required:"true"`
	Value  string `short:"v" long:"value" description:"The value to pass to the action. Optional."`
}

func main() {
	if err := configuration.LoadConfig(); err != nil {
		printErr(err)
	}

	if _, err := flags.Parse(&opts); err != nil {
		printErr(err)
	}

	err := LoadModule()
	printErr(err)
}

func LoadModule() error {
	switch opts.Module {
	case "twitter":
	case "t":
		return twitter.Load(opts.Action, opts.Value)
	}

	return fmt.Errorf("Module '%s' is not supported", opts.Module)
}

func printErr(err error) {
	if err != nil {
		panic(err)
	}
}
