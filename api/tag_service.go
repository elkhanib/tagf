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

package api

import (
	"context"

	"github.com/elkhan-ibrahimov/tagf/cmd"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TagService bla
type TagService struct{}

// Add bla
func (s *TagService) Add(ctx context.Context, t *AddTag) (*Empty, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "TagService.Add canceled")
	}

	command := cmd.AddAPI{AddCmd: cmd.AddCmd{Tags: t.Tags, Files: t.Files, Overwrite: t.Overwrite}}
	if err := command.Run(); err != nil {
		return nil, err
	}
	return &Empty{}, nil
}

// Search bla
func (s *TagService) Search(ctx context.Context, sr *SearchRq) (*SearchRp, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "TagService.Search canceled")
	}

	command := cmd.SearchAPI{SearchCmd: cmd.SearchCmd{Tags: sr.Tags}}
	r, err := command.Run()
	if err != nil {
		return nil, err
	}

	return &SearchRp{Files: convert(r)}, nil
}

func convert(m map[string][]string) []*FileWithTags {
	r := make([]*FileWithTags, 0)
	for file, tags := range m {
		r = append(r, &FileWithTags{File: file, Tags: tags})
	}

	return r
}
