package repos

import (
	"context"

	"github.com/yosa12978/hjkl/types"
)

type PostRepo interface {
	GetAll(ctx context.Context) []types.Post
	GetById(ctx context.Context, id string) (*types.Post, error)
	GetPagedFromTimestamp(ctx context.Context, timestamp int64, page, limit int) *types.Page[types.Post]
}
