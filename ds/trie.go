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

package ds

import (
	"fmt"
	"time"

	"github.com/elkhan-ibrahimov/tagf/util"
)

func (x *Trie) find(tag string) (*Node, int) {
	runes := []rune(tag)

	// return nil if no any child nodes
	if len(x.Node) == 0 {
		return nil, -1
	}

	// return nil if node doesn't exists for the first character
	node, found := x.Node[runes[0]]
	if !found {
		return nil, -1
	}

	// return first found node if given tag length is one character
	if len(runes) == 1 {
		return node, 0
	}

	node, i := node.find(runes[1:])
	return node, i + 1
}

func (x *Trie) rootNode(chr rune) (root *Node, index int) {
	if x.Node == nil {
		x.Node = make(map[int32]*Node)
	}
	node, ok := x.Node[chr]
	if ok {
		return node, 0
	}
	return node, -1
}

// Has checks that provided tag does exist in the trie
func (x *Trie) Has(tag string) *Node {
	node, length := x.find(tag)

	if tagHasFound(tag, length) {
		return node
	} else {
		return nil
	}
}

// Put puts new tag to the tree
func (x *Trie) Put(tag string, files []string) {
	node, length := x.find(tag)

	// in case no any node found
	if node == nil || length == -1 {
		x.Fill(tag, files)

		// if tag partially found, then adding
	} else if !tagHasFound(tag, length) {
		node.FillChild(tag[length:], files)

		// if full node found, then just updating file list
	} else {
		node.Files = append(node.Files, files...)
	}
}

func (x *Trie) Fill(tag string, files []string) {
	runes := []rune(tag)
	childNode := &Node{}
	if len(x.Node) == 0 {
		x.Node = make(map[int32]*Node, 0)
	}
	x.Node[runes[0]] = childNode
	childNode.Fill(tag, files)
}

// Delete deletes tag from the trie
func (x *Trie) Delete(tag string) {
	node, index := x.find(tag)

	if tagHasFound(tag, index) {
		node.Final = false // not deleting, just setting final value to false
	}
}

func (x *Trie) AllTags() []string {
	tags := make([]string, 0)
	for _, node := range x.Node {
		tags = append(tags, node.AllTags()...)
	}
	return tags
}

func (x *Trie) AllFiles() []string {
	files := make([]string, 0)
	start := time.Now()
	defer util.Elapsed("Trie.AllFiles()")
	for _, node := range x.Node {
		nodeFiles := node.AllFiles()

		// file can have a multiple tags. so, I'm checking that this file is already exists in our list or not
		for _, f := range nodeFiles {
			if !util.ExistsInSlice(&files, f) {
				files = append(files, f)
			}
		}

	}
	fmt.Printf("%s took %v\n", "for Trie.AllFiles", time.Since(start))
	fmt.Printf("%s took %v\n", "for Trie.AllFiles 2", time.Now().Sub(start))
	return files
}

func (x *Trie) ToString() string {
	var s string
	for _, node := range x.Node {
		s = s + node.ToString()
	}
	return s
}
