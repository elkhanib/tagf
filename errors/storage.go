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

package errors

import (
	"fmt"
)

const (
	msgCreateFolderError = "couldn't create folder: '%s'"
	msgCreateFileError   = "couldn't create file: '%s'"
	msgWriteFileError    = "couldn't write file content"
	msgFlushFileError    = "couldn't flush cached data to a file"
	msgLoadFileError     = "couldn't open file: '%s'"
	msgAbsolutePathError = "couldn't get absolute path of file: '%s'"
)

func CreateFolderError(err error, folderPath string) error {
	return Wrap(err, fmt.Sprintf(msgCreateFolderError, folderPath))
}

func CreateFileError(err error, fileName string) error {
	return Wrap(err, fmt.Sprintf(msgCreateFileError, fileName))
}

func WriteFileError(err error) error {
	return Wrap(err, msgWriteFileError)
}

func FlushFileError(err error) error {
	return Wrap(err, msgFlushFileError)
}

func LoadFileError(err error, path string) error {
	return Wrap(err, fmt.Sprintf(msgLoadFileError, path))
}

func AbsolutePathError(err error, file string) error {
	return Wrap(err, fmt.Sprintf(msgAbsolutePathError, file))
}
