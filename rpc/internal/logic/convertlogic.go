package logic

import (
	"context"
	"database/sql"
	"errors"

	"shortLink/model"
	"shortLink/pkg/base62"
	"shortLink/pkg/salt"
	"shortLink/pkg/urlp"
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
	// 数据校验
	// 1 不能为空
	// 2 长链接必须有效
	// if !ping.New(req.LongURL) {
	// 	// ping不通则拒绝
	// 	return nil, errors.New("无效的链接")
	// }
	// 3 不能是短链接
	// 获取短链接
	lastURL, err := urlp.GetLastPath(in.GetLongUrl())
	if err != nil {
		logx.Errorw("convertlogic", logx.Field("err", err))
		return nil, err
	}
	_, err = l.svcCtx.MapModel.FindOneBySurl(l.ctx, sql.NullString{String: lastURL, Valid: true})
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
	md5LongURL := salt.GetMD5(in.LongUrl)
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

	var shortPath string
	for {
		// 取号
		// 使用replace插入sequence，获取新的主键
		seq, err := l.svcCtx.Sequence.Get()
		if err != nil {
			logx.Errorw("convertlogic error", logx.Field("l.svcCtx.Sequence.Get", err))
		}

		// 转为base62
		shortPath = base62.Uint2string(seq)
		// 不能出现特殊词
		if _, ok := l.svcCtx.ShortURLBlackList[shortPath]; !ok {
			// 没有出现特殊词
			break
		}
	}

	ShortURL := "127.0.0.1/" + shortPath
	_, err = l.svcCtx.MapModel.Insert(l.ctx, &model.ShortUrlMap{
		Lurl: sql.NullString{String: md5LongURL, Valid: true},
		Surl: sql.NullString{String: ShortURL, Valid: true},
	})
	if err != nil {
		logx.Errorw("convert error", logx.Field("error", err))
		return nil, err
	}

	return &shortLink.ConvertResponse{ShortUrl: ShortURL}, nil
}
