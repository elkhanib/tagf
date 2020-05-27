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
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/elkhan-ibrahimov/tagf/api"

	"github.com/elkhan-ibrahimov/tagf/cmd"
	"github.com/elkhan-ibrahimov/tagf/common"
)

type arrayFlag []string

func (a *arrayFlag) String() string {
	return fmt.Sprintln(*a)
}

func (a *arrayFlag) Set(value string) error {
	*a = common.Split(value, ",")
	return nil
}

var (
	flagTags        arrayFlag
	flagFiles       arrayFlag
	flagOverwrite   bool
	flagSearch      bool
	flagAdd         bool
	flagClear       bool
	flagStartServer bool
	flagAllTags     bool
	flagAllFiles    bool
)

func parseArgs() Cmd {
	flag.BoolVar(&flagClear, "c", false, "clear all tags from files")
	flag.BoolVar(&flagAdd, "a", false, "add tags to files")
	flag.BoolVar(&flagSearch, "s", false, "search for tags")
	flag.BoolVar(&flagOverwrite, "o", false, "overwrite previous tags")
	flag.BoolVar(&flagStartServer, "gs", false, "start tagf api server")
	flag.BoolVar(&flagAllTags, "at", false, "get all tags list")
	flag.BoolVar(&flagAllFiles, "af", false, "get all tagged files list")
	flag.Var(&flagTags, "t", "tags")
	flag.Var(&flagFiles, "f", "files to add tag")
	flag.Parse()

	// check for search command
	if flagSearch {
		if len(flagTags) == 0 {
			return cmd.UnknownCmd{Err: fmt.Errorf("'-s' flag needs tags argument to search")}
		}
		return cmd.SearchCmd{Tags: flagTags}
	}

	// check for add command
	if flagAdd {
		if len(flagTags) == 0 || len(flagFiles) == 0 {
			return cmd.UnknownCmd{Err: fmt.Errorf("'-a' flag needs tags and files arguments")}
		}
		return cmd.AddCmd{Tags: flagTags, Files: flagFiles, Overwrite: flagOverwrite}
	}

	// check for clear command
	if flagClear {
		if len(flagFiles) == 0 {
			return cmd.UnknownCmd{Err: fmt.Errorf("'-c' flag needs files argument")}
		}
		return cmd.ClearCmd{Files: flagFiles}
	}

	// check for start-api-server command
	if flagStartServer {
		return api.StartAPICmd{}
	}

	if flagAllTags {
		return cmd.AllTagsCmd{}
	}

	if flagAllFiles {
		return cmd.AllFilesCmd{}
	}

	// otherwise returning help command
	return cmd.HelpCmd{}
}

// getting absolute path of provided file
func absolutePath(p string) (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}
	return path.Join(homeDir, path.Clean(p)), nil
}
