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
