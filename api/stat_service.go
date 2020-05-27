package api

import (
	"context"

	"github.com/elkhan-ibrahimov/tagf/cmd"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// StatsService bla
type StatService struct{}

// AllTags bla
func (s *StatService) AllTags(ctx context.Context, _ *Empty) (*AllTagsRp, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "StatService.AllTags canceled")
	}

	tags, err := cmd.AllTagsAPI{}.Run()
	if err != nil {
		return nil, err
	}
	return &AllTagsRp{Tags: tags}, nil
}

// AllFiles bla
func (s *StatService) AllFiles(ctx context.Context, _ *Empty) (*AllFilesRp, error) {
	if ctx.Err() == context.Canceled {
		return nil, status.Errorf(codes.Canceled, "StatService.AllFiles canceled")
	}

	files, err := cmd.AllFilesAPI{}.Run()
	if err != nil {
		return nil, err
	}

	return &AllFilesRp{Files: files}, nil
}
