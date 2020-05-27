/*
Copyright 2020 Elkhan Ibrahimov

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"os"

	"github.com/elkhan-ibrahimov/tagf/common"
	"github.com/elkhan-ibrahimov/tagf/printer"
)

// Cmd bla bla
type Cmd interface {
	Run() error
}

func main() {
	c := parseArgs()
	if err := c.Run(); err != nil {
		_ = printer.Error(err.Error())

		if _, ok := os.LookupEnv(common.EnableDebug); ok {
			// print stack trace in verbose mode
			_, _ = fmt.Fprintf(os.Stderr, "[DEBUG] error: %+v\n", err)
		}
		defer os.Exit(1)
	}
}
