package logic

import (
	"context"
	"database/sql"
	"errors"
	"net/url"
	"strings"

	"shortLink/internal/svc"
	"shortLink/internal/types"
	"shortLink/model"
	"shortLink/pkg/ping"
	"shortLink/pkg/salt"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	// 数据校验
	// 1 不能为空
	// 2 长链接必须有效
	if !ping.New(req.LongURL) {
		// ping不通则拒绝
		return nil, errors.New("无效的链接")
	}
	// 3 不能是短链接
	// 获取短链接
	parsedURL, err := url.Parse(req.LongURL)
	if err != nil {
		logx.Errorw("err", logx.LogField{Key: "url.Parse", Value: err})
		return nil, err
	}
	path := parsedURL.Path
	parts := strings.Split(path, "/")
	lastURL := parts[len(parts)-1]
	_, err = l.svcCtx.MapModel.FindOneBySurl(l.ctx, sql.NullString{String: lastURL})
	if err != model.ErrNotFound {
		// 只要不是没有找到的错误都不能进行
		if err != nil {
			// 内部错误
			logx.Errorw("MapModel.FindOneBySurl", logx.Field("err", err))
			return nil, err
		}
		return nil, errors.New("已经存在")
	}

	// 4 没有转链过
	md5LongURL := salt.GetMD5(req.LongURL)
	_, err = l.svcCtx.MapModel.FindOneByLurl(l.ctx, sql.NullString{String: md5LongURL, Valid: true})
	if err != model.ErrNotFound {
		// 只要不是没有找到的错误都不能进行
		if err != nil {
			// 内部错误
			logx.Errorw("MapModel.FindOneByLurl", logx.Field("err", err))
			return nil, err
		}
		return nil, errors.New("已经存在")
	}

	return
}
