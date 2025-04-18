package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	SequenceMysql struct {
		Datasource string
	}
	SLMapMysql struct {
		Datasource string
	}
}
