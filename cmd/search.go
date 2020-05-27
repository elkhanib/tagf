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
	"github.com/elkhan-ibrahimov/tagf/common"
	"github.com/elkhan-ibrahimov/tagf/ds"
	"github.com/elkhan-ibrahimov/tagf/printer"
	"github.com/elkhan-ibrahimov/tagf/storage"
	"github.com/golang/protobuf/proto"
)

// SearchCmd bla bla
type SearchCmd struct {
	Tags []string
}

// SearchAPI  bla bla
type SearchAPI struct {
	SearchCmd
}

func search(s *SearchCmd) (map[string][]string, error) {
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

	// looking for files with tagged given tags
	tagAndFileMap := make(map[string][]string, 0)
	for _, t := range s.Tags {
		n := currTags.Has(t)
		if n != nil {
			tagAndFileMap[t] = n.Files
		}
	}

	return reverse(tagAndFileMap), nil
}

// reverse map from 'tag to files' relation to 'file to tags' relation
func reverse(m map[string][]string) map[string][]string {
	fileAndTagMap := make(map[string][]string)
	for tag, files := range m {
		for _, f := range files {
			fileAndTagMap[f] = append(fileAndTagMap[f], tag)
		}
	}

	return fileAndTagMap
}

// Run bla bla
func (s SearchAPI) Run() (map[string][]string, error) {
	return search(&s.SearchCmd)
}

// Run bla bla
func (s SearchCmd) Run() error {
	m, err := search(&s)

	if err != nil {
		return err
	}

	for file, tags := range m {
		t := common.FmtStringSlice(tags, ", ")
		printer.Success("%s -> %s\n", file, t)
	}

	return nil
}
