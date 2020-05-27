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
	"errors"
	"fmt"
	"github.com/elkhan-ibrahimov/tagf/util"
)

const (
	msgInvalidTagError   = "below tag(s) are not in valid format\n%v"
	msgFileNotExistError = "below file(s) doesn't exists\n%v"
)

func InvalidTagError(tags []string) error {
	if len(tags) == 0 {
		return nil
	}
	s := util.FmtStringSlice(tags, " ")
	return errors.New(fmt.Sprintf(msgInvalidTagError, s))
}

func FileNotExistError(files []string) error {
	if len(files) == 0 {
		return nil
	}
	s := util.FmtStringSlice(files, "\n")
	return errors.New(fmt.Sprintf(msgFileNotExistError, s))
}
