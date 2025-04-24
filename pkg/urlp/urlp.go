package urlp

import (
	"net/url"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

func GetLastPath(targetURL string) (string, error) {
	parsedURL, err := url.Parse(targetURL)
	if err != nil {
		logx.Errorw("err", logx.LogField{Key: "url.Parse", Value: err})
		return "", err
	}
	path := parsedURL.Path
	// 忽略/结尾
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	parts := strings.Split(path, "/")
	return parts[len(parts)-1], nil
}
