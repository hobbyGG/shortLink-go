package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	SequenceMysql struct {
		Datasource string
	}
	SLMapMysql struct {
		Datasource string
	}

	ShortURLBlackList []string
}
