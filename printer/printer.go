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

package printer

import (
	"github.com/fatih/color"
)

var (
	errorColor   = color.New(color.FgRed)
	warningColor = color.New(color.FgYellow)
	successColor = color.New(color.FgHiGreen)
)

func init() {
	errorColor.EnableColor()
	warningColor.EnableColor()
	successColor.EnableColor()
}

// Error prints standard error (stderr) with colorized
func Error(msg string) error {
	_, err := errorColor.Fprintln(color.Error, msg)
	return err
}

// Warning prints standard output (stdout) with colorized
func Warning(msg string) error {
	_, err := warningColor.Fprintln(color.Output, msg)
	return err
}

// Success prints standard output (stdout) with colorized
func Success(format string, msg ...interface{}) error {
	_, err := successColor.Fprintf(color.Output, format, msg...)
	return err
}
