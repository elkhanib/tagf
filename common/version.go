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

package common

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	Major uint
	Minor uint
	Patch uint
}

func ParseVersion(version string) (Version, error) {
	numbers := strings.Split(version, ".")

	major, err := strconv.Atoi(numbers[0])
	if err != nil {
		return Version{}, errors.New("invalid major version")
	}

	minor, err := strconv.Atoi(numbers[1])
	if err != nil {
		return Version{}, errors.New("invalid minor version")
	}

	patch, err := strconv.Atoi(numbers[2])
	if err != nil {
		return Version{}, errors.New("invalid patch version")
	}

	return Version{uint(major), uint(minor), uint(patch)}, nil
}

func (v Version) String() string {
	return fmt.Sprintf("%v.%v.%v", v.Major, v.Minor, v.Patch)
}

func (v Version) LessThan(comparedVer Version) bool {
	return v.Major < comparedVer.Major ||
		(v.Major == comparedVer.Major && v.Minor < comparedVer.Minor) ||
		(v.Major == comparedVer.Major && v.Minor == comparedVer.Minor && v.Patch < comparedVer.Patch)
}

func (v Version) GreaterThan(comparedVer Version) bool {
	return v.Major > comparedVer.Major ||
		(v.Major == comparedVer.Major && v.Minor > comparedVer.Minor) ||
		(v.Major == comparedVer.Major && v.Minor == comparedVer.Minor && v.Patch > comparedVer.Patch)
}
