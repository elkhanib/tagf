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
	"github.com/golang/protobuf/proto"
)

// AddCmd represents add tagging through command-line
type AddCmd struct {
	Tags      []string
	Files     []string
	Overwrite bool
}

// AddAPI represents tagging structure through API
type AddAPI struct {
	AddCmd
}

func add(m *AddCmd) error {
	// validating files
	if err := checkFiles(m.Files); err != nil {
		return err
	}

	// validating tags
	if err := checkTags(m.Tags); err != nil {
		return err
	}

	// loading existing file
	file, err := storage.Load()
	if err != nil {
		return err
	}

	// deserialization
	var currTags = ds.Trie{}
	err = proto.Unmarshal(file, &currTags)
	if err != nil {
		return err
	}

	for _, t := range m.Tags {
		currTags.Put(t, m.Files)
	}

	// serialization
	file, err = proto.Marshal(&currTags)
	if err != nil {
		return err
	}

	return storage.Save(file)
}

// Run adding tags to file
func (m AddCmd) Run() error {
	if err := add(&m); err != nil {
		return err
	}

	return printer.Success("Tags are successfully added")
}

// Run adding tags to file
func (m AddAPI) Run() error {
	return add(&m.AddCmd)
}
