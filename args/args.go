// Copyright 2014 BertWednesdays Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Simple command line parser
//

package args

import (
	"fmt"
	"os"
	"strings"
)

// Argument map
type arguments_map map[string]string

type ArgumentsHandler struct {
	args       arguments_map // map containing the parsed command line arguments
	args_error string        // error string
	args_map   arguments_map // argument list passed in
}

// ---------------------------------------------------------------------------------
// Public Functions
// ---------------------------------------------------------------------------------

// New allocates and returns a new ArgumentsHandler
func New(dict arguments_map, in_args []string) *ArgumentsHandler {

	// Parse the command line arguments
	args, error_message := parse(dict, in_args)

	return &ArgumentsHandler{args: args, args_error: error_message, args_map: dict}
}

func CreateWithOSArgs(dict arguments_map) *ArgumentsHandler {

	// Remove the first element for the os args array
	os_args := os.Args
	os_args = os_args[:0+copy(os_args[0:], os_args[0+1:])]

	return New(dict, os_args)
}

// Returns the number of params available
func (self *ArgumentsHandler) Count() int {
	return len(self.args)
}

// Returns true if the command has an argument
func (self *ArgumentsHandler) HasParam(cmd string) bool {
	_, ok := self.args[cmd]
	return ok
}

// Returns the command parameter
func (self *ArgumentsHandler) GetParam(cmd string) string {
	if val, ok := self.args[cmd]; ok {
		return val
	}
	return nil
}

// Returns true if there is an error
func (self *ArgumentsHandler) HasError() bool {
	return len(self.args_error) > 0
}

// Get the error parameter
func (self *ArgumentsHandler) GetError() string {
	return self.args_error
}

// Get a string containing the Usage information
func (self *ArgumentsHandler) GetUsage() string {
	usage := "Please provide the following command line arguments:\n"
	usage += fmt.Sprintf("   --%s %s\n", "help", "Display the usage information.")

	if len(self.args_map) > 0 {
		for key, value := range self.args_map {
			usage += fmt.Sprintf("   --%s <%s>\n", key, value)
		}
	} else {
		usage += "No Argument list provided"
	}
	return usage
}

// Print any Error and a usage message to the console
func (self *ArgumentsHandler) PrintUsageAndErrors() {
	if self.HasError() {
		fmt.Println(self.GetError())
	}
	fmt.Println(self.GetUsage())
}

// ---------------------------------------------------------------------------------
// End Public Functions
// ---------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------
// Private Functions
// ---------------------------------------------------------------------------------

// Parse the command line arguments
func parse(dict arguments_map, arguments []string) (args arguments_map, args_err string) {
	if len(dict) > 0 {

		//fmt.Printf("%v", arguments)

		prefix := "--"
		in_args_len := len(arguments)

		if in_args_len%2 != 0 {

			args_err = "Error : Arguments list is invalid."

			// Check to see if help has been requested
			if in_args_len == 1 {
				if strings.HasPrefix(arguments[0], prefix) {
					cmd := strings.TrimPrefix(arguments[0], prefix)
					if cmd == "help" || cmd == "h" {
						args_err = "Help\n\n"
					}
				}
			}

		} else {
			args = make(arguments_map, len(dict))
			for i := 0; i < in_args_len; i += 2 {
				if strings.HasPrefix(arguments[i], prefix) {
					cmd := strings.TrimPrefix(arguments[i], prefix)
					if _, ok := dict[cmd]; ok {
						args[cmd] = arguments[i+1]
					} else {
						args_err = "Error : Invalid Command " + cmd + "."
					}
				} else {
					args_err = "Error : Invalid prefix on value " + arguments[i] + "."
				}
			}

		}
	} else {
		args_err = "Error : No arguments provided."
	}
	return args, args_err
}

// ---------------------------------------------------------------------------------
// End Private Functions
// ---------------------------------------------------------------------------------
