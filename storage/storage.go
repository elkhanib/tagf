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

package storage

import (
	"github.com/elkhan-ibrahimov/tagf/common"
	"github.com/elkhan-ibrahimov/tagf/errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Save stores file in the storage
func Save(raw []byte) error {
	if err := createDir(common.FolderPath); err != nil {
		return errors.CreateFolderError(err, common.FolderPath)
	}

	file, err := os.Create(filepath.Join(common.FolderPath, common.FileName))
	if err != nil {
		return errors.CreateFileError(err, common.FileName)
	}
	defer file.Close()

	if _, err := file.Write(raw); err != nil {
		return errors.WriteFileError(err)
	}

	// forcing OS to write cached data to file
	return errors.FlushFileError(file.Sync())
}

func createDir(folderPath string) error {
	return os.MkdirAll(folderPath, 0700)
}

// Load loads file content as a byte array
func Load() ([]byte, error) {
	p := filepath.Join(common.FolderPath, common.FileName)

	// reading file content
	file, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, Save(nil)
	}

	return file, nil
}
