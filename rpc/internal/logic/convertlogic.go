package logic

import (
	"context"

	"shortLink/rpc/internal/svc"
	"shortLink/rpc/pb/shortLink"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConvertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConvertLogic) Convert(in *shortLink.ConvertRequest) (*shortLink.ConvertResponse, error) {
	// todo: add your logic here and delete this line

	return &shortLink.ConvertResponse{}, nil
}
