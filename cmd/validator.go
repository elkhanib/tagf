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
	"os"
	"regexp"

	"github.com/elkhan-ibrahimov/tagf/common"
	"github.com/elkhan-ibrahimov/tagf/errors"
)

func checkFiles(files []string) error {
	nonExistingFiles := make([]string, 0)
	for i, f := range files {
		absPath, err := common.AbsolutePath(f)
		if err != nil {
			return errors.AbsolutePathError(err, f)
		}

		files[i] = absPath
		if !fileExists(absPath) {
			nonExistingFiles = append(nonExistingFiles, absPath)
		}
	}

	if len(nonExistingFiles) == 0 {
		return nil
	}

	return errors.FileNotExistError(nonExistingFiles)
}

func fileExists(fileName string) bool {
	f, err := os.Stat(fileName)
	notExists := os.IsNotExist(err) || f.IsDir()
	return !notExists
}

func checkTags(tags []string) error {
	invalidTags := make([]string, 0)
	for _, t := range tags {
		matched, err := regexp.MatchString("^[\\w/\\p{L}!@#%&*'\"]+$", t)
		if err != nil || !matched {
			invalidTags = append(invalidTags, t)
		}
	}

	if len(invalidTags) == 0 {
		return nil
	}

	return errors.InvalidTagError(invalidTags)
}
