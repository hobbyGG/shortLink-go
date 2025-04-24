package svc

import (
	"shortLink/internal/config"
	"shortLink/model"
	mysqlse "shortLink/model/mysql"
	"shortLink/model/sequence"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config   config.Config
	SeqModel model.SequenceModel
	MapModel model.ShortUrlMapModel

	Sequence sequence.Methods

	ShortURLBlackList map[string]struct{}
}

func NewServiceContext(c config.Config) *ServiceContext {
	seqConn := sqlx.NewMysql(c.SLMapMysql.Datasource)
	mapConn := sqlx.NewMysql(c.SLMapMysql.Datasource)

	return &ServiceContext{
		Config:   c,
		SeqModel: model.NewSequenceModel(seqConn),
		MapModel: model.NewShortUrlMapModel(mapConn),
		Sequence: mysqlse.NewClient(c.SequenceMysql.Datasource),
		ShortURLBlackList: mysqlse.InitBlackList(c.ShortURLBlackList),
	}
}
