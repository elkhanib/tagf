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

package cmd

import (
	"github.com/elkhan-ibrahimov/tagf/ds"
	"github.com/elkhan-ibrahimov/tagf/printer"
	"github.com/elkhan-ibrahimov/tagf/storage"
	"github.com/elkhan-ibrahimov/tagf/util"
	"github.com/golang/protobuf/proto"
)

// SearchCmd bla bla
type AllTagsCmd struct {
}

// SearchAPI  bla bla
type AllTagsAPI struct {
	AllTagsCmd
}

func allTags(s *AllTagsCmd) ([]string, error) {
	// loading existing file
	file, err := storage.Load()
	if err != nil {
		return nil, err
	}

	// deserialization
	var currTags = ds.Trie{}
	err = proto.Unmarshal(file, &currTags)
	if err != nil {
		return nil, err
	}

	return currTags.AllTags(), nil
}

// Run bla bla
func (s AllTagsAPI) Run() ([]string, error) {
	return allTags(&s.AllTagsCmd)
}

// Run bla bla
func (s AllTagsCmd) Run() error {
	tags, err := allTags(&s)

	if err != nil {
		return err
	}

	if len(tags) == 0 {
		return printer.Warning("you don't have any tag")
	}

	t := util.FmtStringSlice(tags, ", ")
	return printer.Success(t)
}
