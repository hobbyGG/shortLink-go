package logic

import (
	"context"

	"shortLink/rpc/internal/svc"
	"shortLink/rpc/pb/shortLink"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchLogic) Search(in *shortLink.SearchRequest) (*shortLink.SearchResponse, error) {
	// todo: add your logic here and delete this line

	return &shortLink.SearchResponse{}, nil
}
