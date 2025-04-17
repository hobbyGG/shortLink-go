package main

// 重定向服务

import (
	"demo/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/:tag", redirectHandler)

	r.Run(":8888")
}

func redirectHandler(c *gin.Context) {
	// 获取路径参数
	tag := c.Param("tag")

	// 通过路径参数查数据库
	if tag == "lidglwL" {
		// c.Writer().Header().Set(key, value)
		c.Header(constant.StrLocation, constant.StrRedirectURL)
		c.Status(http.StatusFound)
	} else {
		// 一般要返回404
		c.Status(http.StatusNotFound)
	}
}
