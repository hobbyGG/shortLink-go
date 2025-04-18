package svc

import (
	"shortLink/internal/config"
	"shortLink/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config   config.Config
	SeqModel model.SequenceModel
	MapModel model.ShortUrlMapModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	seqConn := sqlx.NewMysql(c.SLMapMysql.Datasource)
	mapConn := sqlx.NewMysql(c.SLMapMysql.Datasource)
	return &ServiceContext{
		Config:   c,
		SeqModel: model.NewSequenceModel(seqConn),
		MapModel: model.NewShortUrlMapModel(mapConn),
	}
}
