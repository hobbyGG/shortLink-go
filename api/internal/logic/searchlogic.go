package logic

import (
	"context"
	"database/sql"
	"fmt"

	"shortLink/api/internal/svc"
	"shortLink/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchRequest) (resp *types.SearchResponse, err error) {
	shortURL := req.ShortURL
	fmt.Println("shortURL:", shortURL)
	// 通过shortURL找到longURL
	// 将长连接返回给handler
	// handler使用302跳转到longURL
	// 后续添加redis缓存，并处理缓存雪崩和缓存击穿
	slMap, err := l.svcCtx.MapModel.FindOneBySurl(l.ctx, sql.NullString{String: shortURL, Valid: true})
	if err != nil {
		return nil, err
	}

	return &types.SearchResponse{LongURL: slMap.Lurl.String}, nil
}
