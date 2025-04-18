package salt

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5(longURL string) string {
	md := md5.New()
	md.Write([]byte(longURL))
	// 不用加密所以不设置secrete
	return hex.EncodeToString(md.Sum(nil))
}
