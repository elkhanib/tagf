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
limitations under the License.g
*/

package cmd

type ExitStatus uint8

const (
	Success ExitStatus = 0
	Error              = 1
)

func (e ExitStatus) Code() int {
	return int(e)
}

func Run() error {
	//// parsing arguments
	//files, tags, err := main.parseArgs()
	//if err != nil {
	//	return err
	//}
	//
	//// validating files argument
	//if err := checkFiles(files); err != nil {
	//	return err
	//}
	//
	//// validating tags argument
	//if err := checkTags(tags); err != nil {
	//	return err
	//}
	//
	//// loading existing tags
	//existingFile, err := storage.Load()
	//if err != nil {
	//	return err
	//}
	//
	//// deserialization
	//var existingTags = &scheme.Trie{}
	//err = proto.Unmarshal(existingFile, existingTags)
	//if err != nil {
	//	return err
	//}
	//
	//for _, t := range tags {
	//	existingTags.Put(t, files)
	//}
	//
	//// serialization
	//existingFile, err = proto.Marshal(existingTags)
	//if err != nil {
	//	return err
	//}
	//
	//if err := storage.Save(existingFile); err != nil {
	//	return err
	//}
	//
	return nil
}
