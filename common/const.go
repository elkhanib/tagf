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
	"github.com/elkhan-ibrahimov/tagf/printer"
	"os"
	"path/filepath"
)

const TagMaxLength = 100
const FileName = ".tags"
const EnableDebug = "DEBUG_TAGF"

// AppVersion indicates version of application
var AppVersion, _ = ParseVersion("0.1.0")

// SchemaVersion indicates database schema version
var SchemaVersion, _ = ParseVersion("0.1.0")

// FolderPath represents absolute path of application config folder
var FolderPath string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		_ = printer.Error(fmt.Sprintf("couldn't get user's home folder: %s", err.Error()))
		os.Exit(1)
	}

	FolderPath = filepath.Join(homeDir, ".tagf")
}
