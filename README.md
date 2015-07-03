# README #

This is a very simple command line parser.

### Install

```bash
go get github.com/bertmcdowell/goargs
```

### Usage

app --config init.ini

### Sample
```go
package main

import (
	"github.com/bertmcdowell/goargs/args"
	"fmt"
)

// Create a map to specify the commands and a help message
var commands = map[string]string{
    "config": "Path to configuration file.",
}

func main() {

	// Parse the command line arguments
	args := args.CreateWithOSArgs(commands)

	if args.HasError() {
		// Print any error and the usage information
		args.PrintUsageAndErrors() 
	} else {
		// You can check to see if we have the parameter
		if args.HasParam("config") {
			fmt.Println("Parameter was found")
		}

		// Or you can grab the parameter
		configFile = args.GetParam("config")

		// Check the parameter is not nil
		if configFile == "" {
			fmt.Println("Parameter was not found")	
		} else { 
			fmt.Println(configFile)
		}
	}
}
```

