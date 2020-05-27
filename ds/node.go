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

func (x *Node) find(runes []rune) (*Node, int) {
	node, ok := x.Child[runes[0]]
	i := 0
	if ok {
		if len(runes) > 1 {
			node, i = node.find(runes[1:])
		}
		i++
	}

	if i == 0 {
		return x, i
	}
	return node, i
}

func (x *Node) Fill(tag string, files []string) {
	runes := []rune(tag)
	x.Char = runes[0]
	childNode := &Node{}
	x.Child = make(map[int32]*Node)
	if len(runes) == 1 {
		x.Final = true
		x.Files = files
	} else {
		x.Final = false
		x.Child[runes[1]] = childNode
		childNode.Fill(tag[1:], files)
	}
}

func (x *Node) FillChild(tag string, files []string) {
	runes := []rune(tag)
	if len(x.Child) == 0 {
		x.Child = make(map[int32]*Node)
	}
	childNode := &Node{}
	x.Child[runes[0]] = childNode
	childNode.Fill(tag, files)
}

func (x *Node) ToString() string {
	tags := x.FormatOutput()
	var tagsConcatenated string

	for tag := range tags {
		tagsConcatenated = tagsConcatenated + tag + "\n"
	}

	return tagsConcatenated
}

func (x *Node) AllTags() []string {
	tagAndFilesMap := x.FormatOutput()
	tags := make([]string, 0)
	for t := range tagAndFilesMap {
		tags = append(tags, t)
	}

	return tags
}

func (x *Node) AllFiles() []string {
	tagAndFilesMap := x.FormatOutput()
	files := make([]string, 0)
	for _, f := range tagAndFilesMap {
		files = append(files, f...)
	}

	return files
}

func (x *Node) FormatOutput() map[string][]string {
	return formatNodeOutput(x, "")
}

func formatNodeOutput(node *Node, rootTag string) map[string][]string {
	rootTag = rootTag + string(node.Char)

	tag := make(map[string][]string)
	for v, _ := range node.Child {
		s := node.Child[v]
		concatenateMap(&tag, formatNodeOutput(s, rootTag))
	}

	if node.Final {
		tag[rootTag] = node.Files
	}

	return tag
}

func tagHasFound(tag string, foundIndex int) bool {
	return len(tag) == foundIndex
}

func concatenateMap(a *map[string][]string, b map[string][]string) {
	for k, v := range b {
		(*a)[k] = v
	}
}
