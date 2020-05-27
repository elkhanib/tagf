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

package common

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

// Split breaks a string down into a list of substrings using separator
func Split(str string, sp string) []string {
	f := func(c rune) bool {
		return string(c) == sp
	}
	return strings.FieldsFunc(str, f)
}

// AbsolutePath gets absolute path of provided file
func AbsolutePath(p string) (string, error) {
	if path.IsAbs(p) {
		return p, nil
	}
	currentDir, err := os.Getwd()

	if err != nil {
		return "", err
	}
	return path.Join(currentDir, path.Clean(p)), nil
}

func FmtStringSlice(str []string, separator string) string {
	if len(str) == 1 {
		return str[0]
	}

	var formatted string
	for i, f := range str {
		if i == len(str)-1 {
			formatted = formatted + f
		} else {
			formatted = formatted + f + separator
		}

	}
	return formatted
}

func ExistsInSlice(s *[]string, elem string) bool {
	for _, e := range *s {
		if e == elem {
			return true
		}
	}

	return false
}

func Elapsed(action string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", action, time.Since(start))
	}
}
