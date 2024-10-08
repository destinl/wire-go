package service

import (
	"context"
	"fmt"
)

type IPostService interface {
	GetPostById(ctx context.Context, id string) string
}

var _ IPostService = (*PostService)(nil)

type PostService struct {
}

func (s *PostService) GetPostById(ctx context.Context, id string) string {
	return fmt.Sprint("欢迎关注公众号:Go技术干货,作者:陈明勇")
}

func NewPostService() IPostService {
	return &PostService{}
}

// 在 Go 中的 最佳实践 是返回 具体的类型 的值，所以最好让 NewPostService 返回具体类型 PostService 的值，而不是接口类型 IPostService。
func NewPostServiceV2() *PostService {
	return &PostService{}
}
